package event

import (
	"time"

	"labix.org/v2/mgo/bson"

	"hp/db"
)

type Attendee struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Eid          string
	Email        string
	Fbuid        string
	FirstName    string
	LastName     string
    Timestamp    time.Time
}

func NewAttendee() *Attendee {
	return &Attendee{
		Timestamp: time.Now(),
	}
}

// implementations for the Attendee struct.

func (a *Attendee) Collection() string { return "attendees" }

func (a *Attendee) Indexes() [][]string {
	return [][]string{
		[]string{"timestamp"},
	}
}

func (a *Attendee) PreSave() {}

func (a *Attendee) Unique() bson.M {
	if len(a.ID) > 0 {
		return bson.M{"_id": a.ID}
	}
	return bson.M{"eid": a.Eid}
}

type Event struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Eid          string
	Slug         string
	Name         string
	Description  string
    Timestamp    time.Time
	Speaker      int
	Published    bool

	Data struct {
		Id           string
		Owner struct {
			Name    string
			Id      string
		}
		Name            string
		Description     string
		StartTime		time.Time    `bson:"start_time" json:"start_time"`
		EndTime         time.Time    `bson:"end_time" json:"end_time"`
		TimeZone		string
		IsDateOnly		bool
		Location		string
		Venue struct {
			Latitude	float64
			Longitude	float64
			City		string
			Country	    string
			Id			string
			Street		string
			Zip		    string
		}
		UpdatedTime	    time.Time    `bson:"updated_time" json:"updated_time"`
	}
}

func NewEvent() *Event {
	return &Event{
		Timestamp: time.Now(),
	}
}


// implementations for the Event struct.

func (e *Event) Collection() string { return "events" }

func (e *Event) Indexes() [][]string {
	return [][]string{
		[]string{"timestamp"},
	}
}

func (e *Event) PreSave() {}

func (e *Event) Unique() bson.M {
	if len(e.ID) > 0 {
		return bson.M{"_id": e.ID}
	}
	return bson.M{"eid": e.Eid}
}

// register "models"

func init() {
	db.Register(&Event{})
	db.Register(&Attendee{})
}

