package repository

import "siem-system/internal/model"

var Logs []model.Log
var Users []model.User
var Alerts []model.Alert

func GetLogs() []model.Log {
	return Logs
}

func GetUsers() []model.User {
	return Users
}

func GetAlerts() []model.Alert {
	return Alerts
}
