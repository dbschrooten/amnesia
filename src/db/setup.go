package db

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v2"
	"github.com/genjidb/genji"
	"github.com/genjidb/genji/engine/badgerengine"
	"github.com/spf13/viper"
)

const (
	TableServiceEvents = "service_events"
	TableCheckEvents   = "check_events"
	TableAlertEvents   = "alert_events"
)

var (
	Conn *genji.DB
)

func CreateTables() error {
	if err := Conn.Exec(fmt.Sprintf("CREATE TABLE %s", TableServiceEvents)); err != nil {
		if err.Error() != "table already exists" {
			return err
		}
	}

	if err := Conn.Exec(fmt.Sprintf("CREATE TABLE %s", TableCheckEvents)); err != nil {
		if err.Error() != "table already exists" {
			return err
		}
	}

	if err := Conn.Exec(fmt.Sprintf("CREATE TABLE %s", TableAlertEvents)); err != nil {
		if err.Error() != "table already exists" {
			return err
		}
	}

	return nil
}

func Setup() error {
	var (
		err error
		ng  *badgerengine.Engine
	)

	ng, err = badgerengine.NewEngine(
		badger.DefaultOptions(viper.GetString("system.db_path")),
	)

	if err != nil {
		return err
	}

	Conn, err = genji.New(ng)

	if err != nil {
		return err
	}

	if err := CreateTables(); err != nil {
		return err
	}

	log.Print("Setup embedded db")

	return nil
}
