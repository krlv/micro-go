package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "Welcome!\n", string(msg), "Error index response")
}

func TestHelloHandler(t *testing.T) {
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

	assert.Equal(t, "Hello, test!\n", string(msg), "Error hello response")
}
