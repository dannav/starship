/*
	Initially only support markdown and text files

	enhance to support .docx, .doc, .pdf, .rtf, etc...
	supporting multiple file formats can be done with apache tika
*/

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sort"

	"github.com/dannav/starship/services/internal/shared"

	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/embedding"
	annoyindex "github.com/dannav/starship/services/internal/platform/spotify"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/dannav/starship/services/internal/store"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"gopkg.in/neurosnap/sentences.v1/english"
)

// vDims represents how many dimensions a spotify/annoy index vector is (sized returned by tensorflow model)
var vDims = 512

// MimeToDocType translates a tika mimeType to a document type
var MimeToDocType = map[string]int{
	"text/plain":      document.TypeMarkdown,
	"application/pdf": document.TypePDF,
}

// GetDocType gets the document type from a mimeType
func GetDocType(s string) int {
	if _, ok := MimeToDocType[s]; !ok {
		return document.TypeUnsupported
	}

	return MimeToDocType[s]
}

// Index handles creating word embeddings from a multi-parse/file upload and indexing it for search purposes
func (a *App) Index(ds *document.Service, ss *store.Service) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file, header, err := r.FormFile("content")
		defer file.Close()

		if err != nil {
			err = errors.Wrap(err, "read multi part file")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// TODO store the file in block storage so we can download from a search

		// pass file to apache tika for parsing to plaintext
		tikaResp, err := a.TikaParse(file, header.Filename)
		if err != nil {
			err = errors.Wrap(err, "contents parse failed")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}
		content := tikaResp.Body

		// break up text into sentences
		var inputs []string
		tokenizer, err := english.NewSentenceTokenizer(nil)
		if err != nil {
			err = errors.Wrap(err, "tokenizing text to sentences")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		sentences := tokenizer.Tokenize(content)
		for _, s := range sentences {
			inputs = append(inputs, s.Text)
		}

		// generate word embeddings from given text
		es, err := embedding.Generate(inputs, a.ServingURL, a.HTTPClient)
		if err != nil {
			err = errors.New("cannot generate embeddings from text")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}

		// create document
		teamID := "1" // TODO replace with user auth
		d := &document.Document{
			Body:   content, // we store the content so we can do full text search with postgres
			Name:   header.Filename,
			TypeID: GetDocType(tikaResp.DocumentType),
			TeamID: teamID,
		}

		d, err = ds.CreateDocument(d)
		if err != nil {
			err = errors.Wrap(err, "create document")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// create a new annoy index to load data into
		t := annoyindex.NewAnnoyIndexAngular(vDims)
		defer t.Unload()

		// create new store for this team if it does not exist
		s := &store.Store{
			Location: fmt.Sprintf("indexes/%v.ann", teamID),
			TeamID:   teamID,
		}

		st, foundStore, err := ss.CreateStoreIfNotExists(s)
		if err != nil {
			err = errors.Wrap(err, "creating annoy store")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// add newly indexed data into annoy and db
		for i := range inputs {
			emb, err := json.Marshal(es[i])
			if err != nil {
				err = errors.Wrapf(err, "converting embedding to json, embedding %v", i)
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}

			// build context of sentence (include surrounding sentences)
			context := inputs[i]
			if i != 0 {
				context = fmt.Sprintf("%v%v", inputs[i-1], context)
			}

			if len(inputs) > i+1 {
				context = fmt.Sprintf("%v%v", context, inputs[i+1])
			}

			sen := &document.Sentence{
				Body:       inputs[i],
				Context:    context,
				DocumentID: d.ID,
				StoreID:    st.ID,
				Embedding:  emb,
			}

			sen, err = ds.CreateSentence(sen)
			if err != nil {
				err := errors.Wrap(err, "create document index")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}

			// setting the key as AnnoyID allows us to map back to a sentence later during search
			t.AddItem(sen.AnnoyID, es[i])
		}

		// previous store exists so we should load all previously indexed data into annoy
		if foundStore {
			ic, err := ds.GetIndexContentForTeam(teamID)
			if err != nil {
				err = errors.Wrap(err, "getting index content for team")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}

			for _, c := range ic {
				emb, err := c.GetEmbeddings()
				if err != nil {
					err = errors.Wrap(err, "converting embedding json to slice")
					web.RespondError(w, r, http.StatusInternalServerError, err)
					return
				}

				t.AddItem(c.AnnoyID, emb)
			}
		}

		// build index with N trees, more trees gives higher precision when querying
		t.Build(10)
		if s := t.Save(st.Location); !s {
			err := errors.Errorf("failed to save index at location %v", st.Location)
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		web.Respond(w, r, http.StatusNoContent, nil)
		return
	}
}

// Search performs a search on all documents with the given text
func (a *App) Search(ds *document.Service, ss *store.Service) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var b struct {
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			err = errors.Wrap(err, "decode request body")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// validation
		if len(b.Text) == 0 {
			err := errors.New("text field is required and must not be empty")
			web.RespondError(w, r, http.StatusBadRequest, err)
			return
		}

		// generate word embeddings from given text
		searchText := []string{b.Text}
		es, err := embedding.Generate(searchText, a.ServingURL, a.HTTPClient)
		if err != nil {
			err = errors.New("cannot generate embeddings from text")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}

		// TODO :- when authentication is done use users teamID
		teamID := "1"
		s, err := ss.GetStoreByTeamID(teamID)
		if err != nil {
			err = errors.Wrap(err, "get store by teamID")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// check if file exists at store location
		if _, err := os.Stat(s.Location); os.IsNotExist(err) {
			err := errors.Errorf("file `%v` does not exist", s.Location)
			web.RespondError(w, r, http.StatusBadRequest, err)
			return
		}

		// create and load index on disk into memory
		t := annoyindex.NewAnnoyIndexAngular(vDims)
		defer t.Unload()

		if success := t.Load(s.Location); !success {
			err := errors.Errorf("could not load file %v", s.Location)
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// perform KNN search on embedding for store. annoy index ids are mapped to sentence ids
		var sIDs []int
		var distances []float32
		t.GetNnsByVector(es[0], 20, 1, &sIDs, &distances) // embedding, # items, bool distances?, annoykeys, distances

		// sort sIDs by closest distance
		sort.Slice(sIDs, func(i, j int) bool {
			return distances[i] < distances[j]
		})

		// sort distance to match order of sIDs
		sort.Slice(distances, func(i, j int) bool {
			return distances[i] < distances[j]
		})

		// get documents and sentences from above IDs
		docs, err := ds.GetSearchResults(sIDs)
		if err != nil {
			err = errors.Wrap(err, "get search results")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// sort docs in order of best match
		var docsSorted []document.SearchResult
		for i, sid := range sIDs {
			for _, d := range docs {
				if d.AnnoyID == sid {
					d.Rank = distances[i] // set rank to distance
					docsSorted = append(docsSorted, d)
				}
			}
		}

		ftsDocs, err := ds.FullTextSearch(b.Text)
		if err != nil {
			err = errors.Wrap(err, "get search results")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// Ranking System Process (TODO run tests and tweak this for best results)
		// =======================================================================

		// 1. Start with distance from searchtext for annoy search results (set above when creating docsSorted)

		// ultimately combine fts results and annoy search results
		uniqResults := map[string]bool{}
		docs = []document.SearchResult{}

		// 2. If sentence_id or document_id exists in fts subtract the relevancy in fts from distance
		for _, d := range docsSorted {
			for _, f := range ftsDocs {
				if d.SentenceID.String() == f.SentenceID.String() || d.DocumentID.String() == f.DocumentID.String() {
					if _, ok := uniqResults[d.SentenceID.String()]; ok {
						continue
					} else {
						uniqResults[d.SentenceID.String()] = true
					}

					d.Rank = d.Rank - f.Rank
				}
			}

			docs = append(docs, d)
		}

		// 3. use 1 - relevancy as rank for fts results
		for _, f := range ftsDocs {
			if _, ok := uniqResults[f.SentenceID.String()]; ok {
				continue
			} else {
				uniqResults[f.SentenceID.String()] = true
			}

			f.Rank = 1 - f.Rank
			docs = append(docs, f)
		}

		// Finally, order by lowest rank first
		sort.Slice(docs, func(i, j int) bool {
			return docs[i].Rank < docs[j].Rank
		})

		result := struct {
			Distances []float32               `json:"distances"`
			Results   []document.SearchResult `json:"results"`
		}{
			Distances: distances,
			Results:   docs,
		}

		web.Respond(w, r, http.StatusOK, result)
	}
}

// TikaParse makes a request to the tikad endpoint to parse a file
func (a *App) TikaParse(file io.Reader, filename string) (*shared.TikaResponse, error) {
	var parseResp struct {
		Results shared.TikaResponse `json:"results"`
	}

	var buf bytes.Buffer
	encoder := multipart.NewWriter(&buf)
	field, err := encoder.CreateFormFile("content", filename)
	if err != nil {
		err = errors.Wrap(err, "creating content form field for tikad request")
		return nil, err
	}

	_, err = io.Copy(field, file)
	if err != nil {
		err = errors.Wrap(err, "copying file to tikad request")
		return nil, err
	}
	encoder.Close()

	endpoint := fmt.Sprintf("http://%v/v1/parse", a.Cfg.TikaURL)
	req, err := http.NewRequest(http.MethodPost, endpoint, &buf)
	if err != nil {
		err = errors.Wrap(err, "preparing tikad request")
		return nil, err
	}
	req.Header.Set("Content-Type", encoder.FormDataContentType())

	res, err := a.HTTPClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing tikad request")
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("tikad request failed")
		return nil, err
	}

	// get response from tikad
	if err := json.NewDecoder(res.Body).Decode(&parseResp); err != nil {
		err = errors.Wrap(err, "decoding tikad response")
		return nil, err
	}

	return &parseResp.Results, nil
}
