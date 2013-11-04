package event

import (
	"time"
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

// This handler renders the event on the upcoming Friday
func EventNextHandler(w http.ResponseWriter, req *http.Request) {

	var event *Event
	search := bson.M{"data.start_time": bson.M{"$gte": time.Now()}}
	sort := "data.start_time"
	err := db.Find(event, search).Sort(sort).One(&event)
	if err != nil {
		panic(err)
	}

	var eventnext = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/event_next.html",
	))

	type templateData struct {
		Context *conf.Context
		E *Event
	}

	data := templateData{conf.DefaultContext(conf.Config), event}

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
	//event := NewEvent()

	for i := 0; i < len(event_ids); i++ {
		e := facebook.GetEvent(&MyToken, event_ids[i])

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
