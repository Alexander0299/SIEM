package main

import (
	"fmt"
	"os"
	"os/signal"
	"siem-system/internal/service"
	"syscall"
)

func main() {
	fmt.Println("Запуск SIEM системы...")

	srv := service.NewSIEMService()
	go srv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("\nЗавершаем работу...")
	srv.Stop()
	fmt.Println("SIEM система остановлена.")
}
