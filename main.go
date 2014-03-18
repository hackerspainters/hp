package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"path"

	"github.com/gorilla/mux"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"hp/event"
	"hp/db"
	"hp/conf"
)

func HomeHandler(r render.Render) {
	r.HTML(200, "home", "")
}

func NotFoundHandler(w http.ResponseWriter, req *http.Request) {

	var notfound = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/404.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	if err := notfound.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func FacebookChannelHandler(w http.ResponseWriter, req *http.Request) {

	var fbchannel = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/channel.html"),
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
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/facebook_login.html"),
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
	r.StrictSlash(true)
	handleFuncPrefix(r, "/404/", NotFoundHandler)
	handleFuncPrefix(r, "/channel.html", FacebookChannelHandler)
	handleFuncPrefix(r, "/facebook/login/", FacebookLoginHandler)

	m := martini.Classic()

	m.Use(martini.Static("static"))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout", // Specify a layout template
		Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		//IndentJSON: true, // Output human readable JSON
	}))

	m.Get("/", HomeHandler)
	m.Get("/events/", event.EventListHandler)
	m.Get("/events/past/", event.EventPastHandler)
	m.Get("/events/next/", event.EventNextHandler)
	m.Get("/organise/", event.OrganiseHandler)
	m.Get("/event/add/", event.EventAddHandler)

	// one-off link that allows event owner to grab group-specific events set with group-only perms
	m.Get("/events/grab/", event.EventGrabHandler)
	m.Get("/events/import/", event.EventImportHandler)

	m.Run()

}
