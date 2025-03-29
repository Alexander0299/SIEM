package main

import (
	"log"
	"net/http"
	"siem-system/internal/handlers"
	"siem-system/internal/repository"
)

func main() {

	repo := repository.NewRepository("logs.csv")

	http.HandleFunc("/api/logs", handlers.GetLogs(repo))
	http.HandleFunc("/api/log/", handlers.GetLogByID(repo))

	log.Println("Сервер работает на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
