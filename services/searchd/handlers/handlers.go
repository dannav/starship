// Package handlers creates the application struct that handles all routing
// for the prometheus daemon
package handlers

import (
	"errors"
	"log"
	"net/http"

	"starship/services/internal/platform/web"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pq is the postgresql driver configuration

	"github.com/julienschmidt/httprouter"
)

// App represents the application and configuration
type App struct {
	handler http.Handler
	db      *sqlx.DB
}

// NewApp returns a new instance of the with routes loaded.
func NewApp() *App {
	var a App

	// Eventually other things like configuration/etc will be attached to this struct..
	// Initialize them here.

	a.initDB()
	a.initHandler()

	return &a
}

// ServeHTTP is the handler interface to handle serving up the handlers.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}

// initDB initializes the database connection
func (a *App) initDB() {
	a.db = nil
}

// initHandler creates a new router initializes all the routes.
// and wraps them in a global middleware.
func (a *App) initHandler() {
	r := httprouter.New()

	// Handle 404 Errors
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		web.RespondError(w, r, http.StatusNotFound, errors.New("not found"))
	})

	// Gracefully handle and recover from web panics
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Printf("panic:\n %+v", i)

		web.RespondError(w, r, http.StatusInternalServerError, web.ErrInternalServer)
	}

	// API Routes
	r.Handle(http.MethodPost, "/v1/index", a.Index)

	// Wrap all the routes with a global middleware.
	a.handler = web.RequestMW(r)
}
