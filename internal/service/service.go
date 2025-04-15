package service

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"siem-sistem/internal/model"
	"time"
)

func AddUsers(ctx context.Context, usersChan chan model.Inter) {
	for {
		select {
		case <-ctx.Done():
			return
		case usersChan <- model.User{Login: "Alex"}:
			time.Sleep(30 * time.Second)
		}

	}
}
func AddAlerts(ctx context.Context, alertsChan chan model.Inter) {
	for {
		select {
		case <-ctx.Done():
			return
		case alertsChan <- model.Alert{Massage: "Попытка взлома"}:
			time.Sleep(30 * time.Second)
		}

	}
}
func AddLogs(ctx context.Context, logsChan chan model.Inter) {
	for {
		select {
		case <-ctx.Done():
			return
		case logsChan <- model.Log{Area: "Антивирус Касперский"}:
			time.Sleep(30 * time.Second)
		}

	}
}
func Logger(usersChan, alertsChan, logsChan chan model.Inter) {
	users := []model.Inter{}
	alerts := []model.Inter{}
	logs := []model.Inter{}
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case user, ok := <-usersChan:
			if !ok {
				return
			}
			users = append(users, user)

		case alert, ok := <-alertsChan:
			if !ok {
				return
			}
			alerts = append(alerts, alert)

		case log, ok := <-logsChan:
			if !ok {
				return
			}
			logs = append(logs, log)

		case <-ticker.C:
			log.Printf("Количество пользователей: %d,Количество уведомлений: %d, Количество логов: %d\n", len(users), len(alerts), len(logs))
		}
	}
}
func SaveUsersCsv(users []model.User, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		writer.Write([]string{"Пользователи:"})
	}

	for _, user := range users {
		writer.Write([]string{user.Login})
	}
	return nil
}

func SaveAlertsCsv(alerts []model.Alert, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		writer.Write([]string{"Уведомления:"})
	}

	for _, alert := range alerts {
		writer.Write([]string{alert.Massage})
	}
	return nil
}

func SaveLogsCsv(logs []model.Log, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		writer.Write([]string{"Логи:"})
	}

	for _, log := range logs {
		writer.Write([]string{log.Area})
	}
	return nil
}
