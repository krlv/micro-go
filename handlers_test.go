package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func Test_index(t *testing.T) {
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

func Test_redirect(t *testing.T) {
	router := httprouter.New()
	router.GET("/:hash", redirect)

	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Get(server.URL + "/example")
	if err != nil {
		t.Fatal(err)
	}

	msg, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Redirecting to example\n", string(msg), "Error redirect response")
}
