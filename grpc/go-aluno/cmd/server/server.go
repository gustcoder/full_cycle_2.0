package main

import (
	"app/pb"
	"app/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{})
	reflection.Register(grpcServer) // necessário para que o evans client consiga acessar o server

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Server Error: %v", err)
	}
}
