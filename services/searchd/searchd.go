package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"starship/services/searchd/handlers"
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
)

func main() {
	app := handlers.NewApp()
	port := fmt.Sprintf(":%d", appPort)

	// Start API
	// ===============

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
		log.Printf("server started, listening on %s", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Wait for osSignal or error starting server
	select {
	case e := <-serverErrors:
		log.Fatalf("server failed to start: %+v", e)

	case <-osSignals:
	}

	// Shutdown Server
	// ===============

	// Create context for Shutdown call.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Asking listener to shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown: graceful shutdown did not complete in %v : %v", shutdownTimeout, err)

		if err := server.Close(); err != nil {
			log.Printf("shutdown: error killing server : %v", err)
		}
	}
}
