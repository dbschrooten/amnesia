package cmd

import (
	"amnesia/src/channels"
	"amnesia/src/config"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func Server() error {
	config.Load()

	if err := config.Setup(); err != nil {
		return err
	}

	channels.Setup()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/health", Health)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.SrvAddr,
		WriteTimeout: config.SrvWriteTimeout,
		ReadTimeout:  config.SrvReadTimeout,
	}

	log.Printf("Listening at %s", config.SrvAddr)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
