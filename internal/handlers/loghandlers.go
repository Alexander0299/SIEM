package handlers

import (
	"encoding/json"
	"net/http"
	"siem-system/internal/repository"
	"strconv"
	"strings"
)

func GetLogs(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(repo.Logs)
	}
}

func GetLogByID(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/log/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		for _, log := range repo.Logs {
			if log.ID == id {
				json.NewEncoder(w).Encode(log)
				return
			}
		}
		http.NotFound(w, r)
	}
}
