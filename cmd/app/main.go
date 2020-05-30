package main

import (
	"fmt"
	"os"

	"github.com/christian298/metrics-aggegator/config"
	"github.com/christian298/metrics-aggegator/db"
	"github.com/christian298/metrics-aggegator/server"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	config := config.ReadConfig()

	r := mux.NewRouter()

	srv := server.New()

	influxDB, err := db.New(&config)

	if err != nil {
		return fmt.Errorf("error creating DB: %w", err)
	}

	srv.Router = r
	srv.Db = influxDB
	srv.Config = &config

	srv.StartServer()

	return nil
}
