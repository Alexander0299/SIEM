package service

import (
	"fmt"
	"siem-system/internal/model"
	"time"
)

func ProcessLog(log model.Log) {
	time.Sleep(time.Second)
	fmt.Println("Processed log:", log.Message)
}

func ProcessUser(user model.User) {
	time.Sleep(time.Second)
	fmt.Println("Processed user:", user.Username)
}

func ProcessAlert(alert model.Alert) {
	time.Sleep(time.Second)
	fmt.Println("Processed alert:", alert.Details)
}
