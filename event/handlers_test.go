package event

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/martini"

	// TODO: begin using suite and assert from testify
	//"github.com/stretchr/testify/suite"
	//"github.com/stretchr/testify/assert"

	"hp/db"
)



func TestEventNextHandler(t *testing.T) {

	// set up test database

	fmt.Println("testing EventNext")
	db.Connect("127.0.0.1", "test_db")
	db.RegisterAllIndexes()

	// integration test on http requests to EventNextHandler
	m := martini.Classic()
	m.Get("/events/next/", EventNextHandler)

	request, _ := http.NewRequest("GET", "/events/next/", nil)
	response := httptest.NewRecorder()

	m.ServeHTTP(response, request)

}
