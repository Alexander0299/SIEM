package handler

import (
	"fmt"
	"net/http"
	"siem-sistem/internal/config"
)

type Handler struct {
	Cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{Cfg: cfg}
}

func (h *Handler) Base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Base handler")
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login handler")
}

func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Update handler")
}
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
