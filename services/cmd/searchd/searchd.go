package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/dannav/starship/services/cmd/searchd/handlers"
	"github.com/dannav/starship/services/cmd/searchd/healthcheck"
	"github.com/dannav/starship/services/internal/platform/db"
	_ "github.com/lib/pq" // postgres sql driver
	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"
)

const (
	// readTimeout is timeout for reading the request
	readTimeout = 30 * time.Second

	// writeTimeout is timeout for reading the response
	writeTimeout = 30 * time.Second

	// shutdownTimeout is the timeout for shutdown
	shutdownTimeout = 30 * time.Second

	// defaultAppPort is the port that the application listens on if no PORT env is set
	defaultAppPort = 8080

	// appPortEnv is the name of the environment variable that contains the port to run searchd on
	appPortEnv = "PORT"

	// postgresDSNEnv is the name of the environment variable that contains the postgres connection string dsn.
	postgresDSNEnv = "POSTGRES_DSN"

	// modelURLEnv is the name of the environment variable that holds the URL
	// to machine learning model endpoint. it should point directly to the machine learning model
	// and follow tensorflow serving conventions. i.e. http://tfserving/v1/models/{model_name}
	modelURLEnv = "MODEL_URL"

	// modelVectorDimensionsEnv is the name of the environment variable that defines how many vector dimensions the
	// machine learning model returns for a sentence embedding
	modelVectorDimensionsEnv = "MODEL_VECTOR_DIMENSIONS"

	// tikadURLEnv is the name of the environment variable that holds the URL
	// to the tika service which converts documents to plain text
	tikadURLEnv = "TIKAD_URL"

	// objectStorageURLEnv is the URL to the object storage provider
	objectStorageURLEnv = "OBJECT_STORAGE_URL"

	// objectStorageBucketNameEnv is the bucket to connect to with your object storage provider
	objectStorageBucketNameEnv = "OBJECT_STORAGE_BUCKET"

	// objectStorageKey is the key to use when connecting to object storage
	objectStorageKeyEnv = "OBJECT_STORAGE_KEY"

	// objectStorageSecret is the secret ot use when connecting to object storage
	objectStorageSecretEnv = "OBJECT_STORAGE_SECRET"

	// storagePathEnv is the location to store the index and if no object storage is used documents
	// ensure that the user running this app has permissions to read / write here
	storagePathEnv = "STORAGE_PATH"

	// storageDir is the root folder to store cfg, indexes, and files given the storagePath
	storageDir = ".starship"
)

