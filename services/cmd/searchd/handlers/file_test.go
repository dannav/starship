package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/pkg/errors"
)

// TestFileExists tests the exists endpoint
func TestFileExists(t *testing.T) {
	err := index("hello world", "exists.md", "/")
	if err != nil {
		t.Error(err)
	}

	e, err := exists("/exists.md")
	if err != nil {
		t.Error(err)
	}

	if e != true {
		t.Errorf("expected %v, got %v", true, e)
	}

	e, err = exists("/random/file/that/should/not/exist.md")
	if err != nil {
		t.Error(err)
	}

	if e != false {
		t.Errorf("expected %v, got %v", false, e)
	}
}

// exists handles making a request to the exists endpoint
func exists(path string) (bool, error) {
	endpoint := "/v1/exists?path=" + path
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		err = errors.Wrap(err, "preparing exists request")
		return false, err
	}

	w := httptest.NewRecorder()
	ts.App.ServeHTTP(w, req)

	var res bool
	d := web.Response{
		Results: &res,
	}

	if w.Code != http.StatusOK {
		return false, errors.New("received error performing exists request")
	}

	if err := json.NewDecoder(w.Body).Decode(&d); err != nil {
		return false, errors.Wrap(err, "unable to unmarshal json")
	}

	return res, nil
}
