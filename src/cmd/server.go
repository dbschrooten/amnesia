package cmd

import (
	"amnesia/src/channels"
	"amnesia/src/config"
	"amnesia/src/db"
	"amnesia/src/db/schema"
	"amnesia/src/dispatch"
	"amnesia/src/extension"
	"amnesia/src/status"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

func Records(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var records []schema.CheckEvent

	if err := db.Conn.All(&records); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving records from db"))
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"Records": records,
	}); err != nil {
		log.Fatal(err)
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	status.Current.AddFailure("website_healthcheck", "elasticsearch")
	log.Print(status.Current.GetFailure("website_healthcheck", "elasticsearch"))

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")

	if err := enc.Encode(status.Current); err != nil {
		log.Fatal(err)
	}
}

func Routes() *http.Server {
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/health", Health).Methods("GET")
	r.HandleFunc("/api/v1/records", Records).Methods("GET")
	r.HandleFunc("/api/v1/status", Status).Methods("GET")

	return &http.Server{
		Handler:      r,
		Addr:         viper.GetString("server.address"),
		WriteTimeout: viper.GetDuration("server.write_timeout"),
		ReadTimeout:  viper.GetDuration("server.read_timeout"),
	}
}

func Server() error {
	if err := extension.Setup(); err != nil {
		return err
	}

	if err := config.Setup(); err != nil {
		return err
	}

	if err := status.Setup(); err != nil {
		return err
	}

	if err := db.Setup(); err != nil {
		return err
	}

	defer db.Conn.Close()

	channels.Setup()

	srv := Routes()

	go dispatch.Setup()

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
