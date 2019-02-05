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

	"github.com/dannav/starship/services/cmd/tikad/handlers"
	"github.com/google/go-tika/tika"
	_ "github.com/lib/pq" // postgres sql driver
	log "github.com/sirupsen/logrus"

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

	// tikaVersion is the version of apache tika to download
	tikaVersion = "1.16"

	// tikaPort is the port that apache tika runs on
	tikaPort = "9998"
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

	// check if tika jar exists or download it
	tikaJar := fmt.Sprintf("tika-server-%v.jar", tikaVersion)
	if _, err := os.Stat(tikaJar); os.IsNotExist(err) {
		log.Infof("apache tika not found, downloading version %v", tikaVersion)
		err := tika.DownloadServer(context.Background(), tika.Version116, tikaJar)
		if err != nil {
			mainErr = errors.Wrap(err, "downloading apache tika")
			return
		}
		log.Info("apache tika downloaded successfully")
	}

	// force 30 second timeouts on all http requests
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	opts := tika.WithPort(tikaPort)
	s, err := tika.NewServer(tikaJar, opts)
	if err != nil {
		mainErr = errors.Wrap(err, "instantiating new apache tika server")
		return
	}

	// start tika server
	shutdownTika, err := s.Start(context.Background())
	if err != nil {
		mainErr = errors.Wrap(err, "starting apache tika")
		return
	}

	tikaClient := tika.NewClient(client, s.URL())
	app := handlers.NewApp(tikaClient)

	server := http.Server{
		Addr:           port,
		Handler:        app,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// start the API server
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
	// create context for Shutdown call.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// shutdown apache tika
	shutdownTika()

	// asking listener to shutdown
	if err := server.Shutdown(ctx); err != nil {
		mainErr = errors.Wrapf(err, "shutdown: graceful shutdown did not complete in %v", shutdownTimeout)

		if err := server.Close(); err != nil {
			mainErr = errors.Wrapf(err, "shutdown: error killing server")
		}

		return
	}
}
