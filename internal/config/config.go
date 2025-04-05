package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
	Username  string
	Password  string
	LogFile   string
	UserFile  string
	AlertFile string
	ItemFile  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Username:  os.Getenv("USERNAME"),
		Password:  os.Getenv("PASSWORD"),
		LogFile:   os.Getenv("LOG_FILE"),
		UserFile:  os.Getenv("USER_FILE"),
		AlertFile: os.Getenv("ALERT_FILE"),
		ItemFile:  os.Getenv("ITEM_FILE"),
	}
}
