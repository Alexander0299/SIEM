package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"siem-sistem/internal/app"
	"siem-sistem/internal/model"
	"siem-sistem/internal/service"
	"syscall"
)

func main() {
	// Инициализация CSV-данных
	service.SaveUsersCsv([]model.User{{Login: "Alex"}}, "users.csv")
	service.SaveAlertsCsv([]model.Alert{{Massage: "Попытка взлома"}}, "alerts.csv")
	service.SaveLogsCsv([]model.Log{{Area: "Антивирус Касперского"}}, "logs.csv")

	// Каналы
	usersChan := make(chan model.Inter)
	alertsChan := make(chan model.Inter)
	logsChan := make(chan model.Inter)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запуск воркеров
	go service.AddUsers(ctx, usersChan)
	go service.AddAlerts(ctx, alertsChan)
	go service.AddLogs(ctx, logsChan)
	go service.Logger(usersChan, alertsChan, logsChan)

	// Запуск HTTP-сервера в горутине
	srv, err := app.NewService(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("server stopped: %v", err)
			cancel()
		}
	}()

	// сигнал завершения
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	fmt.Println("Получен сигнал завершения...")

	cancel() // корректное завершение

	fmt.Println("Программа завершена")
}
