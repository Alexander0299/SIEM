package service

import (
	"fmt"
	"siem-system/internal/repository"
)

func PrintLogs() {
	for _, log := range repository.Logs {
		fmt.Printf("Log: %s, Level: %s\n", log.Message, log.Level)
	}
}
