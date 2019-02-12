package handlers

import (
	"net/http"

	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

var errNotAuthorized = errors.New("not authorized")

// IndexAuthorized checks to see if an indexKey was supplied at app start and ensures the connecting client has provided it
// in the proper header
func IndexAuthorized(next httprouter.Handle, indexKey string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if indexKey != "" {
			key := r.Header.Get("X-INDEX-KEY")

			if indexKey != key {
				web.RespondError(w, r, http.StatusUnauthorized, errNotAuthorized)
				return
			}
		}

		next(w, r, p)
	}
}

// AccessAuthorized checks to see if an accessKey was supplied at app start and ensures the connecting client has provided it
// in the proper header
func AccessAuthorized(next httprouter.Handle, accessKey string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if accessKey != "" {
			key := r.Header.Get("X-ACCESS-KEY")

			if accessKey != key {
				web.RespondError(w, r, http.StatusUnauthorized, errNotAuthorized)
				return
			}
		}

		next(w, r, p)
	}
}
