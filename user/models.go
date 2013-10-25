package user

import (
	"time"

	"labix.org/v2/mgo/bson"

	"hp/db"
)


type User struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	Username    string
	Firstname	string
	Lastname	string
	Speaker		int
	Active      bool  // if false, user cannot log in
	Admin       bool  // for restriction to admin-specific management pages subsequently
	Timestamp	time.Time
}

// implementations for the User struct. 

func (u *User) Collection() string { return "users" }

func (u *User) Indexes() [][]string {
	return [][]string{
		[]string{"username"},
	}
}

func (u *User) PreSave() {}


func (u *User) Unique() bson.M {
        if len(u.ID) > 0 {
                return bson.M{"_id": u.ID}
        }
        return bson.M{"username": u.Username}
}

// register "models"

func init() {
	db.Register(&User{})
}
