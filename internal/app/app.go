package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"siem-sistem/internal/config"
	"siem-sistem/internal/handler"
	"syscall"
	"time"

	_ "siem-sistem/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg *config.Config
	ctx context.Context
}

func NewService(ctx context.Context) (*App, error) {
	return &App{
		ctx: ctx,
		cfg: config.NewConfig(),
	}, nil
}

func (a *App) Start() error {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// user
	router.HandleFunc("/api/users", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", handler.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user", handler.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", handler.DeleteUser).Methods("DELETE")
	//alert
	router.HandleFunc("/api/alerts", handler.GetAllAlerts).Methods("GET")
	router.HandleFunc("/api/alert/{id}", handler.GetAlertByID).Methods("GET")
	router.HandleFunc("/api/alert", handler.CreateAlert).Methods("POST")
	router.HandleFunc("/api/alert/{id}", handler.UpdateAlert).Methods("PUT")
	router.HandleFunc("/api/alert/{id}", handler.DeleteAlert).Methods("DELETE")
	//log
	router.HandleFunc("/api/logs", handler.GetAllLogs).Methods("GET")
	router.HandleFunc("/api/log/{id}", handler.GetLogByID).Methods("GET")
	router.HandleFunc("/api/log", handler.CreateLog).Methods("POST")
	router.HandleFunc("/api/log/{id}", handler.UpdateLog).Methods("PUT")
	router.HandleFunc("/api/log/{id}", handler.DeleteLog).Methods("DELETE")

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: router,
	}

	go func() {
		log.Printf("Сервер запущен на %s:%d\n", a.cfg.Host, a.cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(timeout); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	log.Println("Server gracefully stopped.")
	return nil
}
