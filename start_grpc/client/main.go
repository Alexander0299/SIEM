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

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Создание пользователя
	newUser, err := client.CreateUser(ctx, &pb.User{Login: "NewUser"})
	if err != nil {
		log.Fatalf("Ошибка при создании пользователя: %v", err)
	}
	log.Printf("Создан пользователь: ID=%d, Login=%s", newUser.Id, newUser.Login)

	// Получение всех пользователей
	userList, err := client.ListUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Ошибка при получении списка пользователей: %v", err)
	}
	log.Println("Список пользователей:")
	for _, u := range userList.Users {
		log.Printf("ID=%d, Login=%s", u.Id, u.Login)
	}
}
