package main

import (
	"reflect"
	"html/template"
	"net/http"
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"github.com/gorilla/mux"

	"hp/conf"
	"hp/db"
	"hp/event"
	"hp/user"
)

var session *mgo.Session

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

type M bson.M

func HomeHandler(w http.ResponseWriter, req *http.Request) {

	//s := session.Clone()
	//defer s.Close()

	// set up collection and query
	//coll := s.DB("hp_db").C("events")
	//query := coll.Find(nil).Sort("-timestamp")

	// execute query
	//var events []event.Event
	//err = db.Find(&event.Event{}, 
	//if err := db.Find(&event.Event{}, M{}).Sort("-timestamp"); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
	//}

	if err := index.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fmt.Printf("Using config %s\n", conf.Path)
	fmt.Printf("Using models:\n")
	for _, m := range db.Models {
		t := reflect.TypeOf(m)
		fmt.Printf("    %s\n", fmt.Sprintf("%s", t)[1:])
	}

	// Establish session with mongodb
	db.Connect(conf.Config.DbHostString(), conf.Config.DbName)
	db.RegisterAllIndexes()

	// Routing with Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users/", user.UsersHandler)
	r.HandleFunc("/events/", event.EventListHandler)
	r.HandleFunc("/events/next/", event.EventNextHandler)
	r.HandleFunc("/events/past/", event.EventPastHandler)
	r.HandleFunc("/event/add/", event.EventAddHandler)

    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

	http.Handle("/", r)

	if err := http.ListenAndServe(":5050", nil); err != nil {
		panic(err)
	}
}
