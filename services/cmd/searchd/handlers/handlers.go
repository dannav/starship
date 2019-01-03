package handlers

import (
	"net/http"
	"runtime"

	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/dannav/starship/services/internal/store"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq" // pq is the postgresql driver configuration
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Digital Ocean Spaces info
var doKey = "E7GJPSBYFMF5SJU7P4TC"
var doSecret = "N/K398tqDdLak67JrXYMzYEW/a1juFhC3rSxVuC5s5M"
var spacesURL = "sfo2.digitaloceanspaces.com"

// Cfg represents the app config
type Cfg struct {
	ServingURL string
	TikaURL    string
}

// App represents the application and configuration
type App struct {
	handler    http.Handler
	DB         *sqlx.DB
	HTTPClient *http.Client
	Cfg
}

// NewApp returns a new instance of App with config and DB connections loaded
func NewApp(cfg Cfg, db *sqlx.DB, client *http.Client) *App {
	a := App{
		Cfg:        cfg,
		DB:         db,
		HTTPClient: client,
	}

	a.initHandler()

	return &a
}

// ServeHTTP satisfies the http handler interface so App can perform as an http server
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}

// initHandler creates a new router initializes and all the routes
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

		stack := make([]byte, 4096)
		stack = stack[:runtime.Stack(stack, false)]

		log.Println(string(stack))

		web.RespondError(w, r, http.StatusInternalServerError, web.ErrInternalServer)
	}

	// instantiate services
	ds := document.NewService(a.DB)
	ss := store.NewService(a.DB)

	// api routes
	r.Handle(http.MethodPost, "/v1/index", a.Index(ds, ss))
	r.Handle(http.MethodGet, "/v1/search", a.Search(ds, ss))
	r.Handle(http.MethodGet, "/v1/download", a.DownloadFile())
	r.Handle(http.MethodGet, "/v1/exists", a.CheckFileExistence(ds))

	// wrap all the routes with global middleware
	a.handler = web.RequestMW(r)
}
