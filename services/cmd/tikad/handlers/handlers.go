package handlers

import (
	"net/http"

	"github.com/google/go-tika/tika"

	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq" // pq is the postgresql driver configuration
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// App represents the application and configuration
type App struct {
	handler http.Handler
	Tika    *tika.Client
}

// NewApp returns a new instance of App with config and DB connections loaded
func NewApp(t *tika.Client) *App {
	a := App{
		Tika: t,
	}

	a.initHandler()

	return &a
}

// ServeHTTP satisfies the http handler interface so App can perform as an http server
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}

// initHandler creates a new router and initializes all the routes
func (a *App) initHandler() {
	r := httprouter.New()

	// handle 404 Errors
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		web.RespondError(w, r, http.StatusNotFound, errors.New("not found"))
	})

	// gracefully handle and recover from web panics
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.WithFields(log.Fields{
			"error": i,
		}).Error("panic")

		web.RespondError(w, r, http.StatusInternalServerError, web.ErrInternalServer)
	}

	// api routes
	r.Handle(http.MethodGet, "/ready", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		web.Respond(w, r, http.StatusOK, nil)
	})

	r.POST("/v1/parse", a.Parse())

	// wrap all the routes with global middleware
	a.handler = web.RequestMW(r)
}
