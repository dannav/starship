package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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

	// defaultAppPort is the port that the application listens on if not PORT env is set
	defaultAppPort = 8080

	// appPortEnv is the name of the environment variable that contains the port to run searchd on
	appPortEnv = "PORT"

	// postgresDSNEnv is the name of the environment variable that contains the postgres connection string dsn.
	postgresDSNEnv = "POSTGRES_DSN"

	// servingURLEnv is the name of the environment variable that holds the URL
	// to serving endpoint of the universal sentence encoder
	servingURLEnv = "SERVING_URL"

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

	servingURL := os.Getenv(servingURLEnv)
	if servingURL == "" {
		mainErr = errors.Errorf("missing required environment variable %s", servingURLEnv)
		return
	}

	tikadURL := os.Getenv(tikadURLEnv)
	if tikadURL == "" {
		mainErr = errors.Errorf("missing required environment variable %s", tikadURLEnv)
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

	// wait for tfserving to be ready
	var servingReady bool
	for {
		log.Info("waiting for serving to be ready")
		ticker := time.NewTicker(time.Second)

		select {
		case <-ticker.C:
			ready, err := healthcheck.ServingReady(servingURL)
			if err != nil && strings.Index(err.Error(), "connection refused") == -1 { // skip exiting if the service isn't started yet
				mainErr = errors.Wrap(err, "serving connection error")
				return
			}

			if ready == true {
				servingReady = true
			}
		}

		if servingReady {
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
		ServingURL: servingURL,
		TikaURL:    tikadURL,
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
