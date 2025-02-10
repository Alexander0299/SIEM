package repository

import (
	"fmt"
	"siem-system/internal/model"
)

var alerts []model.Alert
var logs []model.Log

func StoreEntity(e model.Entity) {
	switch v := e.(type) {
	case model.Alert:
		alerts = append(alerts, v)
		fmt.Println("New alert stored:", v.Message)
	case model.Log:
		logs = append(logs, v)
		fmt.Println("New log stored:", v.Content)
	default:
		fmt.Println("Unknown type received")
	}
}
