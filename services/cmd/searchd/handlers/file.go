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

// TODO - have this configurable in an environment variable
// RootFolder represents the root folder that documents for starship are stored on object storage
const RootFolder = "starship_documents"

// DownloadFile handles downloading a file from object storage
func (a *App) DownloadFile() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file := RootFolder + r.URL.Query().Get("file")
		client, err := minio.New(a.ObjectStorageConfig.URL, a.ObjectStorageConfig.Key, a.ObjectStorageConfig.Secret, true)
		if err != nil {
			err = errors.Wrap(err, "connecting to object storage, is the config provided correct?")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		opts := minio.GetObjectOptions{}
		o, err := client.GetObject(a.ObjectStorageConfig.BucketName, file, opts)
		if err != nil && o != nil {
			err = errors.Wrap(err, "downloading file from object storage")
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
