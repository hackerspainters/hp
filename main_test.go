package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExampleIncr(t *testing.T) {
	// example unit testing in golang
	const in, out = 2, 3
	if x := ExampleIncr(in); x != out {
		t.Errorf("Example(%v) = %v, want %v", in, x, out)
	}
}

func TestHomeHandler(t *testing.T) {
	// integration test on http requests to HomeHandler

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	HomeHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}

}
