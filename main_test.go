package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/codegangsta/martini"
)

func TestHomeHandler(t *testing.T) {
	// integration test on http requests to HomeHandler

	m := martini.Classic()
	m.Get("/", HomeHandler)

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	m.ServeHTTP(response, request)

}
