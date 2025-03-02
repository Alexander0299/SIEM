package model

import "time"

type Log struct {
	ID        int
	Message   string
	Timestamp time.Time
}

type User struct {
	ID       int
	Username string
	Email    string
}

type Alert struct {
	ID      int
	Level   string
	Details string
}
