package service

import (
	"fmt"
	"siem-system/internal/model"
)

type Processor struct {
	store *Store
}

func NewProcessor(store *Store) *Processor {
	return &Processor{store: store}
}

func (p *Processor) ProcessLog(log model.Log) {
	fmt.Println("Processing log:", log.Message)
	p.store.Logs = append(p.store.Logs, log)
}

func (p *Processor) ProcessUser(user model.User) {
	fmt.Println("Processing user:", user.Username)
	p.store.Users = append(p.store.Users, user)
}

func (p *Processor) ProcessAlert(alert model.Alert) {
	fmt.Println("Processing alert:", alert.Details)
	p.store.Alerts = append(p.store.Alerts, alert)
}
