package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"siem-sistem/internal/model"
	"strconv"
	"sync"
	"time"
)

var (
	idMutex sync.Mutex
)

func IdFilePath(entity string) string {
	return filepath.Join("restartcsv", fmt.Sprintf("%s_id.txt", entity))
}

func GetCurrentID(entity string) int {
	path := IdFilePath(entity)
	data, err := os.ReadFile(path)
	if err == nil {
		if val, err := strconv.Atoi(string(data)); err == nil {
			return val
		}
	}
	return 0
}

func GetNextID(entity string) int {
	idMutex.Lock()
	defer idMutex.Unlock()
	current := GetCurrentID(entity)
	next := current + 1
	os.WriteFile(IdFilePath(entity), []byte(strconv.Itoa(next)), 0644)
	return next
}

func AddUsers(ctx context.Context, usersChan chan model.Inter) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			usersChan <- model.User{
				ID:    GetNextID("user"),
				Login: "Alex",
			}
		}
	}
}

func AddAlerts(ctx context.Context, alertsChan chan model.Inter) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			alertsChan <- model.Alert{
				ID:      GetNextID("alert"),
				Massage: "Попытка взлома",
			}
		}
	}
}

func AddLogs(ctx context.Context, logsChan chan model.Inter) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			logsChan <- model.Log{
				ID:   GetNextID("log"),
				Area: "Антивирус Касперского",
			}
		}
	}
}

func Logger(usersChan, alertsChan, logsChan chan model.Inter) {
	users := []model.User{}
	alerts := []model.Alert{}
	logs := []model.Log{}

	var totalUsers, totalAlerts, totalLogs int

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case item, ok := <-usersChan:
			if !ok {
				usersChan = nil
				continue
			}
			if user, ok := item.(model.User); ok {
				users = append(users, user)
				totalUsers++
			}

		case item, ok := <-alertsChan:
			if !ok {
				alertsChan = nil
				continue
			}
			if alert, ok := item.(model.Alert); ok {
				alerts = append(alerts, alert)
				totalAlerts++
			}

		case item, ok := <-logsChan:
			if !ok {
				logsChan = nil
				continue
			}
			if logItem, ok := item.(model.Log); ok {
				logs = append(logs, logItem)
				totalLogs++
			}

		case <-ticker.C:
			log.Printf("Количество пользователей=%d, Количество уведомлений=%d, Количество логов=%d",
				totalUsers, totalAlerts, totalLogs)

			if err := RewriteUsersCSV(users, "users.csv"); err != nil {
				log.Printf("Ошибка сохранения пользователей: %v", err)
			}
			if err := RewriteAlertsCSV(alerts, "alerts.csv"); err != nil {
				log.Printf("Ошибка сохранения уведомлений: %v", err)
			}
			if err := RewriteLogsCSV(logs, "logs.csv"); err != nil {
				log.Printf("Ошибка сохранения логов: %v", err)
			}

			users = []model.User{}
			alerts = []model.Alert{}
			logs = []model.Log{}
		}

		if usersChan == nil && alertsChan == nil && logsChan == nil {
			return
		}
	}
}

func RewriteUsersCSV(users []model.User, filename string) error {
	return SaveCsv(filename, []string{"ID", "Пользователи:"}, func(w *csv.Writer) error {
		for _, user := range users {
			if err := w.Write([]string{
				strconv.Itoa(user.ID),
				user.Login,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func RewriteAlertsCSV(alerts []model.Alert, filename string) error {
	return SaveCsv(filename, []string{"ID", "Уведомления:"}, func(w *csv.Writer) error {
		for _, alert := range alerts {
			if err := w.Write([]string{
				strconv.Itoa(alert.ID),
				alert.Massage,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func RewriteLogsCSV(logs []model.Log, filename string) error {
	return SaveCsv(filename, []string{"ID", "Источники:"}, func(w *csv.Writer) error {
		for _, log := range logs {
			if err := w.Write([]string{
				strconv.Itoa(log.ID),
				log.Area,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func SaveCsv(filename string, header []string, writeRows func(w *csv.Writer) error) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		if err := writer.Write(header); err != nil {
			return err
		}
	}

	return writeRows(writer)
}

func LoadUsersFromCSV(filename string) []model.User {
	var users []model.User
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Ошибка открытия users CSV:", err)
		return users
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Ошибка чтения users CSV:", err)
		return users
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		users = append(users, model.User{
			ID:    id,
			Login: row[1],
		})
	}
	return users
}

func LoadAlertsFromCSV(filename string) []model.Alert {
	var alerts []model.Alert
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Ошибка открытия alerts CSV:", err)
		return alerts
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Ошибка чтения alerts CSV:", err)
		return alerts
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		alerts = append(alerts, model.Alert{
			ID:      id,
			Massage: row[1],
		})
	}
	return alerts
}

func LoadLogsFromCSV(filename string) []model.Log {
	var logs []model.Log
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Ошибка открытия logs CSV:", err)
		return logs
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Ошибка чтения logs CSV:", err)
		return logs
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		logs = append(logs, model.Log{
			ID:   id,
			Area: row[1],
		})
	}
	return logs
}
