package handler

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"siem-sistem/internal/model"

	"github.com/gorilla/mux"
)

// ///// USER
// ///// Чтение пользователей из CSV
func readUsers() ([]model.User, error) {
	file, err := os.Open("users.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var users []model.User

	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) < 1 {
			continue
		}
		// Пропускаем заголовок
		if i == 0 && strings.Contains(strings.ToLower(record[0]), "польз") {
			continue
		}
		users = append(users, model.User{Login: record[0]})
	}
	return users, nil
}

// ////// Сохранение пользователей в CSV
func saveUsers(users []model.User) error {
	file, err := os.Create("users.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Пользователи"})
	for _, u := range users {
		writer.Write([]string{u.Login})
	}
	return nil
}

// ////// POST /api/user
// CreateUser godoc
// @Summary Создать пользователя
// @Description Добавляет нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "Пользователь"
// @Success 201 {string} string "Создано"
// @Failure 400 {string} string "Неверный ввод"
// @Router /api/user [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	users, _ := readUsers()
	users = append(users, user)
	saveUsers(users)

	w.WriteHeader(http.StatusCreated)
}

// //////// GET /api/users
// GetAllUsers godoc
// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Router /api/users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := readUsers()
	if err != nil {
		http.Error(w, "Ошибка чтения CSV", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// //////// GET /api/user/{id}
// GetUserByID godoc
// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} model.User
// @Failure 404 {string} string "Пользователь не найден"
// @Router /api/user/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}
	users, _ := readUsers()
	if id < 0 || id >= len(users) {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(users[id])
}

// PUT /api/user/{id}
// UpdateUser godoc
// @Summary Обновить пользователя
// @Description Обновляет данные пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body model.User true "Пользователь"
// @Success 200 {string} string "Обновлено"
// @Failure 400 {string} string "Неверный ввод"
// @Failure 404 {string} string "Пользователь не найден"
// @Router /api/user/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	users, _ := readUsers()
	if id < 0 || id >= len(users) {
		http.NotFound(w, r)
		return
	}
	users[id] = newUser
	saveUsers(users)
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/user/{id}
// DeleteUser godoc
// @Summary Удалить пользователя
// @Description Удаляет пользователя по ID
// @Tags users
// @Param id path int true "ID пользователя"
// @Success 200 {string} string "Удалено"
// @Failure 404 {string} string "Пользователь не найден"
// @Router /api/user/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	users, _ := readUsers()
	if id < 0 || id >= len(users) {
		http.NotFound(w, r)
		return
	}

	users = append(users[:id], users[id+1:]...)
	saveUsers(users)
	w.WriteHeader(http.StatusNoContent)
}

// ALERT
// Чтение уведомлений из CSV
func readAlerts() ([]model.Alert, error) {
	file, err := os.Open("alerts.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var alerts []model.Alert

	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) < 1 {
			continue
		}
		// Пропускаем заголовок
		if i == 0 && strings.Contains(strings.ToLower(record[0]), "Уведомления") {
			continue
		}
		alerts = append(alerts, model.Alert{Massage: record[0]})
	}
	return alerts, nil
}

// Сохранение пользователей в CSV
func saveAlerts(alerts []model.Alert) error {
	file, err := os.Create("alerts.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Уведомления"})
	for _, a := range alerts {
		writer.Write([]string{a.Massage})
	}
	return nil
}

// POST /api/alert
// CreateAlert godoc
// @Summary Создать уведомление
// @Description Добавляет нового уведомление
// @Tags alerts
// @Accept json
// @Produce json
// @Param alert body model.Alert true "Уведомление"
// @Success 201 {string} string "Создано"
// @Failure 400 {string} string "Неверный ввод"
// @Router /api/alert [post]
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var alert model.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	alerts, _ := readAlerts()
	alerts = append(alerts, alert)
	saveAlerts(alerts)

	w.WriteHeader(http.StatusCreated)
}

