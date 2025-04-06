package routes

import (
	"siem-sistem/internal/handler"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", handler.Login)

	protected := r.Group("/")
	protected.Use(handler.AuthMiddleware())
	{
		protected.GET("/protected", handler.ProtectedEndpoint)
	}
}
func RegisterRoutes(mux *http.ServeMux, h *handler.Handler) {
	mux.HandleFunc("/ping", h.Ping)
}
