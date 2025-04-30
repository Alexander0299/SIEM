package service

import (
	"context"
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"siem-sistem/internal/model"
)

func createTempIDFile(t *testing.T, entity string, id int) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, entity+"_id.txt")
	os.MkdirAll(filepath.Dir(path), 0755)
	os.WriteFile(path, []byte(strconv.Itoa(id)), 0644)
	return path
}

func TestGetNextID(t *testing.T) {
	entity := "test_entity"
	path := createTempIDFile(t, entity, 10)
	os.MkdirAll("restartcsv", 0755)
	os.Rename(path, IdFilePath(entity))
	defer os.Remove(IdFilePath(entity))

	next := GetNextID(entity)
	if next != 11 {
		t.Errorf("expected next ID 11, got %d", next)
	}
}

func TestRewriteUsersCSV(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "test_users.csv")
	users := []model.User{{ID: 1, Login: "testuser"}}
	err := RewriteUsersCSV(users, filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, _ := os.ReadFile(filename)
	if !strings.Contains(string(data), "testuser") {
		t.Errorf("expected contents to include 'testuser'")
	}
}

func TestRewriteAlertsCSV(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "test_alerts.csv")
	alerts := []model.Alert{{ID: 1, Massage: "alert1"}}
	err := RewriteAlertsCSV(alerts, filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, _ := os.ReadFile(filename)
	if !strings.Contains(string(data), "alert1") {
		t.Errorf("expected contents to include 'alert1'")
	}
}

func TestRewriteLogsCSV(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "test_logs.csv")
	logs := []model.Log{{ID: 1, Area: "log1"}}
	err := RewriteLogsCSV(logs, filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, _ := os.ReadFile(filename)
	if !strings.Contains(string(data), "log1") {
		t.Errorf("expected contents to include 'log1'")
	}
}

func TestLoadUsersFromCSV(t *testing.T) {
	file := filepath.Join(t.TempDir(), "users.csv")
	os.WriteFile(file, []byte("ID,Пользователи:\n1,testuser\n"), 0644)
	users := LoadUsersFromCSV(file)
	if len(users) != 1 || users[0].Login != "testuser" {
		t.Errorf("expected user 'testuser', got: %+v", users)
	}
}

func TestLoadAlertsFromCSV(t *testing.T) {
	file := filepath.Join(t.TempDir(), "alerts.csv")
	os.WriteFile(file, []byte("ID,Уведомления:\n1,alert1\n"), 0644)
	alerts := LoadAlertsFromCSV(file)
	if len(alerts) != 1 || alerts[0].Massage != "alert1" {
		t.Errorf("expected alert 'alert1', got: %+v", alerts)
	}
}

func TestLoadLogsFromCSV(t *testing.T) {
	file := filepath.Join(t.TempDir(), "logs.csv")
	os.WriteFile(file, []byte("ID,Источники:\n1,log1\n"), 0644)
	logs := LoadLogsFromCSV(file)
	if len(logs) != 1 || logs[0].Area != "log1" {
		t.Errorf("expected log 'log1', got: %+v", logs)
	}
}

func TestSaveCsv_HeaderWrite(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "header_test.csv")
	err := SaveCsv(filename, []string{"A", "B"}, func(w *csv.Writer) error {
		return w.Write([]string{"1", "2"})
	})
	if err != nil {
		t.Fatalf("SaveCsv failed: %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("cannot read csv: %v", err)
	}
	if !strings.Contains(string(data), "A,B") {
		t.Errorf("expected header A,B in file, got: %s", string(data))
	}
}

func TestAddUsers_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan model.Inter, 1)

	go AddUsers(ctx, ch)
	time.Sleep(21 * time.Second)
	cancel()
	time.Sleep(21 * time.Second)

	select {
	case u := <-ch:
		if _, ok := u.(model.User); !ok {
			t.Errorf("expected model.User, got %T", u)
		}
	default:
		t.Errorf("no user added to channel")
	}
}

func TestAddAlerts_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan model.Inter, 1)

	go AddAlerts(ctx, ch)
	time.Sleep(21 * time.Second)
	cancel()
	time.Sleep(21 * time.Second)

	select {
	case a := <-ch:
		if _, ok := a.(model.Alert); !ok {
			t.Errorf("expected model.Alert, got %T", a)
		}
	default:
		t.Errorf("no alert added to channel")
	}
}

func TestAddLogs_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan model.Inter, 1)

	go AddLogs(ctx, ch)
	time.Sleep(21 * time.Second)
	cancel()
	time.Sleep(21 * time.Second)

	select {
	case l := <-ch:
		if _, ok := l.(model.Log); !ok {
			t.Errorf("expected model.Log, got %T", l)
		}
	default:
		t.Errorf("no log added to channel")
	}
}
