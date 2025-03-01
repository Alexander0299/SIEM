package service

import (
	"fmt"
	"siem-system/internal/model"
)

func ProcessData(logCh chan model.Log, userCh chan model.User, alertCh chan model.Alert) {
	go func() {
		for log := range logCh {
			fmt.Println("Обработан лог:", log.Content)
		}
	}()

	go func() {
		for user := range userCh {
			fmt.Println("Обработан пользователь:", user.Name)
		}
	}()

	go func() {
		for alert := range alertCh {
			fmt.Println("Обработано предупреждение:", alert.Message)
		}
	}()
}
