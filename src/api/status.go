package api

import (
	"amnesia/src/helpers"
	"amnesia/src/status"
	"log"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	status.Current.Check()

	if err := helpers.JSONResponse(w, map[string]interface{}{
		"Status": "success",
		"Result": status.Current,
	}, map[string]interface{}{
		"PrettyPrint": true,
	}); err != nil {
		log.Fatal(err)
	}
}
