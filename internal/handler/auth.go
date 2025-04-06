package handler

import (
	"net/http"
	"siem-sistem/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login проверяет учетные данные и возвращает JWT
// @Summary Вход в систему
// @Description Проверка учетных данных и получение JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Учетные данные"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cfg := config.GetConfig()
	if input.Username != cfg.Login || input.Password != cfg.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = input.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// ProtectedEndpoint пример защищенного маршрута
// @Summary Защищенный маршрут
// @Description Доступен только с действительным JWT
// @Tags protected
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "Доступ разрешен"
// @Router /protected [get]
func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Доступ разрешен"})
}
