package event

import (
	"time"

	"labix.org/v2/mgo/bson"

	"hp/db"
)

type Event struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	Slug        string
	Timestamp	time.Time
	Name		string
	Description	string
	Speaker		int
	Published   bool
}

func NewEvent() *Event {
	return &Event{
		Timestamp:time.Now(),
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
        return bson.M{"slug": e.Slug}
}

// register "models"

func init() {
	db.Register(&Event{})
}
