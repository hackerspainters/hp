package event

import (
	"fmt"
	"reflect"
	"net/http"
	"path"
	"encoding/json"
	"github.com/hackerspainters/facebook"
	"html/template"
	"labix.org/v2/mgo/bson"

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
		"templates/_base.html",
		"templates/event_grab.html",
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	eventgrab.Execute(w, data)

}

func EventImportHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	MyToken := facebook.AccessToken{}
	err := decoder.Decode(&MyToken)
	if err != nil {
		panic(err)
	}

	events := facebook.GetGroupEvents(&MyToken, conf.Config.FacebookGroupId)
	event_ids := facebook.GetGroupEventIds(events)
	event := NewEvent()

	for i := 0; i < len(event_ids); i++ {
		e := facebook.GetEvent(&MyToken, event_ids[i])
		fmt.Println(e.Id)
		fmt.Println(reflect.TypeOf(e.Id))

		var ev *Event
		err := db.Find(ev, bson.M{"eid": e.Id}).One(&ev)
		if err != nil {
			// Not found, so insert our event object
			event.Eid = e.Id
			event.Data = e
			db.Upsert(event)
		} else {
			// Already exists, so simply update as the retrieved ev object
			ev.Data = e
			db.Upsert(ev)
		}
	}

}
