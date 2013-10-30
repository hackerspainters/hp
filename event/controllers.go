package event

import (
	"net/http"
	"path"
	//"fmt"
	"html/template"

	//"github.com/hackerspainters/facebook"
	//"github.com/gorilla/context"

	"hp/conf"
	"hp/db"
)

func EventAddHandler(w http.ResponseWriter, req *http.Request) {

	// Public facing page that allows users (required field email and event title) to submit a topic
	// TODO: implement check to ensure user submitted event provides an email

	var eventadd = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_add.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	// if request method is a GET, we will simply render the page
	if req.Method != "POST" {
		eventadd.Execute(w, data)
		return
	}

	// else if it is a POST, let's add our event
	event := NewEvent()
	event.Name = req.FormValue("name")
	event.Description = req.FormValue("description")

	// TODO: validation
	//if event.Name == "" {
	//fmt.Println("No event name submitted")
	//}

	//if event.Description == "" {
	//fmt.Println("No event description submitted")
	//}

	db.Upsert(event)

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)

}

func EventListHandler(w http.ResponseWriter, req *http.Request) {

	// TODO: implement db.Find to retrieve data dynamically

	var eventlist = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_list.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	eventlist.Execute(w, data)

}

func EventPastHandler(w http.ResponseWriter, req *http.Request) {

	// TODO: implement db.Find to retrieve data dynamically

	var eventpast = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_past.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	eventpast.Execute(w, data)

}

func EventNextHandler(w http.ResponseWriter, req *http.Request) {

	// TODO: implement db.Find to retrieve data dynamically

	// TODO: simplify this with os.Glob
	var eventnext = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_next.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	eventnext.Execute(w, data)

}

func EventGrabHandler(w http.ResponseWriter, req *http.Request) {

	var eventgrab = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_grab.html"),
		//"templates/_base.html",
		//"templates/event_grab.html",
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	eventgrab.Execute(w, data)

}
