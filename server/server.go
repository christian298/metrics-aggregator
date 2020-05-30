package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/christian298/metrics-aggegator/config"
	"github.com/christian298/metrics-aggegator/db"
	"github.com/gorilla/mux"
)

// Server handels request
type Server struct {
	Router *mux.Router
	Db     *db.Db
	Config *config.Config
}

// New creates a new server
func New() *Server {
	s := &Server{}

	return s
}

// StartServer starts the server and attaches routes and middleware
func (s *Server) StartServer() {
	serverAddr := fmt.Sprintf(": %s", s.Config.Server.Port)

	s.routes()
	s.staticFileRoutes()
	log.Fatal(http.ListenAndServe(serverAddr, s.Router))
}
