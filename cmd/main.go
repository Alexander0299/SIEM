package main

import (
	"log"
	"siem-sistem/internal/app"
	"siem-sistem/internal/handler"
	"siem-sistem/internal/repository"
	"siem-sistem/internal/service"
)

func main() {
	dataSource := "internal/repository/items.csv"

	repo := repository.NewRepository(dataSource)
	svc := service.NewService(repo)
	hndlr := handler.NewHandler(svc)

	application := app.NewApp(repo, svc, hndlr)

	log.Println("Приложение успешно запущено:", application)

	handler := hndlr.InitRoutes()

	log.Println("Сервер запущен на :8080")
	if err := application.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Ошибка сервера: %s", err.Error())
	}
}
