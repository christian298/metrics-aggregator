package server

import (
	"net/http"
)

func (s *Server) routes() {
	s.Router.HandleFunc("/metrics", s.handleMetrics()).Methods(http.MethodPost)
}

func (s *Server) staticFileRoutes() {
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static")))).Methods(http.MethodGet)
	s.Router.PathPrefix("/web").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/index.html")
	})
}
