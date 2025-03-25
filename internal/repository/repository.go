package repository

import (
	"encoding/csv"
	"os"
	"siem-system/internal/model"
	"strconv"
	"sync"
)

type Repository struct {
	mu            sync.Mutex
	Logs          []model.Log
	Users         []model.User
	Alerts        []model.Alert
	logFilePath   string
	userFilePath  string
	alertFilePath string
}

func NewRepository(logFile, userFile, alertFile string) *Repository {
	return &Repository{
		logFilePath:   logFile,
		userFilePath:  userFile,
		alertFilePath: alertFile,
	}
}

func (r *Repository) AddLog(log model.Log) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Logs = append(r.Logs, log)
	saveToCSV(r.logFilePath, logsToCSV(r.Logs))
}

func (r *Repository) AddUser(user model.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Users = append(r.Users, user)
	saveToCSV(r.userFilePath, usersToCSV(r.Users))
}

func (r *Repository) AddAlert(alert model.Alert) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Alerts = append(r.Alerts, alert)
	saveToCSV(r.alertFilePath, alertsToCSV(r.Alerts))
}

func logsToCSV(logs []model.Log) [][]string {
	data := [][]string{{"ID", "Message", "Timestamp"}}
	for _, log := range logs {
		data = append(data, []string{
			strconv.Itoa(log.ID),
			log.Message,
			log.Timestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return data
}

func usersToCSV(users []model.User) [][]string {
	data := [][]string{{"ID", "Name", "Role"}}
	for _, user := range users {
		data = append(data, []string{
			strconv.Itoa(user.ID),
			user.Username,
			user.Email,
		})
	}
	return data
}

func alertsToCSV(alerts []model.Alert) [][]string {
	data := [][]string{{"ID", "Details", "Level"}}
	for _, alert := range alerts {
		data = append(data, []string{
			strconv.Itoa(alert.ID),
			alert.Level,
			alert.Details,
		})
	}
	return data
}

func saveToCSV(filePath string, data [][]string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
