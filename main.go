package main

import (
	"os"
	"os/signal"
	"siem-system/internal/model"
	"siem-system/internal/repository"
	"siem-system/internal/service"
	"syscall"
)

func main() {
	logCh := make(chan model.Log)
	userCh := make(chan model.User)
	alertCh := make(chan model.Alert)

	go service.GenerateData(logCh, userCh, alertCh)
	go repository.ProcessLogs(logCh)
	go repository.ProcessUsers(userCh)
	go repository.ProcessAlerts(alertCh)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	close(logCh)
	close(userCh)
	close(alertCh)

	println("Приложение завершено корректно.")
}
