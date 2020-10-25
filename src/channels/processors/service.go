package processors

import (
	"amnesia/src/db"
	"amnesia/src/db/schema"
	"amnesia/src/lib"

	"github.com/google/uuid"
)

func ServiceEvent(e lib.ChannelEvent) error {
	event := schema.ServiceEvent{
		ID:        uuid.New().String(),
		ServiceID: e.ID,
		Time:      e.Time,
	}

	if err := db.Conn.Exec(
		`INSERT INTO service_events VALUES ?`,
		&event,
	); err != nil {
		return err
	}

	return nil
}
