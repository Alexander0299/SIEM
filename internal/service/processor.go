package service

import (
	"fmt"
	"siem-system/internal/model"
	"sync"
)

var (
	logs   []model.Log
	users  []model.User
	alerts []model.Alert
)

func ProcessLogs(logCh <-chan model.Log, mutex *sync.Mutex) {
	for log := range logCh {
		mutex.Lock()
		logs = append(logs, log)
		mutex.Unlock()
		fmt.Println("Обработан лог:", log.Message)
	}
}

func ProcessUsers(userCh <-chan model.User, mutex *sync.Mutex) {
	for user := range userCh {
		mutex.Lock()
		users = append(users, user)
		mutex.Unlock()
		fmt.Println("Обработан пользователь:", user.Username)
	}
}

func ProcessAlerts(alertCh <-chan model.Alert, mutex *sync.Mutex) {
	for alert := range alertCh {
		mutex.Lock()
		alerts = append(alerts, alert)
		mutex.Unlock()
		fmt.Println("Обработано предупреждение:", alert.Details)
	}
}
