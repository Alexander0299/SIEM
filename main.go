package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"siem-system/internal/model"
	"siem-system/internal/service"
	"syscall"
)

func main() {
	logCh := make(chan model.Log)
	userCh := make(chan model.User)
	alertCh := make(chan model.Alert)

	store := service.NewStore()

	processor := service.NewProcessor(store)

	go func() {
		for {
			select {
			case log := <-logCh:
				processor.ProcessLog(log)
			case user := <-userCh:
				processor.ProcessUser(user)
			case alert := <-alertCh:
				processor.ProcessAlert(alert)
			}
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	fmt.Println("\nПриложение завершено корректно")
}
