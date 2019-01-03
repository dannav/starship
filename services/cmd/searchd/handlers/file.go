/*
	Initially only support markdown and text files

	enhance to support .docx, .doc, .pdf, .rtf, etc...
	supporting multiple file formats can be done with apache tika
*/

package handlers

import (
	"database/sql"
	"io"
	"net/http"
	"strings"

	minio "github.com/minio/minio-go"

	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

// DownloadFile handles downloading a file from blob storage
func (a *App) DownloadFile() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file := r.URL.Query().Get("file")
		client, err := minio.New(spacesURL, doKey, doSecret, true)
		if err != nil {
			err = errors.Wrap(err, "connecting to do spaces")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		opts := minio.GetObjectOptions{}
		o, err := client.GetObject("stuph", file, opts)
		if err != nil {
			err = errors.Wrap(err, "downloading file")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		// get filename by copying everything after last "/" in path
		filename := file[strings.LastIndex(file, "/")+1:]
		s, err := o.Stat()

		// write headers needed when allowing client to download file
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", s.ContentType)
		w.Header().Set("Content-Length", string(s.Size))

		// write file to response
		_, err = io.Copy(w, o)
		if err != nil {
			err = errors.Wrap(err, "writing file to response")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}

// CheckFileExistence determines whether a file exists at the given path for an account
func (a *App) CheckFileExistence(ds *document.Service) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// teamID := "1" // TODO replace with user auth

		path := r.URL.Query().Get("path")
		_, err := ds.GetDocumentByPath(path)
		if err != nil {
			if err := errors.Cause(err); err == sql.ErrNoRows {
				web.Respond(w, r, http.StatusOK, false)
				return
			}

			err = errors.Wrap(err, "get document by path")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		web.Respond(w, r, http.StatusOK, true)
	}
}
