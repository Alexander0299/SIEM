package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims представляет утверждения для JWT.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
