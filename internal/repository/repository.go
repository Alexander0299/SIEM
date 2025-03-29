package repository

import (
	"encoding/csv"
	"os"
	"siem-system/internal/model"
	"strconv"
	"sync"
)

type Repository struct {
	Logs []model.Log
	mu   sync.Mutex
}

func NewRepository(logFile string) *Repository {
	repo := &Repository{}
	repo.loadLogs(logFile)
	return repo
}

func (r *Repository) loadLogs(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		log := model.Log{ID: id, Message: record[1]}
		r.Logs = append(r.Logs, log)
	}
}

func (r *Repository) AddLog(log model.Log) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Logs = append(r.Logs, log)
}
