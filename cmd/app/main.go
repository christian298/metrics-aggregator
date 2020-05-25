package main

import (
	"fmt"
	"github.com/christian298/metrics-aggegator/db"
	"github.com/christian298/metrics-aggegator/server"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	r := mux.NewRouter()

	srv := server.New()

	influxDB, err := db.New()

	if err != nil {
		return fmt.Errorf("error creating DB: %w", err)
	}

	srv.Router = r
	srv.Db = influxDB

	srv.StartServer()

	return nil
}
