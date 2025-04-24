package main

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"siem-sistem/internal/handler"

	pb "siem-sistem/internal/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при создании слушателя: %v", err)

	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	h := &handler.SiemHandler{}

	pb.RegisterUserServiceServer(grpcServer, h)
	pb.RegisterAlertServiceServer(grpcServer, h)
	pb.RegisterLogServiceServer(grpcServer, h)

	log.Println("gRPC-сервер запущен на порту :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске gRPC-сервера: %v", err)
	}
}
