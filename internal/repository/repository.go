package service

import (
	"encoding/json"
	"fmt"
	"os"
	"siem-system/internal/model"
)

type Store struct {
	Logs   []model.Log
	Users  []model.User
	Alerts []model.Alert
}

func (s *Store) SaveToFile() {

	file, err := os.Create("logs.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.Encode(s.Logs)

	file, err = os.Create("users.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	enc.Encode(s.Users)

	file, err = os.Create("alerts.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	enc.Encode(s.Alerts)
}

func (s *Store) LoadFromFile() {

	file, err := os.Open("logs.json")
	if err == nil {
		defer file.Close()
		dec := json.NewDecoder(file)
		dec.Decode(&s.Logs)
	}

	file, err = os.Open("users.json")
	if err == nil {
		defer file.Close()
		dec := json.NewDecoder(file)
		dec.Decode(&s.Users)
	}

	file, err = os.Open("alerts.json")
	if err == nil {
		defer file.Close()
		dec := json.NewDecoder(file)
		dec.Decode(&s.Alerts)
	}
}
