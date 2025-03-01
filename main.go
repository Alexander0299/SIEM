package main

import (
	"fmt"
	"siem-system/internal/model"
	"siem-system/internal/service"
	"time"
)

func main() {

	logCh := make(chan model.Log)
	userCh := make(chan model.User)
	alertCh := make(chan model.Alert)

	service.ProcessData(logCh, userCh, alertCh)

	go service.GenerateData(logCh, userCh, alertCh)

	time.Sleep(2 * time.Second)

	fmt.Println("Программа завершила работу")
}
