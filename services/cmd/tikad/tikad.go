package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dannav/starship/services/cmd/tikad/handlers"
	"github.com/google/go-tika/tika"
	_ "github.com/lib/pq" // postgres sql driver
	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

const (
	// readTimeout is timeout for reading the request.
	readTimeout = 5 * time.Second

	// writeTimeout is timeout for reading the response.
	writeTimeout = 10 * time.Second

	// shutdownTimeout is the timeout for shutdown.
	shutdownTimeout = 5 * time.Second

	// appPort is the port that the application listens on
	appPort = 8080

	tikaVersion = "1.16"
)

func main() {
	port := fmt.Sprintf(":%d", appPort)

	// If there was an error, exit main with non-zero status code.
	var mainErr error
	defer func() {
		if mainErr != nil {
			log.WithFields(log.Fields{
				"error": mainErr,
			}).Error("error in main")

			os.Exit(1)
		}
	}()

	// Start Apache Tika (used for detecting and converting filetypes)
	// ===============================================================

	tikaJar := fmt.Sprintf("tika-server-%v.jar", tikaVersion)

	// check if tika jar exists or download it
	if _, err := os.Stat(tikaJar); os.IsNotExist(err) {
		log.Infof("apache tika not found, downloading version %v", tikaVersion)
		err := tika.DownloadServer(context.Background(), tika.Version116, tikaJar)
		if err != nil {
			mainErr = errors.Wrap(err, "downloading apache tika")
			return
		}
		log.Info("apache tika downloaded successfully")
	}

	// force 10 second timeouts on all http requests
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	opts := tika.WithPort("9998")
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

	// Starting the service, listening for requests.
	serverErrors := make(chan error, 1)
	go func() {
		log.Infof("server started, listening on %s", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Wait for osSignal or error starting server
	select {
	case e := <-serverErrors:
		mainErr = e
		return

	case <-osSignals:
	}

	// Shutdown Server
	// ===============

	// Create context for Shutdown call.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// shutdown apache tika
	shutdownTika()

	// Asking listener to shutdown
	if err := server.Shutdown(ctx); err != nil {
		mainErr = errors.Wrapf(err, "shutdown: graceful shutdown did not complete in %v", shutdownTimeout)

		if err := server.Close(); err != nil {
			mainErr = errors.Wrapf(err, "shutdown: error killing server")
		}

		return
	}
}
