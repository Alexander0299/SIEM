package main

import (
	"fmt"
	"siem-system/internal/model"
	"siem-system/internal/service"
	"sync"
	"time"
)

func main() {

	logCh := make(chan model.Log, 10)
	userCh := make(chan model.User, 10)
	alertCh := make(chan model.Alert, 10)

	var logMutex, userMutex, alertMutex sync.Mutex

	go service.ProcessLogs(logCh, &logMutex)
	go service.ProcessUsers(userCh, &userMutex)
	go service.ProcessAlerts(alertCh, &alertMutex)

	go service.LogChanges(&logMutex, &userMutex, &alertMutex)

	go service.GenerateData(logCh, userCh, alertCh)

	time.Sleep(10 * time.Second)

	fmt.Println("Программа завершила работу")
}
