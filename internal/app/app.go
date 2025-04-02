package app

import (
	"log"
	"net/http"

	"siem-sistem/internal/handler"
	"siem-sistem/internal/repository"

	"github.com/gorilla/mux"
)

// App - структура для хранения зависимостей
type App struct {
	Router *mux.Router
	Repo   *repository.Repository
}

// NewApp создаёт и настраивает приложение
func NewApp() *App {
	// Создаём репозиторий
	repo := repository.NewRepository()

	// Загружаем данные (если метод Load нужен)
	if err := repo.Load(); err != nil {
		log.Fatalf("Ошибка загрузки данных: %v", err)
	}

	// Создаём обработчик
	h := handler.NewHandler(repo)

	// Создаём маршрутизатор
	router := mux.NewRouter()
	router.HandleFunc("/api/items", h.GetItems).Methods("GET")
	router.HandleFunc("/api/items/{id:[0-9]+}", h.GetItemByID).Methods("GET")
	router.HandleFunc("/api/items", h.CreateItem).Methods("POST")
	router.HandleFunc("/api/items/{id:[0-9]+}", h.UpdateItem).Methods("PUT")
	router.HandleFunc("/api/items/{id:[0-9]+}", h.DeleteItem).Methods("DELETE")

	return &App{
		Router: router,
		Repo:   repo,
	}
}

// Run запускает сервер
func (a *App) Run(addr string) {
	log.Printf("Сервер запущен на %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
