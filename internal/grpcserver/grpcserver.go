package grpcserver

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"

	"siem-sistem/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	proto.UnimplementedAlertServiceServer
	proto.UnimplementedLogServiceServer
}

// Users

func (s *Server) CreateUser(ctx context.Context, in *proto.User) (*proto.User, error) {
	file, _ := os.OpenFile("users.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	id := getNextID("users.csv")
	in.Id = int32(id)

	writer.Write([]string{strconv.Itoa(id), in.Login})
	return in, nil
}

func (s *Server) ListUsers(ctx context.Context, _ *proto.Empty) (*proto.UserList, error) {
	file, err := os.Open("users.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	var users []*proto.User
	for i, row := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		users = append(users, &proto.User{Id: int32(id), Login: row[1]})
	}

	return &proto.UserList{Users: users}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *proto.UserID) (*proto.Empty, error) {
	return deleteByID("users.csv", int(in.Id))
}

func (s *Server) GetUser(ctx context.Context, in *proto.UserID) (*proto.User, error) {
	file, err := os.Open("users.csv")
	if err != nil {
		return nil, status.Error(codes.NotFound, "users file not found")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to read users")
	}

	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[0])
		if id == int(in.Id) {
			return &proto.User{
				Id:    in.Id,
				Login: row[1],
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "user not found")
}

func (s *Server) UpdateUser(ctx context.Context, in *proto.User) (*proto.User, error) {

	_, err := s.DeleteUser(ctx, &proto.UserID{Id: in.Id})
	if err != nil {
		return nil, err
	}

	file, _ := os.OpenFile("users.csv", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{strconv.Itoa(int(in.Id)), in.Login})
	writer.Flush()

	return in, nil
}

// Alert

func (s *Server) CreateAlert(ctx context.Context, in *proto.Alert) (*proto.Alert, error) {
	file, _ := os.OpenFile("alerts.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	id := getNextID("alerts.csv")
	in.Id = int32(id)

	writer.Write([]string{strconv.Itoa(id), in.Message})
	return in, nil
}

func (s *Server) ListAlerts(ctx context.Context, _ *proto.Empty) (*proto.AlertList, error) {
	file, err := os.Open("alerts.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	var alerts []*proto.Alert
	for i, row := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		alerts = append(alerts, &proto.Alert{Id: int32(id), Message: row[1]})
	}

	return &proto.AlertList{Alerts: alerts}, nil
}

func (s *Server) DeleteAlert(ctx context.Context, in *proto.AlertID) (*proto.Empty, error) {
	return deleteByID("alerts.csv", int(in.Id))
}

func (s *Server) GetAlert(ctx context.Context, in *proto.AlertID) (*proto.Alert, error) {
	file, err := os.Open("alerts.csv")
	if err != nil {
		return nil, status.Error(codes.NotFound, "alerts file not found")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to read alerts")
	}

	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[0])
		if id == int(in.Id) {
			return &proto.Alert{
				Id:      in.Id,
				Message: row[1],
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "log not found")
}

func (s *Server) UpdateAlert(ctx context.Context, in *proto.Alert) (*proto.Alert, error) {

	_, err := s.DeleteAlert(ctx, &proto.AlertID{Id: in.Id})
	if err != nil {
		return nil, err
	}

	file, _ := os.OpenFile("alerts.csv", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{strconv.Itoa(int(in.Id)), in.Message})
	writer.Flush()

	return in, nil
}

// Log

func (s *Server) CreateLog(ctx context.Context, in *proto.Log) (*proto.Log, error) {
	file, _ := os.OpenFile("logs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	id := getNextID("logs.csv")
	in.Id = int32(id)

	writer.Write([]string{strconv.Itoa(id), in.Area})
	return in, nil
}

func (s *Server) ListLogs(ctx context.Context, _ *proto.Empty) (*proto.LogList, error) {
	file, err := os.Open("logs.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	var logs []*proto.Log
	for i, row := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		logs = append(logs, &proto.Log{Id: int32(id), Area: row[1]})
	}

	return &proto.LogList{Logs: logs}, nil
}

func (s *Server) DeleteLog(ctx context.Context, in *proto.LogID) (*proto.Empty, error) {
	return deleteByID("logs.csv", int(in.Id))
}
func (s *Server) GetLog(ctx context.Context, in *proto.LogID) (*proto.Log, error) {
	file, err := os.Open("logs.csv")
	if err != nil {
		return nil, status.Error(codes.NotFound, "logs file not found")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to read logs")
	}

	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[0])
		if id == int(in.Id) {
			return &proto.Log{
				Id:   in.Id,
				Area: row[1],
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "log not found")
}

func (s *Server) UpdateLog(ctx context.Context, in *proto.Log) (*proto.Log, error) {

	_, err := s.DeleteLog(ctx, &proto.LogID{Id: in.Id})
	if err != nil {
		return nil, err
	}

	file, _ := os.OpenFile("logs.csv", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{strconv.Itoa(int(in.Id)), in.Area})
	writer.Flush()

	return in, nil
}

// доп.функции

func getNextID(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	maxID := 0
	for i, row := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

func deleteByID(path string, targetID int) (*proto.Empty, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	newRows := [][]string{rows[0]}

	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[0])
		if id != targetID {
			newRows = append(newRows, row)
		}
	}

	file, _ = os.Create(path)
	writer := csv.NewWriter(file)
	writer.WriteAll(newRows)
	writer.Flush()

	return &proto.Empty{}, nil
}
