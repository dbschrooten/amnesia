package cmd

import (
	"amnesia/src/channels"
	"amnesia/src/config"
	"amnesia/src/dispatch"
	"amnesia/src/extension"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r *mux.Router
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"Status": "ok",
		"Services": map[string]interface{}{
			"Elasticsearch": "ok",
		},
	}); err != nil {
		log.Fatal(err)
	}
}

func Routes() *http.Server {
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/health", Health)

	return &http.Server{
		Handler:      r,
		Addr:         config.SrvAddr,
		WriteTimeout: config.SrvWriteTimeout,
		ReadTimeout:  config.SrvReadTimeout,
	}
}

func Server() error {
	config.Load()

	if err := extension.Setup(); err != nil {
		return err
	}

	if err := config.Setup(); err != nil {
		return err
	}

	channels.Setup()

	srv := Routes()

	go dispatch.Setup()

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
