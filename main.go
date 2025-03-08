package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"siem-system/internal/model"
	"siem-system/internal/service"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	logCh := make(chan model.Log)
	userCh := make(chan model.User)
	alertCh := make(chan model.Alert)

	go service.ProcessData(ctx, logCh, userCh, alertCh)

	go service.GenerateData(ctx, logCh, userCh, alertCh)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	fmt.Println("\nПолучен сигнал завершения, останавливаем программу...")

	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Программа завершена.")
}
