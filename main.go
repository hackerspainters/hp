package main

import (
	"fmt"
	"reflect"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"hp/event"
	"hp/db"
	"hp/conf"
)

func HomeHandler(r render.Render) {
	r.HTML(200, "home", "")
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
