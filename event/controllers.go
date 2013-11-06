package event

import (
	"fmt"
	"time"
	"net/http"
	"path"
	"encoding/json"
	"html/template"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/hackerspainters/facebook"

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

func EventListHandler(w http.ResponseWriter, r *http.Request) {

	search := bson.M{"data.start_time": bson.M{"$gte": time.Now()}}
	sort := "data.start_time"
	var results []Event
	err := db.Find(&Event{}, search).Sort(sort).All(&results)
	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}
	if err == mgo.ErrNotFound {
		fmt.Println("No such object in db. Redirect")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	// TODO: 
	// This is the absolute path parsing of template files so tests will pass
	// Code can be better organized
	var eventlist = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_list.html"),
	))

	type templateData struct {
		Context *conf.Context
		Events []Event
	}

	data := templateData{conf.DefaultContext(conf.Config), results}

	eventlist.Execute(w, data)

}

func EventPastHandler(w http.ResponseWriter, r *http.Request) {

	search := bson.M{"data.start_time": bson.M{"$lte": time.Now()}}
	sort := "-data.start_time"
	var results []Event
	err := db.Find(&Event{}, search).Sort(sort).All(&results)
	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}
	if err == mgo.ErrNotFound {
		fmt.Println("No such object in db. Redirect")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	// TODO: 
	// This is the absolute path parsing of template files so tests will pass
	// Code can be better organized
	var eventpast = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_past.html"),
	))

	type templateData struct {
		Context *conf.Context
		Events []Event
	}

	data := templateData{conf.DefaultContext(conf.Config), results}

	eventpast.Execute(w, data)

}

// This handler renders the event on the upcoming Friday
func EventNextHandler(w http.ResponseWriter, r *http.Request) {

	search := bson.M{"data.start_time": bson.M{"$gte": time.Now()}}
	sort := "data.start_time"
	var result *Event
	err := db.Find(&Event{}, search).Sort(sort).One(&result)
	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}
	if err == mgo.ErrNotFound {
		fmt.Println("No such object in db. Redirect")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	// TODO: 
	// This is the absolute path parsing of template files so tests will pass
	// Code can be better organized
	var eventnext = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/event_next.html"),
	))

	type templateData struct {
		Context *conf.Context
		E *Event
	}

	data := templateData{conf.DefaultContext(conf.Config), result}

	eventnext.Execute(w, data)

}

func OrganiseHandler(w http.ResponseWriter, r *http.Request) {

	var organise = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/organise.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	organise.Execute(w, data)

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
	updateEventDetails(event_ids, &MyToken)

	// one-off implementation to also grab event details in the `General Hackers` group
    generalhackers := facebook.GetGroupEvents(&MyToken, "314660778669731")
	generalhackers_ids := facebook.GetGroupEventIds(generalhackers)
	updateEventDetails(generalhackers_ids, &MyToken)

}

// helper function which updates the events given a slice of event ids (string) and the token
func updateEventDetails(event_ids []string, token *facebook.AccessToken) {

	for i := 0; i < len(event_ids); i++ {
		e := facebook.GetEvent(token, event_ids[i])

		var result *Event
		err := db.Find(&Event{}, bson.M{"eid": e.Id}).One(&result)
		if err != nil {
			// Not found, so insert our event object
			event := wrangleData(e)
			db.Upsert(event)
		} else {
			// Already exists, so simply update as the retrieved result object
			result := wrangleData(e)
			db.Upsert(result)
		}
	}

}

// cast to correct types and handle our own custom fields as needed
func wrangleData(e facebook.Event) *Event {
	const layout = "2006-01-02T15:04:05-0700"
	event := NewEvent()
	event.Eid = e.Id
	event.Data.Name = e.Name
	event.Data.Description = e.Description
	event.Data.StartTime, _ = time.Parse(layout, e.StartTime)
	event.Data.EndTime, _ = time.Parse(layout, e.EndTime)
	event.Data.UpdatedTime, _ = time.Parse(layout, e.UpdatedTime)
	event.Data.TimeZone = e.TimeZone
	event.Data.IsDateOnly = e.IsDateOnly
	event.Data.Location = e.Location
	event.Data.Venue.Latitude = e.Venue.Latitude
	event.Data.Venue.Longitude = e.Venue.Longitude
	event.Data.Venue.City = e.Venue.City
	event.Data.Venue.Country = e.Venue.Country
	event.Data.Venue.Id = e.Venue.Id
	event.Data.Venue.Street = e.Venue.Street
	event.Data.Venue.Zip = e.Venue.Zip
	event.Data.UpdatedTime, _ = time.Parse(layout, e.UpdatedTime)
	return event
}
