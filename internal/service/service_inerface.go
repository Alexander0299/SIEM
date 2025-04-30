package service

import (
	"siem-sistem/internal/model"
)

type EntityIDGenerator interface {
	GetNextID(entity string) int
	GetCurrentID(entity string) int
}

type CSVWriter interface {
	RewriteUsersCSV(users []model.User, filename string) error
	RewriteAlertsCSV(alerts []model.Alert, filename string) error
	RewriteLogsCSV(logs []model.Log, filename string) error
}
