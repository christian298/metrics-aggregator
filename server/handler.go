package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/christian298/metrics-aggegator/models"
	"github.com/mssola/user_agent"
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
		isMobile := ua.Mobile()

		if err := json.NewDecoder(r.Body).Decode(&perfMetrics); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		deviceType := "Mobile"
		if !isMobile {
			deviceType = "Desktop"
		}

		browser := models.Browser{
			Name:     name,
			Version:  version,
			Platform: platform,
			Type:     deviceType,
		}

		fmt.Println(perfMetrics)

		s.Db.Insert(perfMetrics, browser)

		w.WriteHeader(http.StatusOK)
	}
}
