package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/hello/world
func main() {
    router := httprouter.New()
    router.GET("/", index)
    router.GET("/hello/:name", hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
