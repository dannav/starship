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

// TikaParse makes a request to the tikad endpoint to parse a file
func (a *App) TikaParse(file io.Reader, filename string) (string, error) {
	var parseResp struct {
		Results string `json:"results"`
	}

	var buf bytes.Buffer
	encoder := multipart.NewWriter(&buf)
	field, err := encoder.CreateFormFile("content", filename)
	if err != nil {
		err = errors.Wrap(err, "creating content form field for tikad request")
		return "", err
	}

	_, err = io.Copy(field, file)
	if err != nil {
		err = errors.Wrap(err, "copying file to tikad request")
		return "", err
	}
	encoder.Close()

	endpoint := fmt.Sprintf("http://%v/v1/parse", a.Cfg.TikaURL)
	req, err := http.NewRequest(http.MethodPost, endpoint, &buf)
	if err != nil {
		err = errors.Wrap(err, "preparing tikad request")
		return "", err
	}
	req.Header.Set("Content-Type", encoder.FormDataContentType())

	res, err := a.HTTPClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing tikad request")
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("tikad request failed")
		return "", err
	}

	// get response from tikad
	if err := json.NewDecoder(res.Body).Decode(&parseResp); err != nil {
		err = errors.Wrap(err, "decoding tikad response")
		return "", err
	}

	return parseResp.Results, nil
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

		// pass file to apache tika for parsing to plaintext
		content, err := a.TikaParse(file, header.Filename)
		if err != nil {
			err = errors.Wrap(err, "contents parse failed")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}

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
			Body:   content, // TODO we shouldn't store the body but link to a file on cloud storage
			Name:   header.Filename,
			TypeID: document.TypeText,
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
			Location: fmt.Sprintf("%v.ann", teamID),
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

			sen := &document.Sentence{
				Body:       inputs[i],
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
		for _, sid := range sIDs {
			for _, d := range docs {
				if d.AnnoyID == sid {
					docsSorted = append(docsSorted, d)
				}
			}
		}

		result := struct {
			Distances []float32               `json:"distances"`
			Results   []document.SearchResult `json:"results"`
		}{
			Distances: distances,
			Results:   docsSorted,
		}

		web.Respond(w, r, http.StatusOK, result)
	}
}
