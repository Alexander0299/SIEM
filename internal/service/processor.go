package service

import (
	"fmt"
	"siem-sistem/internal/model"
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
func ProcessItem(item model.Item) {
	time.Sleep(time.Second)
	fmt.Println("Processed alert:", item.Name)
}
