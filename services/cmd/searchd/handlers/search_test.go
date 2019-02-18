package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dannav/starship/services/cmd/searchd/handlers"
	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/pkg/errors"
)

// TestIndexAndSearchWithPaths tests indexing a document and searching for it
// it also ensures that the path created by the index handler is correct
func TestIndexAndSearchWithPaths(t *testing.T) {
	type testContent struct {
		Content            string
		Filename           string
		Path               string
		ExpectedParsedPath string
		SearchQuery        string
	}

	tt := []testContent{
		// test valid index and search
		testContent{
			Content:            "Hello World!",
			Filename:           "test.md",
			Path:               "/",
			ExpectedParsedPath: "_rootfolder_.test.md",
			SearchQuery:        "hello",
		},
		// test a nested path with trailing '/'
		testContent{
			Content:            "Lorem ipsum dolor sit amet!",
			Filename:           "ipsum.txt",
			Path:               "/project/lipsum/",
			ExpectedParsedPath: "_rootfolder_.project.lipsum.ipsum.txt",
			SearchQuery:        "lorem",
		},
		// test no path defined (should default to root folder)
		testContent{
			Content:            "file with no path defined on index",
			Filename:           "nopath.txt",
			Path:               "",
			ExpectedParsedPath: "_rootfolder_.nopath.txt",
			SearchQuery:        "no path",
		},
		// test nested path with no trailing '/'
		testContent{
			Content:            "Lorem ipsum dolor sit amet!",
			Filename:           "notrailing.txt",
			Path:               "/anotherproject/lipsum",
			ExpectedParsedPath: "_rootfolder_.anotherproject.lipsum.notrailing.txt",
			SearchQuery:        "lorem",
		},
		// TODO :- make modifications to make this test pass
		// test weird characters in path
		// testContent{
		// 	Content:            "Lorem ipsum dolor sit amet!",
		// 	Filename:           "notrailing.txt",
		// 	Path:               "/_()$hello/.*&^/valid",
		// 	ExpectedParsedPath: "_rootfolder_.anotherproject.lipsum.notrailing.txt",
		// 	SearchQuery:        "lorem",
		// },
	}

	for _, test := range tt {
		// index test content
		err := index(test.Content, test.Filename, test.Path)
		if err != nil {
			t.Fatal(err)
		}

		sr, err := search(test.SearchQuery)
		if err != nil {
			t.Fatal(err)
		}

		if len(sr.Documents) > 0 {
			var d *document.SearchResult
			for _, srd := range sr.Documents {
				if srd.DocumentName == test.Filename {
					d = &srd
				}
			}

			if d == nil {
				t.Error("could not find document in search results")
				continue
			}

			if d.Path != test.ExpectedParsedPath {
				t.Errorf("expected parsed path %v, got %v", test.ExpectedParsedPath, d.Path)
			}
		}
	}
}

// search handles making a search request
func search(text string) (*handlers.SearchResponse, error) {
	s := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	endpoint := ("/v1/search")
	req, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(b))
	if err != nil {
		err = errors.Wrap(err, "preparing search request")
		return nil, err
	}

	w := httptest.NewRecorder()
	ts.App.ServeHTTP(w, req)

	var res handlers.SearchResponse
	d := web.Response{
		Results: &res,
	}

	if w.Code != http.StatusOK {
		return nil, errors.New("received error performing search")
	}

	if err := json.NewDecoder(w.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal json: %v")
	}

	return &res, nil
}

// index handles making an index request
func index(content, filename, path string) error {
	buf := &bytes.Buffer{}
	encoder := multipart.NewWriter(buf)

	contentField, err := encoder.CreateFormFile("content", filename)
	if err != nil {
		err = errors.Wrap(err, "creating content form field for index request")
		return err
	}

	// create a reader from string to simulate a file open
	f := strings.NewReader(content)

	// copy file into request body for upload
	_, err = io.Copy(contentField, f)
	if err != nil {
		err = errors.Wrap(err, "copying file to index request")
		return err
	}
	encoder.Close()

	endpoint := ("/v1/index")
	req, err := http.NewRequest(http.MethodPost, endpoint, buf)
	if err != nil {
		err = errors.Wrap(err, "preparing index request")
		return err
	}
	req.Header.Set("Content-Type", encoder.FormDataContentType())
	req.Header.Set("X-PATH", path)

	w := httptest.NewRecorder()
	ts.App.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		if w.Code == http.StatusUnsupportedMediaType {
			err = errors.New("indexing for the given file type is not supported")
		} else if w.Code == http.StatusUnauthorized {
			err = errors.New("not authorized")
		} else {
			err = errors.New("index request failed")
		}

		return err
	}

	return nil
}
