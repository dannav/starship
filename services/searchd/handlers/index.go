package handlers

import (
	"net/http"

	"starship/services/internal/platform/web"

	"github.com/julienschmidt/httprouter"
)

// Index handles creating word embeddings from a string and storing it for search purposes
func (a *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	web.Respond(w, r, http.StatusOK, "Hello World")
}
