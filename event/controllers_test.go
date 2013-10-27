package event

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestEventNextHandler(t *testing.T) {

	// integration test on http requests to EventNextHandler

	request, _ := http.NewRequest("GET", "/events/next/", nil)
	response := httptest.NewRecorder()

	EventNextHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}

}
