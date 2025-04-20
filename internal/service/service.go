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

func idFilePath(entity string) string {
	return filepath.Join("restartcsv", fmt.Sprintf("%s_id.txt", entity))
}

func getCurrentID(entity string) int {
	path := idFilePath(entity)
	data, err := os.ReadFile(path)
	if err == nil {
		if val, err := strconv.Atoi(string(data)); err == nil {
			return val
		}
	}
	return 0
}

func getNextID(entity string) int {
	idMutex.Lock()
	defer idMutex.Unlock()
	current := getCurrentID(entity)
	next := current + 1
	os.WriteFile(idFilePath(entity), []byte(strconv.Itoa(next)), 0644)
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
				ID:    getNextID("user"),
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
				ID:      getNextID("alert"),
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
				ID:   getNextID("log"),
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
			if log, ok := item.(model.Log); ok {
				logs = append(logs, log)
				totalLogs++
			}

		case <-ticker.C:
			log.Printf("Количество пользователей=%d, Количество уведомлений=%d, Количество логов=%d",
				totalUsers, totalAlerts, totalLogs)

			if err := SaveUsersCsv(users, "users.csv"); err != nil {
				log.Printf("Ошибка сохранения пользователей: %v", err)
			}
			if err := SaveAlertsCsv(alerts, "alerts.csv"); err != nil {
				log.Printf("Ошибка сохранения уведомлений: %v", err)
			}
			if err := SaveLogsCsv(logs, "logs.csv"); err != nil {
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

func SaveUsersCsv(users []model.User, filename string) error {
	return saveCsv(filename, []string{"ID", "Пользователи:"}, func(w *csv.Writer) error {
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

func SaveAlertsCsv(alerts []model.Alert, filename string) error {
	return saveCsv(filename, []string{"ID", "Уведомления:"}, func(w *csv.Writer) error {
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

func SaveLogsCsv(logs []model.Log, filename string) error {
	return saveCsv(filename, []string{"ID", "Источники:"}, func(w *csv.Writer) error {
		for _, logItem := range logs {
			if err := w.Write([]string{
				strconv.Itoa(logItem.ID),
				logItem.Area,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

// Упрощённая общая логика сохранения CSV
func saveCsv(filename string, header []string, writeRows func(w *csv.Writer) error) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Пишем заголовки, если файл пустой
	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		if err := writer.Write(header); err != nil {
			return err
		}
	}

	return writeRows(writer)
}
