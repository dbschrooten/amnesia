package db

import (
	"log"

	"github.com/asdine/storm/v3"
	"github.com/spf13/viper"
)

var (
	Conn *storm.DB
)

func Setup() error {
	var err error
	Conn, err = storm.Open(viper.GetString("system.db_path"))

	if err != nil {
		return err
	}

	log.Print("Setup embedded db")

	return nil
}
