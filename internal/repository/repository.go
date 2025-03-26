package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"siem-system/internal/model"
	"strconv"
	"time"
)

type Repository struct {
	logFile   string
	userFile  string
	alertFile string
	logs      []model.Log
	users     []model.User
	alerts    []model.Alert
}

func NewRepository(logFile, userFile, alertFile string) *Repository {
	return &Repository{
		logFile:   logFile,
		userFile:  userFile,
		alertFile: alertFile,
		logs:      []model.Log{},
		users:     []model.User{},
		alerts:    []model.Alert{},
	}
}

func (r *Repository) Load() error {

	if err := r.loadLogs(); err != nil {
		return err
	}

	if err := r.loadUsers(); err != nil {
		return err
	}

	if err := r.loadAlerts(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) loadLogs() error {
	file, err := os.Open(r.logFile)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла логов: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("ошибка чтения CSV: %v", err)
	}
	if len(records) > 0 {
		records = records[1:]
	}
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("ошибка парсинга ID: %v", err)
		}

		timeValue, err := time.Parse("2006-01-02 15:04:05", record[2])
		if err != nil {
			return fmt.Errorf("ошибка парсинга времени: %v", err)
		}

		logEntry := model.Log{
			ID:        id,
			Message:   record[1],
			Timestamp: timeValue,
		}

		r.logs = append(r.logs, logEntry)
	}

	fmt.Println("Логи загружены успешно.")
	return nil
}

func (r *Repository) loadUsers() error {
	file, err := os.Open(r.userFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		r.users = append(r.users, model.User{
			ID:       id,
			Username: record[1],
			Email:    record[2],
		})
	}
	fmt.Println("Пользователи загружены")
	return nil
}

func (r *Repository) loadAlerts() error {
	file, err := os.Open(r.alertFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		r.alerts = append(r.alerts, model.Alert{
			ID:      id,
			Details: record[1],
			Level:   record[2],
		})
	}
	fmt.Println("Алерты загружены")
	return nil
}
