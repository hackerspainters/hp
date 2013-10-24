package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Event struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	Timestamp	time.Time
	Name		string
	Description	string
	Speaker		int
}

func NewEvent() *Event {
	return &Event{
		Timestamp:time.Now(),
	}
}
