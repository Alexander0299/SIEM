package app

import (
	"net/http"
	"siem-sistem/internal/handler"
	"siem-sistem/internal/repository"
	"siem-sistem/internal/service"
)

type App struct {
	repo  *repository.Repository
	svc   *service.Service
	hndlr *handler.Handler
}

func NewApp(repo *repository.Repository, svc *service.Service, hndlr *handler.Handler) *App {
	return &App{repo: repo, svc: svc, hndlr: hndlr}
}

func (a *App) ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
