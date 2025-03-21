package model

type Log struct {
	Message string `json:"message"`
	Level   string `json:"level"`
}

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Alert struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
