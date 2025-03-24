package service

import (
	"log"
	"os"
)

func InitializeLogger() *log.Logger {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return log.New(file, "SIEM: ", log.LstdFlags)
}
