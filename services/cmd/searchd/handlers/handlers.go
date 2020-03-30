package handlers

import (
	"net/http"
	"runtime"

	"github.com/dannav/starship/services/internal/document"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/dannav/starship/services/internal/store"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ObjectStorageCfg represents configuration passed for connecting to object storage
type ObjectStorageCfg struct {
	URL        string
	BucketName string
	Key        string
	Secret     string
}

// Enabled checks if ObjectStorageCfg is configured properly
func (o *ObjectStorageCfg) Enabled() bool {
	if o.URL != "" && o.BucketName != "" && o.Key != "" && o.Secret != "" {
		return true
	}

	return false
}

// Cfg represents the app config
// ModelVectorDims represents how many dimensions a spotify/annoy index should be be (size of sentence embedding returned from ml model)
type Cfg struct {
	ModelURL            string
	ModelName           string
	ModelVectorDims     int
	TikaURL             string
	StoragePath         string
	IndexKey            string
	AccessKey           string
	ObjectStorageConfig ObjectStorageCfg
}

// App represents the application and configuration
type App struct {
	handler              http.Handler
	DB                   *sqlx.DB
	HTTPClient           *http.Client
	ObjectStorageEnabled bool
	Cfg
}

// NewApp returns a new instance of App with config and DB connections loaded
func NewApp(cfg Cfg, db *sqlx.DB, client *http.Client) *App {
	a := App{
		Cfg:                  cfg,
		DB:                   db,
		HTTPClient:           client,
		ObjectStorageEnabled: cfg.ObjectStorageConfig.Enabled(),
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

	// gracefully handle and recover from web panics, print stack to log for debugging purposes
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		stack := make([]byte, 4096)
		stack = stack[:runtime.Stack(stack, false)]

		log.WithFields(log.Fields{
			"error": i,
			"stack": string(stack),
		}).Error("panic")

		web.RespondError(w, r, http.StatusInternalServerError, web.ErrInternalServer)
	}

	// instantiate services
	ds := document.NewService(a.DB)
	ss := store.NewService(a.DB)

	// get keys for authentication purposes
	indexKey := a.IndexKey
	accessKey := a.AccessKey

	// ready route responds when the server is booted up and ready to accept incoming connections
	r.Handle(http.MethodGet, "/ready", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		web.Respond(w, r, http.StatusOK, nil)
	})

	// define api routes
	r.Handle(http.MethodPost, "/v1/index", IndexAuthorized(a.Index(ds, ss), indexKey))
	r.Handle(http.MethodGet, "/v1/search", AccessAuthorized(a.Search(ds, ss), accessKey))
	r.Handle(http.MethodGet, "/v1/download", AccessAuthorized(a.DownloadFile(), accessKey))
	r.Handle(http.MethodGet, "/v1/exists", AccessAuthorized(a.CheckFileExistence(ds), accessKey))

	// wrap all the routes with global middleware
	a.handler = web.RequestMW(r)
}
