package model

type Entity interface {
	GetType() string
}

type Alert struct {
	Message string
}

func (a Alert) GetType() string {
	return "Alert"
}

type Log struct {
	Content string
}

func (l Log) GetType() string {
	return "Log"
}
