package server

import (
	"github.com/christian298/metrics-aggegator/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Server handels request
type Server struct {
	Router *mux.Router
	Db     *db.Db
}

// New creates a new server
func New() *Server {
	s := &Server{}

	return s
}

// StartServer starts the server and attaches routes and middleware
func (s *Server) StartServer() {
	s.Router.Use(LoggingMiddleware)

	s.routes()
	s.staticFileRoutes()
	log.Fatal(http.ListenAndServe(":4000", s.Router))
}
