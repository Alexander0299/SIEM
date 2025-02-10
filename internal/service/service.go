package service

import (
	"siem-system/internal/model"
	"siem-system/internal/repository"
	"time"
)

func GenerateAndSendData() {
	for {
		alert := model.Alert{Message: "Security alert detected!"}
		log := model.Log{Content: "User logged in"}

		repository.StoreEntity(alert)
		repository.StoreEntity(log)

		time.Sleep(5 * time.Second)
	}
}
