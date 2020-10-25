package api

import (
	"amnesia/src/db"
	"amnesia/src/db/schema"
	"amnesia/src/helpers"
	"log"
	"net/http"

	"github.com/genjidb/genji/document"
)

const (
	DefaultLimit = 10
)

func GetCheckEvents(w http.ResponseWriter, r *http.Request) {
	res, err := db.Conn.Query("SELECT * FROM check_events ORDER BY time DESC LIMIT 10")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result []schema.CheckEvent

	if err := res.Iterate(func(d document.Document) error {
		var doc schema.CheckEvent

		if err := document.StructScan(d, &doc); err != nil {
			return err
		}

		result = append(result, doc)

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := helpers.JSONResponse(w, map[string]interface{}{
		"Status": "success",
		"Result": result,
	}, map[string]interface{}{
		"PrettyPrint": true,
	}); err != nil {
		log.Fatal(err)
	}
}

func GetServiceEvents(w http.ResponseWriter, r *http.Request) {
	res, err := db.Conn.Query("SELECT * FROM service_events ORDER BY time DESC LIMIT 10")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result []schema.ServiceEvent

	if err := res.Iterate(func(d document.Document) error {
		var doc schema.ServiceEvent

		if err := document.StructScan(d, &doc); err != nil {
			return err
		}

		result = append(result, doc)

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := helpers.JSONResponse(w, map[string]interface{}{
		"Status": "success",
		"Result": result,
	}, map[string]interface{}{
		"PrettyPrint": true,
	}); err != nil {
		log.Fatal(err)
	}
}
