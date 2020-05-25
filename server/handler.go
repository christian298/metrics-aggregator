package server

import (
	"encoding/json"
	"fmt"
	"github.com/christian298/metrics-aggegator/models"
	"github.com/mssola/user_agent"
	"net/http"
)

func (s *Server) handleMetrics() http.HandlerFunc {
	var perfMetrics models.Metric

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "localhost")
		if r.Method == http.MethodOptions {
			return
		}

		uaString := r.Header.Get("User-Agent")
		ua := user_agent.New(uaString)
		name, version := ua.Browser()
		platform := ua.Platform()

		if err := json.NewDecoder(r.Body).Decode(&perfMetrics); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		browser := models.Browser{
			Name:     name,
			Version:  version,
			Platform: platform,
		}

		fmt.Println(perfMetrics)

		s.Db.Insert(perfMetrics, browser)

		w.WriteHeader(http.StatusOK)
	}
}
