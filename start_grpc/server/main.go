package main

import (
	"log"
	"net"
	"siem-sistem/internal/grpcserver"
	"siem-sistem/internal/proto"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}

	s := grpc.NewServer()
	srv := &grpcserver.Server{}
	reflection.Register(s)
	proto.RegisterUserServiceServer(s, srv)
	proto.RegisterAlertServiceServer(s, srv)
	proto.RegisterLogServiceServer(s, srv)

	log.Println("gRPC сервер запущен на порту 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка при запуске gRPC: %v", err)
	}
}
