package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	Login      string
	Password   string
	JWTSecret  string
	Host       string
	Port       int
}

var cfg Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используем переменные окружения")
	}

	cfg = Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		Host:       os.Getenv("HOST"),
		Port:       mustParseInt(os.Getenv("PORT")),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Login:      os.Getenv("LOGIN"),
		Password:   os.Getenv("PASSWORD"),
	}
}

func GetConfig() Config {
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %v", err)
	}
	return i
}
