package handler

import (
	"net/http"
)

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/items", h.GetAllItems)

	return mux
}

func (h *Handler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.Service.GetAllItems()
	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(items))
}