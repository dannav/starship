package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/dannav/starship/services/cmd/searchd/handlers"
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

	// connect to db
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
			ticker := time.NewTicker(time.Millisecond * 150)

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

	// create mysql db schema
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
	}

	// start the API
	app := handlers.NewApp(cfg, dba, client)
	server := http.Server{
		Addr:           port,
		Handler:        app,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// starting the service, listening for requests.
	serverErrors := make(chan error, 1)
	go func() {
		log.Infof("server started, listening on %s", port)
		serverErrors <- server.ListenAndServe()
	}()

	// blocking main and waiting for shutdown.
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
	// create context for shutdown call.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// asking listener to shutdown
	if err := server.Shutdown(ctx); err != nil {
		mainErr = errors.Wrapf(err, "shutdown: graceful shutdown did not complete in %v", shutdownTimeout)

		if err := server.Close(); err != nil {
			mainErr = errors.Wrapf(err, "shutdown: error killing server")
		}

		return
	}
}