// GET /api/alerts
// GetAllAlerts godoc
// @Summary Получить все уведомления
// @Description Возвращает список всех уведомлений
// @Tags alerts
// @Produce json
// @Success 200 {array} model.Alert
// @Router /api/alerts [get]
func GetAllAlerts(w http.ResponseWriter, r *http.Request) {
	alerts, err := readAlerts()
	if err != nil {
		http.Error(w, "Ошибка чтения CSV", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(alerts)
}

// GET /api/alert/{id}
// GetAlertByID godoc
// @Summary Получить уведомление по ID
// @Description Возвращает уведомление по ID
// @Tags alerts
// @Produce json
// @Param id path int true "ID уведомления"
// @Success 200 {object} model.Alert
// @Failure 404 {string} string "Уведомление не найдено"
// @Router /api/alert/{id} [get]
func GetAlertByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}
	alerts, _ := readAlerts()
	if id < 0 || id >= len(alerts) {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(alerts[id])
}

// PUT /api/alert/{id}
// UpdateAlert godoc
// @Summary Обновить уведомление
// @Description Обновляет уведомления по ID
// @Tags alerts
// @Accept json
// @Produce json
// @Param id path int true "ID уведомления"
// @Param alert body model.Alert true "Уведомление"
// @Success 200 {string} string "Обновлено"
// @Failure 400 {string} string "Неверный ввод"
// @Failure 404 {string} string "Уведомление не найдено"
// @Router /api/alert/{id} [put]
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var newAlert model.Alert
	if err := json.NewDecoder(r.Body).Decode(&newAlert); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	alerts, _ := readAlerts()
	if id < 0 || id >= len(alerts) {
		http.NotFound(w, r)
		return
	}
	alerts[id] = newAlert
	saveAlerts(alerts)
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/alert/{id}
// DeleteAlert godoc
// @Summary Удалить уведомление
// @Description Удаляет уведомление по ID
// @Tags alerts
// @Param id path int true "ID уведомления"
// @Success 200 {string} string "Удалено"
// @Failure 404 {string} string "Уведомление не найдено"
// @Router /api/alert/{id} [delete]
func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	alerts, _ := readAlerts()
	if id < 0 || id >= len(alerts) {
		http.NotFound(w, r)
		return
	}

	alerts = append(alerts[:id], alerts[id+1:]...)
	saveAlerts(alerts)
	w.WriteHeader(http.StatusNoContent)
}

// LOG
// Чтение логов из CSV
func readLogs() ([]model.Log, error) {
	file, err := os.Open("logs.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var logs []model.Log

	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) < 1 {
			continue
		}
		// Пропускаем заголовок
		if i == 0 && strings.Contains(strings.ToLower(record[0]), "Логи") {
			continue
		}
		logs = append(logs, model.Log{Area: record[0]})
	}
	return logs, nil
}

// Сохранение логов в CSV
func saveLogs(logs []model.Log) error {
	file, err := os.Create("logs.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Логи"})
	for _, l := range logs {
		writer.Write([]string{l.Area})
	}
	return nil
}

// POST /api/log
// CreateLog godoc
// @Summary Создать лог
// @Description Добавляет новый лог
// @Tags logs
// @Accept json
// @Produce json
// @Param log body model.Log true "Лог"
// @Success 201 {string} string "Создано"
// @Failure 400 {string} string "Неверный ввод"
// @Router /api/log [post]
func CreateLog(w http.ResponseWriter, r *http.Request) {
	var log model.Log
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	logs, _ := readLogs()
	logs = append(logs, log)
	saveLogs(logs)

	w.WriteHeader(http.StatusCreated)
}

// GET /api/logs
// GetAllLogs godoc
// @Summary Получить все логи
// @Description Возвращает список всех логов
// @Tags logs
// @Produce json
// @Success 200 {array} model.Log
// @Router /api/logs [get]
func GetAllLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := readLogs()
	if err != nil {
		http.Error(w, "Ошибка чтения CSV", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(logs)
}

// GET /api/log/{id}
// GetLogByID godoc
// @Summary Получить логи по ID
// @Description Возвращает логи по ID
// @Tags logs
// @Produce json
// @Param id path int true "ID логов"
// @Success 200 {object} model.Log
// @Failure 404 {string} string "Лог не найден"
// @Router /api/log/{id} [get]
func GetLogByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}
	logs, _ := readLogs()
	if id < 0 || id >= len(logs) {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(logs[id])
}

// PUT /api/log/{id}
// UpdateLog godoc
// @Summary Обновить логи
// @Description Обновляет логи по ID
// @Tags logs
// @Accept json
// @Produce json
// @Param id path int true "ID лога"
// @Param log body model.Log true "Лог"
// @Success 200 {string} string "Обновлено"
// @Failure 400 {string} string "Неверный ввод"
// @Failure 404 {string} string "Лог не найдено"
// @Router /api/log/{id} [put]
func UpdateLog(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var newLog model.Log
	if err := json.NewDecoder(r.Body).Decode(&newLog); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	logs, _ := readLogs()
	if id < 0 || id >= len(logs) {
		http.NotFound(w, r)
		return
	}
	logs[id] = newLog
	saveLogs(logs)
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/log/{id}
// DeleteLog godoc
// @Summary Удалить лог
// @Description Удаляет лог по ID
// @Tags logs
// @Param id path int true "ID лога"
// @Success 200 {string} string "Удалено"
// @Failure 404 {string} string "Лог не найдено"
// @Router /api/log/{id} [delete]
func DeleteLog(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	logs, _ := readLogs()
	if id < 0 || id >= len(logs) {
		http.NotFound(w, r)
		return
	}

	logs = append(logs[:id], logs[id+1:]...)
	saveLogs(logs)
	w.WriteHeader(http.StatusNoContent)
}
