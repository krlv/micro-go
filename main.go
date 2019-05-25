package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/hello/world
func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
