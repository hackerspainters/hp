package main

import (
	"labix.org/v2/mgo"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

var session *mgo.Session

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func homeHandler(w http.ResponseWriter, req *http.Request) {

	s := session.Clone()
	defer s.Close()

	// set up collection and query
	coll := s.DB("hp_db").C("events")
	query := coll.Find(nil).Sort("-timestamp")

	// execute query
	var events []Event
	if err := query.All(&events); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := index.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	// Establish session with mongodb
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Routing with Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/event/add/", eventAddHandler)

    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

	http.Handle("/", r)

	if err := http.ListenAndServe(":5050", nil); err != nil {
		panic(err)
	}
}
