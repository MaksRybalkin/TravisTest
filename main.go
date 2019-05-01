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
)

func main() {
	var mainErr error
	defer func() {
		if mainErr != nil {
			log.Fatalf("Shutting down with error: %s", mainErr)
			os.Exit(1)
		}
	}()

	server := http.Server{
		Addr:           ":8080",
		Handler:        handlers(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start listening for requests made to the daemon and create a channel
	// to collect non-HTTP related server errors on.
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("server started, listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown of the daemon.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Waiting for an osSignal or a non-HTTP related server error.
	select {
	case e := <-serverErrors:
		mainErr = fmt.Errorf("server failed to start: %+v", e)
		return

	case <-osSignals:
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown : Graceful shutdown did not complete in %v : %v", 5*time.Second, err)

		if err := server.Close(); err != nil {
			log.Printf("shutdown : Error killing server : %v", err)
		}
	}
}
