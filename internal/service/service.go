package service

import (
	"fmt"
	"math/rand"
	"siem-system/internal/model"
	"time"
)

type Store struct {
	Logs   []model.Log
	Users  []model.User
	Alerts []model.Alert
}

func NewStore() *Store {
	return &Store{
		Logs:   make([]model.Log, 0),
		Users:  make([]model.User, 0),
		Alerts: make([]model.Alert, 0),
	}
}

func GenerateData(logCh chan<- model.Log, userCh chan<- model.User, alertCh chan<- model.Alert) {
	for {
		time.Sleep(500 * time.Millisecond)

		logCh <- model.Log{
			ID:        rand.Intn(1000),
			Message:   "New log entry",
			Timestamp: time.Now(),
		}

		userCh <- model.User{
			ID:       rand.Intn(1000),
			Username: fmt.Sprintf("User_%d", rand.Intn(100)),
			Email:    "user@example.com",
		}

		alertCh <- model.Alert{
			ID:      rand.Intn(1000),
			Level:   "High",
			Details: "Potential security threat detected",
		}
	}
}
