package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"

	"hp/conf"
	"hp/db"
	"hp/event"
	"hp/user"
)

func ExampleIncr(x float64) float64 {
	// example function that will be used for unit testing in golang
	return x + 1
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	// homepage controller

	var index = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/index.html",
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	if err := index.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FacebookChannelHandler(w http.ResponseWriter, req *http.Request) {

	var fbchannel = template.Must(template.ParseFiles(
		"templates/channel.html",
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	fbchannel.Execute(w, data)
}

func FacebookLoginHandler(w http.ResponseWriter, req *http.Request) {

	// simple static page for user to click on fb connect button

	var fblogin = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/facebook_login.html",
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	fblogin.Execute(w, data)

}

func FacebookAuthHandler(w http.ResponseWriter, req *http.Request) {

	// construct fb graph's oauth end-point, then redirect user to this end-point

}

func FacebookRedirectHandler(w http.ResponseWriter, req *http.Request) {

	// returns here

}

func handleFuncPrefix(r *mux.Router, s string, h func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(conf.Config.HttpPrefix+s, h)
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
	handleFuncPrefix(r, "/", HomeHandler)
	handleFuncPrefix(r, "/channel.html", FacebookChannelHandler)
	handleFuncPrefix(r, "/facebook/login/", FacebookLoginHandler)

	handleFuncPrefix(r, "/users/", user.UsersHandler)
	handleFuncPrefix(r, "/events/", event.EventListHandler)
	handleFuncPrefix(r, "/events/next/", event.EventNextHandler)
	handleFuncPrefix(r, "/events/past/", event.EventPastHandler)
	handleFuncPrefix(r, "/event/add/", event.EventAddHandler)

	// one-off link that allows event owner to grab group-specific events set with group-only perms
	handleFuncPrefix(r, "/events/grab/", event.EventGrabHandler)
	handleFuncPrefix(r, "/events/import/", event.EventImportHandler)

	handleFuncPrefix(r, "/static/{_:.*}", func(w http.ResponseWriter, r *http.Request) {
		// Ignore prefix + leading /
		http.ServeFile(w, r, r.URL.Path[len(conf.Config.HttpPrefix)+1:])
	})

	http.Handle("/", r)

	if err := http.ListenAndServe(":5050", nil); err != nil {
		panic(err)
	}
}
