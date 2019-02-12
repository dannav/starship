package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	// accessKeyHeader is the header to set an access key
	accessKeyHeader = "X-ACCESS-KEY"

	// indexKeyHeader is the header to set an index key
	indexKeyHeader = "X-INDEX-KEY"
)

var (
	// errSearch is the error returned a search request could not be made
	errSearch = errors.New("there was an error performing your search")

	// errDownloadingFile is the error returned when a request to download a file could not be made
	errDownloadingFile = errors.New("could not download file")

	// ErrUnauthorized is the error returned when an unauthorized request is made
	ErrUnauthorized = errors.New("you are not authorized to perform this command against the connected server")
)

// Engine represents the search engine
type Engine struct {
	Client      *http.Client
	APIEndpoint string
	IndexKey    string
	AccessKey   string
}

// SearchResponse represents the structure of a returned search query
type SearchResponse struct {
	Distances []float32      `json:"distances"`
	Documents []SearchResult `json:"documents"`
}

// SearchResult represents search results
type SearchResult struct {
	Path         string  `json:"path"`
	DocumentName string  `json:"name"`
	DownloadURL  string  `json:"downloadURL"`
	Text         string  `json:"text"`
	Rank         float32 `json:"rel"`
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
func NewEngine(client *http.Client, indexKey, accessKey string) *Engine {
	return &Engine{
		Client:    client,
		IndexKey:  indexKey,
		AccessKey: accessKey,
	}
}

// Ready checks to see if the API is up and ready
func (e *Engine) Ready() (bool, error) {
	endpoint := fmt.Sprintf("%v/ready", e.APIEndpoint)
	u, err := url.Parse(endpoint)
	if err != nil {
		return false, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, errors.New("starship api is down")
	}

	return true, nil
}

// SetHeader sets a given header with a value on a request
func SetHeader(r *http.Request, header, value string) {
	r.Header.Set(header, value)
}

// DownloadFile attempts to download a file from the api
func (e *Engine) DownloadFile(url string) error {
	// create file on disk and extract filename
	filename := url[strings.LastIndex(url, "/")+1:]
	out, err := os.Create(filename)
	defer out.Close()

	endpoint := fmt.Sprintf("%v/v1/download?file=%v", e.APIEndpoint, url)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		err = errors.Wrap(err, "preparing download request")
		return err
	}

	if e.AccessKey != "" {
		SetHeader(req, accessKeyHeader, e.AccessKey)
	}

	res, err := e.Client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "downloading file")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return ErrUnauthorized
		}

		return errDownloadingFile
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		err = errors.Wrap(err, "writing file to disk")
		return err
	}

	return nil
}

// ExistsAtIndexPath checks if a file already exists at this index path
func (e *Engine) ExistsAtIndexPath(path string) (bool, error) {
	endpoint := fmt.Sprintf("%v/v1/exists?path=%v", e.APIEndpoint, path)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		err = errors.Wrap(err, "preparing exists request")
		return false, err
	}

	if e.AccessKey != "" {
		SetHeader(req, accessKeyHeader, e.AccessKey)
	}

	res, err := e.Client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing exists request")
		return false, err
	}

	var exists bool
	r := Response{
		Results: &exists,
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return false, errors.Wrap(err, "decoding exists response body")
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return false, ErrUnauthorized
		}

		err := WrapAPIErrors(errors.New("error checking index path existence"), r.Errors)
		return false, err
	}

	return exists, nil
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

	if e.AccessKey != "" {
		SetHeader(req, accessKeyHeader, e.AccessKey)
	}

	res, err := e.Client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing search request")
		return nil, err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, ErrUnauthorized
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
func (e *Engine) Index(filePath, filename, indexPath string) error {
	buf := &bytes.Buffer{}
	encoder := multipart.NewWriter(buf)

	pathField, err := encoder.CreateFormField("path")
	if err != nil {
		err = errors.Wrap(err, "creating path form field for index request")
		return err
	}

	_, err = pathField.Write([]byte(indexPath))
	if err != nil {
		err = errors.Wrap(err, "writing index path to index request")
		return err
	}

	contentField, err := encoder.CreateFormFile("content", filename)
	if err != nil {
		err = errors.Wrap(err, "creating content form field for index request")
		return err
	}

	// start reading the file
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// copy file into request body for upload
	_, err = io.Copy(contentField, f)
	if err != nil {
		err = errors.Wrap(err, "copying file to index request")
		return err
	}
	encoder.Close()

	endpoint := fmt.Sprintf("%v/v1/index", e.APIEndpoint)
	req, err := http.NewRequest(http.MethodPost, endpoint, buf)
	if err != nil {
		err = errors.Wrap(err, "preparing index request")
		return err
	}
	SetHeader(req, "Content-Type", encoder.FormDataContentType())

	if e.IndexKey != "" {
		SetHeader(req, indexKeyHeader, e.IndexKey)
	}

	res, err := e.Client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing index request")
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		if res.StatusCode == http.StatusUnsupportedMediaType {
			err = errors.New("indexing for the given file type is not supported")
		} else if res.StatusCode == http.StatusUnauthorized {
			err = ErrUnauthorized
		} else {
			err = errors.New("index request failed")
		}

		return err
	}

	return nil
}
