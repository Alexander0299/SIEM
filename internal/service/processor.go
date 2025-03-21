package service

import "siem-system/internal/repository"

func GetLogs() []string {
	var result []string
	for _, log := range repository.Logs {
		result = append(result, log.Message)
	}
	return result
}

func GetUsers() []string {
	var result []string
	for _, user := range repository.Users {
		result = append(result, user.Name)
	}
	return result
}

func GetAlerts() []string {
	var result []string
	for _, alert := range repository.Alerts {
		result = append(result, alert.Description)
	}
	return result
}
