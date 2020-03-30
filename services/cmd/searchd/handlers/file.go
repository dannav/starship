package handlers

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	minio "github.com/minio/minio-go"

	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

// RootFolder represents the root folder that documents for starship are stored on object storage
// TODO - have this configurable in an environment variable
const RootFolder = "starship_documents"

// DownloadFile handles downloading a file from object storage
func (a *App) DownloadFile() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file := RootFolder + r.URL.Query().Get("file")

		// get filename by copying everything after last "/" in path
		filename := file[strings.LastIndex(file, "/")+1:]

		// download from object storage
		if a.ObjectStorageEnabled {
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

			return
		}

		// if not using object storage load from disk
		file = filepath.Join(a.StoragePath, file)

		// write headers needed when allowing client to download file
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)

		// serve file rejects any requests that contain '..' so a user could not try to download outside of the starship
		// storage dir
		http.ServeFile(w, r, file)
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

// getFileContentType checks the content type of a file using the http package
func getFileContentType(f *os.File) (string, error) {
	// only the first 512 bytes are used to sniff the content type
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	// return content type or "application/octet-stream" if not match is found
	return http.DetectContentType(buffer), nil
}
