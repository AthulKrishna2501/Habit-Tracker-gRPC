package main

import (
	"log"
	"net"
	"user-service/user-service/db"
	pb "user-service/user-service/proto"
	"user-service/user-service/server"

	"google.golang.org/grpc"
)

func main() {
	dbURL := "postgres://postgres:database.accesslog@localhost:5432/habit_tracker"

	repo, err := db.NewUserRepository(dbURL)

	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	log.Print("Database connected successfully...")

	grpcServer := grpc.NewServer()

	userServer := server.NewUserServer(repo)
	pb.RegisterUserServiceServer(grpcServer, userServer)

	listener, err := net.Listen("tcp", ":5001")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	log.Print("Server is running on port : 5001")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

}
