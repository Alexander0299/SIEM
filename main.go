package main

import (
	"fmt"
	"os"
	"os/signal"
	"siem-system/internal/model"
	"siem-system/internal/repository"
	"syscall"
	"time"
)

func main() {

	repo := repository.NewRepository("logs.csv", "users.csv", "alerts.csv")

	repo.AddLog(model.Log{
		ID:        1,
		Message:   "Система запущена",
		Timestamp: time.Now(),
	})

	repo.AddUser(model.User{
		ID:       1,
		Username: "Admin",
		Email:    "Administrator",
	})

	repo.AddAlert(model.Alert{
		ID:      1,
		Details: "Неудачная попытка входа",
		Level:   "High",
	})

	fmt.Println("Данные сохранены.")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	<-sigs
	fmt.Println("\nЗавершаем работу gracefully...")
}
