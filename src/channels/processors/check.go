package processors

import (
	"amnesia/src/db"
	"amnesia/src/db/schema"
	"amnesia/src/lib"
	"log"

	"github.com/google/uuid"
)

func CheckEvent(e lib.ChannelEvent) error {
	log.Print("Processing status event")
	event := schema.CheckEvent{
		ID:        uuid.New().String(),
		ServiceID: e.ID,
		Time:      e.Time,
	}

	if err := db.Conn.Exec(
		`INSERT INTO check_events VALUES ?`,
		&event,
	); err != nil {
		return err
	}

	log.Print("Status event processed")

	return nil
}
