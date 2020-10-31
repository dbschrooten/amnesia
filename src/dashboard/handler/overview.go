package handler

import (
	"amnesia/src/dashboard"
	"amnesia/src/status"
	"log"
	"net/http"
)

func Overview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status.Current.Check()

	if err := dashboard.OverviewTpl.ExecuteTemplate(w, "layout.html", status.Current); err != nil {
		log.Fatal(err)
	}
}