func main() {
	var runPort int

	// if there was an error, exit main with non-zero status code.
	var mainErr error
	defer func() {
		if mainErr != nil {
			log.WithFields(log.Fields{
				"error": mainErr,
			}).Error("error in main")

			os.Exit(1)
		}
	}()

	// parse app config from environment variables
	setPort := os.Getenv(appPortEnv)
	if setPort == "" {
		runPort = defaultAppPort
	} else {
		p, err := strconv.Atoi(setPort)
		if err != nil {
			mainErr = errors.Errorf("cannot convert PORT environment variable to int %s", err.Error())
			return
		}

		runPort = p
	}
	port := fmt.Sprintf(":%d", runPort)

	postgresDSN := os.Getenv(postgresDSNEnv)
	if postgresDSN == "" {
		mainErr = errors.Errorf("missing required environment variable %s", postgresDSNEnv)
		return
	}

	modelURL := os.Getenv(modelURLEnv)
	if modelURL == "" {
		mainErr = errors.Errorf("missing required environment variable %s", modelURLEnv)
		return
	}

	modelVectorDims := os.Getenv(modelVectorDimensionsEnv)
	if modelVectorDims == "" {
		mainErr = errors.Errorf("missing required environment variable %s", modelVectorDimensionsEnv)
		return
	}

	vDims, err := strconv.Atoi(modelVectorDims)
	if err != nil {
		mainErr = errors.Errorf("environment variable %v cannot be converted to an int: %s", modelVectorDimensionsEnv, err.Error())
		return
	}

	tikadURL := os.Getenv(tikadURLEnv)
	if tikadURL == "" {
		mainErr = errors.Errorf("missing required environment variable %s", tikadURLEnv)
		return
	}

	storagePath := os.Getenv(storagePathEnv)
	if storagePath == "" {
		u, err := user.Current()
		if err != nil {
			mainErr = errors.Wrap(err, "could not get user running app")
			return
		}

		// set storagePath to the default storage path to store cfg, indexes, and files
		storagePath = filepath.Join(u.HomeDir, storageDir)
	} else {
		storagePath = filepath.Clean(storagePath)
	}

	// create storagePath deepest dir (indexes)
	err = os.MkdirAll(filepath.Join(storagePath, "indexes"), os.ModePerm)
	if err != nil {
		mainErr = errors.Wrap(err, "could not create storage directory")
		return
	}

	// wait for tikad to be ready
	var tikaReady bool
	for {
		log.Info("waiting for tikad to be ready")
		ticker := time.NewTicker(time.Second)

		select {
		case <-ticker.C:
			ready, err := healthcheck.TikaServiceReady(tikadURL)
			if err != nil && strings.Index(err.Error(), "connection refused") == -1 { // skip exiting if the service isn't started yet
				mainErr = errors.Wrap(err, "tikad connection error")
				return
			}

			if ready == true {
				tikaReady = true
			}
		}

		if tikaReady {
			break
		}
	}

	// wait for ML model API to be ready
	var modelAPIReady bool
	for {
		log.Info("waiting for ml model api to be ready")
		ticker := time.NewTicker(time.Second)

		select {
		case <-ticker.C:
			ready, err := healthcheck.ModelReady(modelURL)
			if err != nil && strings.Index(err.Error(), "connection refused") == -1 { // skip exiting if the service isn't started yet
				mainErr = errors.Wrap(err, "ml model api connection error")
				return
			}

			if ready == true {
				modelAPIReady = true
			}
		}

		if modelAPIReady {
			break
		}
	}

	// create new postgres connection
	dba, err := sqlx.Open("postgres", postgresDSN)
	if err != nil {
		mainErr = errors.Wrap(err, "open postgres database")
		return
	}

	// wait till we get a connection to the database
	var connected bool
	if err := dba.Ping(); err != nil {
		for {
			log.Info("waiting for database connection")
			ticker := time.NewTicker(time.Second)

			select {
			case <-ticker.C:
				if err := dba.Ping(); err == nil {
					connected = true
				}
			}

			if connected {
				break
			}
		}
	}

	// create db schema
	if _, err := dba.Exec(db.Schema); err != nil {
		mainErr = errors.Wrap(err, "creating database schema")
		return
	}

	// create db indexes
	if _, err := dba.Exec(db.Indexes); err != nil {
		mainErr = errors.Wrap(err, "creating database indexes")
		return
	}

	// run db seed
	for _, s := range db.Seed {
		if _, err := dba.Exec(s); err != nil {
			mainErr = errors.Wrap(err, "seeding database")
			return
		}
	}

	// force 30 second timeouts on all http requests
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	cfg := handlers.Cfg{
		ModelURL:        modelURL,
		ModelVectorDims: vDims,
		TikaURL:         tikadURL,
		StoragePath:     storagePath,
		ObjectStorageConfig: handlers.ObjectStorageCfg{ // get object storage cfg if set in env
			URL:        os.Getenv(objectStorageURLEnv),
			BucketName: os.Getenv(objectStorageBucketNameEnv),
			Key:        os.Getenv(objectStorageKeyEnv),
			Secret:     os.Getenv(objectStorageSecretEnv),
		},
	}

	// create the API
	app := handlers.NewApp(cfg, dba, client)
	server := http.Server{
		Addr:           port,
		Handler:        app,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// starting the API server
	serverErrors := make(chan error, 1)
	go func() {
		log.Infof("server started, listening on %s", port)
		serverErrors <- server.ListenAndServe()
	}()

	// blocking main and waiting for shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// wait for osSignal or error starting server
	select {
	case e := <-serverErrors:
		mainErr = e
		return

	case <-osSignals:
	}

	// shutdown server
	// create context for shutdown call
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// asking listener to shutdown
	if err := server.Shutdown(ctx); err != nil {
		mainErr = errors.Wrapf(err, "shutdown: graceful shutdown did not complete in %v", shutdownTimeout)

		if err := server.Close(); err != nil {
			mainErr = errors.Wrapf(err, "shutdown: error killing server")
		}
	}
}
