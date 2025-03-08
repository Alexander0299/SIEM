package service

import (
	"context"
	"fmt"
	"math/rand"
	"siem-system/internal/model"
	"time"
)

func GenerateData(ctx context.Context, logCh chan<- model.Log, userCh chan<- model.User, alertCh chan<- model.Alert) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Генерация данных остановлена.")
			return
		default:
			time.Sleep(500 * time.Millisecond)
			logCh <- model.Log{
				ID:        rand.Intn(1000),
				Message:   "New log entry",
				Timestamp: time.Now(),
			}
			userCh <- model.User{
				ID:       rand.Intn(1000),
				Username: "User_" + fmt.Sprint(rand.Intn(100)),
				Email:    "user@example.com",
			}
			alertCh <- model.Alert{
				ID:      rand.Intn(1000),
				Level:   "High",
				Details: "Potential security threat detected",
			}
		}
	}
}

func ProcessData(ctx context.Context, logCh <-chan model.Log, userCh <-chan model.User, alertCh <-chan model.Alert) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Обработка данных остановлена.")
			return
		case log := <-logCh:
			fmt.Println("Обработан лог:", log.Message)
		case user := <-userCh:
			fmt.Println("Обработан пользователь:", user.Username)
		case alert := <-alertCh:
			fmt.Println("Обработано предупреждение:", alert.Details)
		}
	}
}
