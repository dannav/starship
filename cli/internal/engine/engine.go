package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var errSearch = errors.New("there was an error performing your search")
var errDownloadingFile = errors.New("could not download file")

// Engine represents the search engine
type Engine struct {
	Client      *http.Client
	APIEndpoint string
}

// SearchResponse represents the structure of a returned search query
type SearchResponse struct {
	Distances []float32      `json:"distances"`
	Documents []SearchResult `json:"documents"`
}

// SearchResult represents search results
type SearchResult struct {
	DocumentID   uuid.UUID `json:"id"`
	SentenceID   uuid.UUID `json:"sentenceID"`
	AnnoyID      int       `json:"annoyID"`
	DocumentName string    `json:"name"`
	DownloadURL  string    `json:"downloadURL"`
	Text         string    `json:"text"`
	Rank         float32   `json:"rel"`
}

// Response is the format used for all api responses
type Response struct {
	Results interface{}     `json:"results"`
	Errors  []ResponseError `json:"errors,omitempty"`
}

// ResponseError is the format used for api response errors
type ResponseError struct {
	Message string `json:"message"`
}

// Error implements the error interface
func (a ResponseError) Error() string {
	return a.Message
}

// WrapAPIErrors transforms an api error response into one wrapped error
func WrapAPIErrors(err error, rErrs []ResponseError) error {
	var es []string
	for _, e := range rErrs {
		es = append(es, e.Error())
	}

	rErrors := strings.Join(es, ",")
	return errors.Wrap(err, rErrors)
}

// NewEngine creates a new engine
func NewEngine(client *http.Client) Engine {
	return Engine{
		Client:      client,
		APIEndpoint: "http://localhost:8080",
	}
}

// DownloadFile attempts to download a file from the api
func (e *Engine) DownloadFile(url string) error {
	baseURL := "stuph.sfo2.digitaloceanspaces.com/"

	// create file on disk and extract filename
	filename := url[strings.LastIndex(url, "/")+1:]
	out, err := os.Create(filename)
	defer out.Close()

	// create file url query to download
	fileWithTeam := strings.Replace(url, baseURL, "", 1)
	file := fileWithTeam[strings.Index(fileWithTeam, "/")-1:]

	endpoint := fmt.Sprintf("%v/v1/download?file=%v", e.APIEndpoint, file)
	res, err := http.Get(endpoint)
	if err != nil {
		err = errors.Wrap(err, "downloading file")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode, endpoint)
		return errDownloadingFile
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		err = errors.Wrap(err, "writing file to disk")
		return err
	}

	return nil
}

// Search returns all search results for a query
func (e *Engine) Search(text string) (*SearchResponse, error) {
	rBody := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}

	b, err := json.Marshal(rBody)
	if err != nil {
		err = errors.Wrap(err, "marshaling search body")
		return nil, err
	}

	endpoint := fmt.Sprintf("%v/v1/search", e.APIEndpoint)
	req, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(b))
	if err != nil {
		err = errors.Wrap(err, "preparing search request")
		return nil, err
	}

	res, err := e.Client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing search request")
		return nil, err
	}

	var body SearchResponse
	r := Response{
		Results: &body,
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.Wrap(err, "decoding search response body")
	}

	if res.StatusCode != http.StatusOK {
		err := WrapAPIErrors(errSearch, r.Errors)
		return nil, err
	}

	return &body, nil
}

// Index stores and indexes a file with the API
func (e *Engine) Index(file io.Reader, filename string) error {
	var buf bytes.Buffer

	encoder := multipart.NewWriter(&buf)
	field, err := encoder.CreateFormFile("content", filename)
	if err != nil {
		err = errors.Wrap(err, "creating content form field for index request")
		return err
	}

	_, err = io.Copy(field, file)
	if err != nil {
		err = errors.Wrap(err, "copying file to index request")
		return err
	}
	encoder.Close()

	endpoint := fmt.Sprintf("%v/v1/index", e.APIEndpoint)
	req, err := http.NewRequest(http.MethodPost, endpoint, &buf)
	if err != nil {
		err = errors.Wrap(err, "preparing index request")
		return err
	}

	req.Header.Set("Content-Type", encoder.FormDataContentType())
	res, err := e.Client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing idnex request")
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		err = errors.New("index request failed")
		return err
	}

	return nil
}