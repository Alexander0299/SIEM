package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Log struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

type SIEMService struct {
	logs   []Log
	mu     sync.Mutex
	stopCh chan struct{}
}

func NewSIEMService() *SIEMService {
	s := &SIEMService{
		logs:   []Log{},
		stopCh: make(chan struct{}),
	}
	s.loadLogs()
	return s
}

func (s *SIEMService) Run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.addLog("Система работает...")
		case <-s.stopCh:
			s.saveLogs()
			return
		}
	}
}

func (s *SIEMService) addLog(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log := Log{Timestamp: time.Now().Format(time.RFC3339), Message: message}
	s.logs = append(s.logs, log)
	fmt.Println("Лог добавлен:", message)
}

func (s *SIEMService) Stop() {
	close(s.stopCh)
}

func (s *SIEMService) saveLogs() {
	file, err := os.Create("logs.json")
	if err != nil {
		fmt.Println("Ошибка сохранения JSON:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(s.logs)

	csvFile, err := os.Create("logs.csv")
	if err != nil {
		fmt.Println("Ошибка сохранения CSV:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, log := range s.logs {
		writer.Write([]string{log.Timestamp, log.Message})
	}
}

func (s *SIEMService) loadLogs() {
	file, err := os.Open("logs.json")
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&s.logs)
	}

	csvFile, err := os.Open("logs.csv")
	if err == nil {
		defer csvFile.Close()
		reader := csv.NewReader(csvFile)
		records, _ := reader.ReadAll()
		for _, record := range records {
			s.logs = append(s.logs, Log{Timestamp: record[0], Message: record[1]})
		}
	}
}
