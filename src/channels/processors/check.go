package processors

import (
	"amnesia/src/db"
	"amnesia/src/db/schema"
	"amnesia/src/lib"

	"labix.org/v2/mgo/bson"
)

func CheckEvent(e lib.ChannelEvent) error {
	event := schema.CheckEvent{
		ID:        bson.NewObjectId(),
		ServiceID: e.ID,
		Time:      e.Time,
	}

	if err := db.Conn.Save(&event); err != nil {
		return err
	}

	return nil
}
