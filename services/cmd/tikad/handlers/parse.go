package handlers

import (
	"bytes"
	"context"
	"net/http"

	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/dannav/starship/services/internal/shared"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	stripmd "github.com/writeas/go-strip-markdown"
)

// Parse an uploaded file and return results from apache tika
func (a *App) Parse() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file, _, err := r.FormFile("content")
		if err != nil {
			err = errors.Wrap(err, "read multi part file")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}
		defer file.Close()
		mimeType, err := a.Tika.Detect(context.Background(), file)
		if err != nil {
			err = errors.Wrap(err, "tika detect mime")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}
		file.Seek(0, 0)

		switch mimeType {
		case "application/octet-stream":
			web.RespondError(w, r, http.StatusUnsupportedMediaType, errors.New("unknown mime type"))
			return
		case "text/plain":
			// markdown will register as text/plain mimeType
			// plain text or markdown does not need to be parsed by apache tika
			buf := new(bytes.Buffer)
			buf.ReadFrom(file)

			s := buf.String()
			body := stripmd.Strip(s)

			resp := shared.TikaResponse{
				Body:         body,
				DocumentType: mimeType,
			}

			web.Respond(w, r, http.StatusOK, resp)
			return
		}

		body, err := a.Tika.Parse(context.Background(), file)
		if err != nil {
			err = errors.Wrap(err, "tika parse")
			web.RespondError(w, r, http.StatusInternalServerError, err)
			return
		}

		resp := shared.TikaResponse{
			Body:         body,
			DocumentType: mimeType,
		}

		web.Respond(w, r, http.StatusOK, resp)
	}
}
