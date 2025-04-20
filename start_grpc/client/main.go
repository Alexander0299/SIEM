package main

import (
	"context"
	"log"
	"time"

	pb "siem-sistem/internal/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userClient := pb.NewUserServiceClient(conn)

	// 🧪 Create User
	uResp, err := userClient.CreateUser(ctx, &pb.User{Login: "admin"})
	if err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}
	log.Printf("Создан пользователь: ID=%d, Login=%s", uResp.Id, uResp.Login)

}
