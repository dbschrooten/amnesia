package schema

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type CheckEvent struct {
	ID        bson.ObjectId `storm:"id"`
	ServiceID string
	Time      time.Time `storm:"index"`
}
