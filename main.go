package main

import (
	"log"
	"net/http"

	"siem-sistem/internal/handler"
	"siem-sistem/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Создаём репозиторий
	repo := repository.NewRepository()

	// Загружаем данные (если метод Load нужен)
	if err := repo.Load(); err != nil {
		log.Fatalf("Ошибка загрузки данных: %v", err)
	}

	// Создаём обработчик с репозиторием
	h := handler.NewHandler(repo)

	// Настраиваем маршруты
	r := mux.NewRouter()
	r.HandleFunc("/api/items", h.GetItems).Methods("GET")
	r.HandleFunc("/api/items/{id:[0-9]+}", h.GetItemByID).Methods("GET")
	r.HandleFunc("/api/items", h.CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id:[0-9]+}", h.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id:[0-9]+}", h.DeleteItem).Methods("DELETE")

	// Запускаем сервер
	log.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
