package main

import (
	"net/http"
	"os"

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
	log.SetLevel(log.WarnLevel)
}

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8080/hello/world
func main() {
	log.Info("Initializing service...")

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	log.Info("Service is up and running")
	http.ListenAndServe(":8080", router)
}
