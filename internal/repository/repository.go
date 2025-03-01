package repository

import "siem-system/internal/model"

var Logs []model.Log
var Users []model.User
var Alerts []model.Alert

func SaveLog(log model.Log) {
	Logs = append(Logs, log)
}

func SaveUser(user model.User) {
	Users = append(Users, user)
}

func SaveAlert(alert model.Alert) {
	Alerts = append(Alerts, alert)
}
