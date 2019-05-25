package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestIndexHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", index)

	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	msg, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expected := "Welcome!\n"
	if expected != string(msg) {
		t.Errorf("Erong index response '%s', expected '%s'", msg, expected)
	}
}

func TestHellpHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/hello/:name", hello)

	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Get(server.URL + "/hello/test")
	if err != nil {
		t.Fatal(err)
	}

	msg, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello, test!\n"
	if expected != string(msg) {
		t.Errorf("Erong index response '%s', expected '%s'", msg, expected)
	}
}
