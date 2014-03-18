package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"path"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	//"github.com/codegangsta/martini-contrib/binding"
	"hp/event"
	"hp/db"
	"hp/conf"
)

func HomeHandler(r render.Render) {

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	r.HTML(200, "home", data)
}

func NotFoundHandler(r render.Render) {

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	r.HTML(200, "404", data)
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

func FacebookLoginHandler(r render.Render) {

	// simple static page for user to click on fb connect button

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	r.HTML(200, "facebook_login", data)

}

func FacebookAuthHandler(w http.ResponseWriter, req *http.Request) {

	// construct fb graph's oauth end-point, then redirect user to this end-point

}

func FacebookRedirectHandler(w http.ResponseWriter, req *http.Request) {

	// returns here

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

	m := martini.Classic()

	m.Use(martini.Static("static"))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout", // Specify a layout template
		Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		//IndentJSON: true, // Output human readable JSON
	}))

	m.NotFound(NotFoundHandler)
	m.Get("/", HomeHandler)
	m.Get("/events", event.EventListHandler)
	m.Get("/events/past", event.EventPastHandler)
	m.Get("/events/next", event.EventNextHandler)
	m.Get("/organise", event.OrganiseHandler)
	m.Get("/event/add", event.EventAddHandler)

	// Facebook related features
	// one-off link that allows event owner to grab group-specific events set with group-only perms
	m.Get("/facebook/login", FacebookLoginHandler)
	m.Get("/channel.html", FacebookChannelHandler)
	m.Get("/events/grab", event.EventGrabHandler)
	m.Post("/events/import", event.EventImportHandler)

	m.Run()

}
