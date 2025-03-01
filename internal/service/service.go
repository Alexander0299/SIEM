package service

import (
	"siem-system/internal/model"
	"time"
)

func GenerateData(logCh chan model.Log, userCh chan model.User, alertCh chan model.Alert) {
	go func() {
		for {
			logCh <- model.Log{Content: "Новый лог"}
			userCh <- model.User{Name: "Александр"}
			alertCh <- model.Alert{Message: "Обнаружена угроза"}
			time.Sleep(1 * time.Second)
		}
	}()
}
