package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8080/hello/world
func main() {
	log.Info("Initializing service...")

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/:hash", redirect)

	srv := &http.Server{Addr: ":8080", Handler: router}

	idle := make(chan os.Signal, 1)
	signal.Notify(idle, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	log.Info("Service is up and running")

	<-idle
	log.Info("Service stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		// TODO: extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Info("Service stopped")
}
