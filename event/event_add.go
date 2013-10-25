package event

import (
	"net/http"
	//"fmt"
	"html/template"
	"labix.org/v2/mgo"

	"hp/db"
)

var session *mgo.Session

var eventadd = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/event_add.html",
))

func EventAddHandler(w http.ResponseWriter, req *http.Request) {

	// if request method is a GET, we will simply render the page
	if req.Method != "POST" {
		eventadd.Execute(w, nil)
		return
	}

	// else if it is a POST, let's add our event
	event := NewEvent()
	event.Name = req.FormValue("name")
	event.Description = req.FormValue("description")

	//if event.Name == "" {
		//fmt.Println("No event name submitted")
	//}

	//if event.Description == "" {
		//fmt.Println("No event description submitted")
	//}

	db.Upsert(event)

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)

}
