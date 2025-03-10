package service

import (
	"fmt"
	"math/rand"
	"siem-system/internal/model"
	"time"
)

func GenerateData(logCh chan<- model.Log, userCh chan<- model.User, alertCh chan<- model.Alert) {
	for {
		time.Sleep(500 * time.Millisecond)

		logCh <- model.Log{ID: rand.Intn(1000), Message: "New log entry", Timestamp: time.Now()}
		userCh <- model.User{ID: rand.Intn(1000), Username: "User_" + fmt.Sprint(rand.Intn(100)), Email: "user@example.com"}
		alertCh <- model.Alert{ID: rand.Intn(1000), Level: "High", Details: "Security threat detected"}
	}
}
