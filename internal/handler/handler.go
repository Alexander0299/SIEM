package handler

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"siem-sistem/internal/model"
	pb "siem-sistem/internal/proto"
	"siem-sistem/internal/service"
	"strconv"
	"strings"

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

// Grpc

type SiemHandler struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedAlertServiceServer
	pb.UnimplementedLogServiceServer
}

// --- USERS ---

func (s *SiemHandler) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	id := service.GetNextID("user")
	user := model.User{
		ID:    id,
		Login: req.Login,
	}
	err := service.RewriteUsers([]model.User{user}, "users.csv")
	if err != nil {
		return nil, err
	}
	return &pb.User{Id: int32(id), Login: req.Login}, nil
}

func (s *SiemHandler) GetUser(ctx context.Context, req *pb.UserID) (*pb.User, error) {
	users := service.LoadUsersFrom("users.csv")
	for _, user := range users {
		if int32(user.ID) == req.Id {
			return &pb.User{Id: int32(user.ID), Login: user.Login}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *SiemHandler) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	users := service.LoadUsersFrom("users.csv")

	updated := false
	for i := range users {
		if int32(users[i].ID) == req.Id {
			users[i].Login = req.Login
			updated = true
			break
		}
	}

	if !updated {
		return nil, fmt.Errorf("user with ID %d not found", req.Id)
	}

	os.Remove("users.csv")

	err := service.RewriteUsers(users, "users.csv")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *SiemHandler) DeleteUser(ctx context.Context, req *pb.UserID) (*pb.Empty, error) {
	users := service.LoadUsersFrom("users.csv")

	filtered := []model.User{}
	for _, user := range users {
		if int32(user.ID) != req.Id {
			filtered = append(filtered, user)
		}
	}

	os.Remove("users.csv")

	err := service.RewriteUsers(filtered, "users.csv")
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
func (s *SiemHandler) ListUsers(ctx context.Context, req *pb.Empty) (*pb.UserList, error) {
	users := service.LoadUsersFrom("users.csv")
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{Id: int32(user.ID), Login: user.Login})
	}
	return &pb.UserList{Users: pbUsers}, nil
}

// --- ALERTS ---

func (s *SiemHandler) CreateAlert(ctx context.Context, req *pb.Alert) (*pb.Alert, error) {
	id := service.GetNextID("alert")
	alert := model.Alert{
		ID:      id,
		Massage: req.Message,
	}
	err := service.RewriteAlerts([]model.Alert{alert}, "alerts.csv")
	if err != nil {
		return nil, err
	}
	return &pb.Alert{Id: int32(id), Message: req.Message}, nil
}

func (s *SiemHandler) GetAlert(ctx context.Context, req *pb.AlertID) (*pb.Alert, error) {
	alerts := service.LoadAlertsFrom("alerts.csv")
	for _, alert := range alerts {
		if int32(alert.ID) == req.Id {
			return &pb.Alert{Id: int32(alert.ID), Message: alert.Massage}, nil
		}
	}
	return nil, fmt.Errorf("alert not found")
}

func (s *SiemHandler) UpdateAlert(ctx context.Context, req *pb.Alert) (*pb.Alert, error) {
	alerts := service.LoadAlertsFrom("alerts.csv")

	updated := false
	for i := range alerts {
		if int32(alerts[i].ID) == req.Id {
			alerts[i].Massage = req.Message
			updated = true
			break
		}
	}

	if !updated {
		return nil, fmt.Errorf("alert with ID %d not found", req.Id)
	}

	os.Remove("alerts.csv")

	err := service.RewriteAlerts(alerts, "alerts.csv")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *SiemHandler) DeleteAlert(ctx context.Context, req *pb.AlertID) (*pb.Empty, error) {
	alerts := service.LoadAlertsFrom("alerts.csv")

	filtered := []model.Alert{}
	for _, alert := range alerts {
		if int32(alert.ID) != req.Id {
			filtered = append(filtered, alert)
		}
	}

	os.Remove("alerts.csv")

	err := service.RewriteAlerts(filtered, "alerts.csv")
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
func (s *SiemHandler) ListAlerts(ctx context.Context, req *pb.Empty) (*pb.AlertList, error) {
	alerts := service.LoadAlertsFrom("alerts.csv")
	var pbAlerts []*pb.Alert
	for _, alert := range alerts {
		pbAlerts = append(pbAlerts, &pb.Alert{Id: int32(alert.ID), Message: alert.Massage})
	}
	return &pb.AlertList{Alerts: pbAlerts}, nil
}

// --- LOGS ---

func (s *SiemHandler) CreateLog(ctx context.Context, req *pb.Log) (*pb.Log, error) {
	id := service.GetNextID("log")
	logItem := model.Log{
		ID:   id,
		Area: req.Area,
	}
	err := service.RewriteLogs([]model.Log{logItem}, "logs.csv")
	if err != nil {
		return nil, err
	}
	return &pb.Log{Id: int32(id), Area: req.Area}, nil
}

func (s *SiemHandler) GetLog(ctx context.Context, req *pb.LogID) (*pb.Log, error) {
	logs := service.LoadLogsFrom("logs.csv")
	for _, logItem := range logs {
		if int32(logItem.ID) == req.Id {
			return &pb.Log{Id: int32(logItem.ID), Area: logItem.Area}, nil
		}
	}
	return nil, fmt.Errorf("log not found")
}

func (s *SiemHandler) UpdateLog(ctx context.Context, req *pb.Log) (*pb.Log, error) {
	logs := service.LoadLogsFrom("logs.csv")

	updated := false
	for i := range logs {
		if int32(logs[i].ID) == req.Id {
			logs[i].Area = req.Area
			updated = true
			break
		}
	}

	if !updated {
		return nil, fmt.Errorf("log with ID %d not found", req.Id)
	}

	os.Remove("logs.csv")

	err := service.RewriteLogs(logs, "logs.csv")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *SiemHandler) DeleteLog(ctx context.Context, req *pb.LogID) (*pb.Empty, error) {
	logs := service.LoadLogsFrom("logs.csv")

	filtered := []model.Log{}
	for _, log := range logs {
		if int32(log.ID) != req.Id {
			filtered = append(filtered, log)
		}
	}

	os.Remove("logs.csv")

	err := service.RewriteLogs(filtered, "logs.csv")
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *SiemHandler) ListLogs(ctx context.Context, req *pb.Empty) (*pb.LogList, error) {
	logs := service.LoadLogsFrom("logs.csv")
	var pbLogs []*pb.Log
	for _, logItem := range logs {
		pbLogs = append(pbLogs, &pb.Log{Id: int32(logItem.ID), Area: logItem.Area})
	}
	return &pb.LogList{Logs: pbLogs}, nil
}
