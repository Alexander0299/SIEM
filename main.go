package main

import (
	"log"
	"os"
	_ "siem-sistem/docsn"
	"siem-sistem/internal/config"
	"siem-sistem/internal/routes"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Your Project API
// @version         1.0
// @description     Это сервер API для вашего проекта.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите "Bearer " перед вашим токеном доступа.
func main() {

	config.LoadConfig()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetConfig().ServerPort
	log.Printf("Сервер запущен на порту %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %s", err)
	}
}
func LoadConfig() config.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env не найден, используем переменные окружения")
	}

	return config.Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
		Login:     os.Getenv("LOGIN"),
		Password:  os.Getenv("PASSWORD"),
	}
}
