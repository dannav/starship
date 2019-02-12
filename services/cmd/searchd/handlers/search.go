package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/dannav/starship/services/internal/fileop"
	"github.com/dannav/starship/services/internal/shared"
	"github.com/google/uuid"
	minio "github.com/minio/minio-go"

	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/embedding"
	annoyindex "github.com/dannav/starship/services/internal/platform/spotify"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/dannav/starship/services/internal/store"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	prose "gopkg.in/jdkato/prose.v2"
)

// ******
// TODO - make loading of indexes faster - we can always store the index in memory so search and index functions
// do not need to reload it from disk every time a request is made
// ******

// MimeToDocType translates a tika mimeType to a document type
var MimeToDocType = map[string]int{
	"text/plain":      document.TypeMarkdown,
	"application/pdf": document.TypePDF,
}

// errUnkownMimeType is an error returned from apache tika if trying to parse a file that is not text
var errUnknownMimeType = errors.New("tikad unknown media type")

// GetDocType gets the document type from a mimeType
func GetDocType(s string) int {
	if _, ok := MimeToDocType[s]; !ok {
		return document.TypeUnsupported
	}

	return MimeToDocType[s]
}

// Index handles creating word embeddings from a multi-part/file upload and indexing it for search purposes
func (a *App) Index(ds *document.Service, ss *store.Service) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file, header, err := r.FormFile("content")
		defer file.Close()

		if err != nil {
			err = errors.Wrap(err, "read multi part file")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		path := r.Header.Get("X-PATH")
		if path == "" {
			path = "/"
		}

		// TODO - replace with filepath.Clean() and test
		// if path is multiple dirs remove any possible trailing or repeating '/'
	cleanpath:
		if path != "/" {
			for i := range path {
				runeArr := []rune(path)
				ch := runeArr[i]

				if ch == '/' && i == len(path)-1 {
					path = path[:i] + path[i+1:]
					goto cleanpath
				}

				if ch == '/' && len(path) >= i+1 {
					if ch := runeArr[i+1]; ch == '/' {
						path = path[:i] + path[i+1:]
						goto cleanpath
					}
				}
			}
		}

		// if first char is '/' and have multiple dirs remove leading '/'
		if strings.Count(path, "/") > 1 && path[0] == '/' {
			path = strings.Replace(path, "/", "", 1)
		}

		// pass file to apache tika for parsing to plaintext
		tikaResp, err := a.TikaParse(file, header.Filename)
		if err != nil {
			if err := errors.Cause(err); err == errUnknownMimeType {
				web.RespondError(w, r, http.StatusUnsupportedMediaType, err)
				return
			}

			err = errors.Wrap(err, "failed to extract text from document")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}
		content := tikaResp.Body

		// create regex for non alphanumeric chars, ignore symbols commonly found in text
		reg, err := regexp.Compile(`[^a-zA-Z0-9\s\.!?:;"'\$@&]+`)
		if err != nil {
			err = errors.Wrap(err, "non alpha regex failed to compile")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// create regex for new lines on all systems \r is to catch newlines on windows
		regNewLines, err := regexp.Compile(`\r?\n`)
		if err != nil {
			err = errors.Wrap(err, "newline regex failed to compile")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// convert tika content resp to a prose document and ultimately tokenize the text
		doc, err := prose.NewDocument(content)
		if err != nil {
			err = errors.Wrap(err, "tokenizing text")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// inputs will be sent to the sentence encoder, where uncleaned is stored for search result purposes
		var inputs []string
		var uncleaned []string

		sents := doc.Sentences()
		for _, s := range sents {
			// remove non alpha numeric chars from sentence text
			inputs = append(inputs, reg.ReplaceAllString(s.Text, ""))
			uncleaned = append(uncleaned, regNewLines.ReplaceAllString(s.Text, " "))
		}

		// generate word embeddings from given text from ML model
		es, err := embedding.Generate(inputs, a.ModelURL, a.HTTPClient)
		if err != nil {
			err = errors.Wrap(err, "could not generate embeddings from text")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}

		// set locations for document
		fileLocation := fmt.Sprintf("%v/%v/%v", RootFolder, uuid.New().String(), header.Filename)
		downloadURL := strings.Replace(fileLocation, RootFolder, "", 1)

		// prepare a client to connect to object storage or setup local file path
		var client *minio.Client
		var objectStorageURL, localFilePath string
		if a.ObjectStorageEnabled {
			client, err = minio.New(a.ObjectStorageConfig.URL, a.ObjectStorageConfig.Key, a.ObjectStorageConfig.Secret, true)
			if err != nil {
				err = errors.Wrap(err, "connecting to object storage, is the config provided correct?")
				web.RespondError(w, r, http.StatusServiceUnavailable, err)
				return
			}

			if client == nil {
				err := errors.New("could not instantiate client to connect to object storage")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}
		} else {
			localFilePath = filepath.Join(a.StoragePath, RootFolder, downloadURL)
		}

		if a.ObjectStorageEnabled {
			objectStorageURL = fmt.Sprintf("%v.%v/%v", a.ObjectStorageConfig.BucketName, a.ObjectStorageConfig.URL, fileLocation)
		}

		d := &document.Document{
			Body:             content, // we store the content so we can join FTS search with semantic search results
			Name:             header.Filename,
			TypeID:           GetDocType(tikaResp.DocumentType),
			ObjectStorageURL: objectStorageURL,
			DownloadURL:      downloadURL,
			Path:             path,
		}

		// check if a document at this path exists and delete it so we can update indexes if it does
		pDoc, err := ds.GetDocumentByPath(d.Path)
		if err != nil {
			if err := errors.Cause(err); err != sql.ErrNoRows {
				err = errors.Wrap(err, "get document by path")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		// TODO - add document revisions
		// this document already exists so we're going to delete it
		if pDoc != nil {
			err := ds.DeleteDocument(d.Path)
			if err != nil {
				err = errors.Wrap(err, "delete document")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}

			// delete from object storage or local storage
			if a.ObjectStorageEnabled {
				if err := client.RemoveObject(a.ObjectStorageConfig.BucketName, pDoc.ObjectStorageURL); err != nil {
					err = errors.Wrapf(err, "deleting object from object storage: %v", pDoc.ObjectStorageURL)
					web.RespondError(w, r, http.StatusInternalServerError, err)
					return
				}
			} else {
				err := fileop.DeleteFile(localFilePath)
				if err != nil {
					err = errors.Wrap(err, "using local storage, deleting file")
					web.RespondError(w, r, http.StatusInternalServerError, err)
					return
				}
			}
		}

		// create document
		d, err = ds.CreateDocument(d)
		if err != nil {
			err = errors.Wrap(err, "create document")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// move reader back to beginning of file so we can either upload it to object storage or write it locally
		file.Seek(0, 0)

		// upload file to object storage
		if a.ObjectStorageEnabled {
			opts := minio.PutObjectOptions{
				UserMetadata: map[string]string{},
			}

			_, err = client.PutObjectWithContext(context.Background(), a.ObjectStorageConfig.BucketName, fileLocation, file, header.Size, opts)
			if err != nil {
				err = errors.Wrap(err, "adding object to object storage")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}
		} else {
			err := fileop.WriteFile(file, localFilePath)
			if err != nil {
				err = errors.Wrap(err, "using local storage, writing file")
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		// create a new annoy index to load embeddings into
		t := annoyindex.NewAnnoyIndexAngular(a.ModelVectorDims)
		defer t.Unload()

		// create new store if it does not exist
		s := &store.Store{
			Location: filepath.Join(a.StoragePath, "indexes", "starship.ann"),
		}

		st, foundStore, err := ss.CreateStoreIfNotExists(s)
		if err != nil {
			err = errors.Wrap(err, "creating annoy store")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// add embeddings into annoy and db
		for i := range inputs {
			emb, err := json.Marshal(es[i])
			if err != nil {
				err = errors.Wrapf(err, "converting embedding to json, embedding %v", i)
				web.RespondError(w, r, http.StatusInternalServerError, err)
				return
			}

			// build context of sentence (sentence + surrounding sentences)
			context := uncleaned[i]
			if i != 0 {
				context = fmt.Sprintf("%v %v", uncleaned[i-1], context)
			}

			if len(uncleaned) > i+1 {
				context = fmt.Sprintf("%v %v", context, uncleaned[i+1])
			}

			sen := &document.Sentence{
				Body:       uncleaned[i],
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
			ic, err := ds.GetIndexContent()
			if err != nil {
				err = errors.Wrap(err, "getting index content")
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

		// build index with N trees... more trees gives higher precision when querying
		t.Build(15)
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
		type SearchResponse struct {
			Distances []float32               `json:"distances"`
			Documents []document.SearchResult `json:"documents"`
		}

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
		searchText := []string{strings.ToLower(b.Text)}

		// ml model api should follow tensorflow serving conventions which is where :predict comes form
		es, err := embedding.Generate(searchText, a.ModelURL, a.HTTPClient)
		if err != nil {
			err = errors.New("cannot generate embeddings from text")
			web.RespondError(w, r, http.StatusServiceUnavailable, err)
			return
		}

		s, err := ss.GetStore()
		if err != nil {
			if err := errors.Cause(err); err == sql.ErrNoRows {
				err = errors.Wrap(err, "no documents indexed")
				web.Respond(w, r, http.StatusOK, SearchResponse{})
				return
			}

			err = errors.Wrap(err, "get store")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// check if file exists at store location
		if _, err := os.Stat(s.Location); os.IsNotExist(err) {
			err := errors.Errorf("index file `%v` does not exist, content must be re indexed", s.Location)
			web.RespondError(w, r, http.StatusBadRequest, err)
			return
		}

		// create and load index on disk into memory
		t := annoyindex.NewAnnoyIndexAngular(a.ModelVectorDims)
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

		// Ranking System Process
		// ======================

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

		// 3. use .8 - relevancy as rank for fts results, this ensures that by default fts results are ranked slightly better than semantic results
		for _, f := range ftsDocs {
			if _, ok := uniqResults[f.SentenceID.String()]; ok {
				continue
			} else {
				uniqResults[f.SentenceID.String()] = true
			}

			f.Rank = float32(math.Abs(float64(.8 - f.Rank)))
			docs = append(docs, f)
		}

		// Finally, order by lowest rank first - rank is distance from current search
		// lower is better
		sort.Slice(docs, func(i, j int) bool {
			return docs[i].Rank < docs[j].Rank
		})

		result := SearchResponse{
			Distances: distances,
			Documents: docs,
		}

		web.Respond(w, r, http.StatusOK, result)
	}
}

// TikaParse makes a request to the tikad endpoint to extract the text from a file
func (a *App) TikaParse(file io.Reader, filename string) (*shared.TikaResponse, error) {
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

	endpoint := fmt.Sprintf("%v/v1/parse", a.Cfg.TikaURL)
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
		if res.StatusCode == http.StatusUnsupportedMediaType {
			return nil, errUnknownMimeType
		}

		err = errors.New("tikad request failed")
		return nil, err
	}

	// parse response from tikad
	var parseResp struct {
		Results shared.TikaResponse `json:"results"`
	}

	if err := json.NewDecoder(res.Body).Decode(&parseResp); err != nil {
		err = errors.Wrap(err, "decoding tikad response")
		return nil, err
	}

	return &parseResp.Results, nil
}
