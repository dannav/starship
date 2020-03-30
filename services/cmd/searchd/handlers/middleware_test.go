package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dannav/starship/services/cmd/searchd/handlers"
	"github.com/julienschmidt/httprouter"
)

// TestIndexAuthorized tests the IndexAuthorized middleware
func TestIndexAuthorized(t *testing.T) {
	r := httprouter.New()
	r.GET("/test", handlers.IndexAuthorized(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}, "XXX"))

	// successful request
	{
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		req.Header.Set("X-INDEX-KEY", "XXX")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if e, a := http.StatusOK, w.Code; e != a {
			t.Errorf("expected code: %v, got code: %v", e, a)
		}
	}

	// should 401 if a valid index key header is not set
	{
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if e, a := http.StatusUnauthorized, w.Code; e != a {
			t.Errorf("expected code: %v, got code: %v", e, a)
		}
	}

	// should 401 if an invalid index key header is set
	{
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		req.Header.Set("X-INDEX-KEY", "NOTRIGHT")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if e, a := http.StatusUnauthorized, w.Code; e != a {
			t.Errorf("expected code: %v, got code: %v", e, a)
		}
	}
}

// TestAccessAuthorized tests the AccessAuthorized middleware
func TestAccessAuthorized(t *testing.T) {
	r := httprouter.New()
	r.GET("/test", handlers.AccessAuthorized(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}, "YYY"))

	// successful request
	{
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		req.Header.Set("X-ACCESS-KEY", "YYY")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if e, a := http.StatusOK, w.Code; e != a {
			t.Errorf("expected code: %v, got code: %v", e, a)
		}
	}

	// should 401 if a valid access key header is not set
	{
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if e, a := http.StatusUnauthorized, w.Code; e != a {
			t.Errorf("expected code: %v, got code: %v", e, a)
		}
	}

	// should 401 if an invalid access key header is set
	{
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		req.Header.Set("X-ACCESS-KEY", "NOTRIGHT")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if e, a := http.StatusUnauthorized, w.Code; e != a {
			t.Errorf("expected code: %v, got code: %v", e, a)
		}
	}
}
