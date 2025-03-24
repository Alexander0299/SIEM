package model

import "time"

type Log struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Alert struct {
	ID      int    `json:"id"`
	Level   string `json:"level"`
	Details string `json:"details"`
}
