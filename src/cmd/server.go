package cmd

import (
	"amnesia/src/api"
	"amnesia/src/channels"
	"amnesia/src/config"
	"amnesia/src/dashboard"
	"amnesia/src/dashboard/handler"
	"amnesia/src/db"
	"amnesia/src/dispatch"
	"amnesia/src/extension"
	"amnesia/src/status"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var (
	r *mux.Router
)

func Routes() *http.Server {
	r = mux.NewRouter()

	// Dashboard Routes
	r.HandleFunc("/", handler.Overview).Methods("GET")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("src/dashboard/static"))))

	// REST API routes
	r.HandleFunc("/api/v1/events/check", api.GetCheckEvents).Methods("GET")
	r.HandleFunc("/api/v1/events/service", api.GetServiceEvents).Methods("GET")
	r.HandleFunc("/api/v1/status", api.GetStatus).Methods("GET")

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

	if err := db.Setup(); err != nil {
		return err
	}

	defer db.Conn.Close()

	if err := status.Setup(); err != nil {
		return err
	}

	channels.Setup()

	if err := dashboard.Setup(); err != nil {
		return err
	}

	go dispatch.Setup()

	srv := Routes()

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
