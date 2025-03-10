package repository

import (
	"fmt"
	"siem-system/internal/model"
)

func ProcessLogs(logCh <-chan model.Log) {
	for log := range logCh {
		fmt.Println("Log:", log)
	}
}

func ProcessUsers(userCh <-chan model.User) {
	for user := range userCh {
		fmt.Println("User:", user)
	}
}

func ProcessAlerts(alertCh <-chan model.Alert) {
	for alert := range alertCh {
		fmt.Println("Alert:", alert)
	}
}

func LogChanges() {
	fmt.Println("Logging changes...")
}
