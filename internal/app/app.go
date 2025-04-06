package app

import (
	"fmt"
	"net/http"
	"siem-sistem/internal/config"
	"siem-sistem/internal/handler"
	"siem-sistem/internal/routes"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Start() error {
	mux := http.NewServeMux()

	h := handler.NewHandler(a.cfg)
	routes.RegisterRoutes(mux, h)

	serverHTTP := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: mux,
	}

	fmt.Printf("Сервер запущен на %s:%d\n", a.cfg.Host, a.cfg.Port)

	return serverHTTP.ListenAndServe()
}