package main

import (
	"log"
	"net"
	"user-service/habit-service/db"
	pb "user-service/habit-service/proto"
	"user-service/habit-service/server"

	"google.golang.org/grpc"
)

func main() {
	dbURL := "postgres://postgres:database.accesslog@localhost:5432/habit_tracker"

	repo, err := db.NewHabitRepository(dbURL)

	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	log.Print("Database connected successfully....")

	grpcServer := grpc.NewServer()

	habitServer := server.NewHabitServer(repo)

	pb.RegisterHabitServiceServer(grpcServer, habitServer)

	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	log.Print("Server is running on port: 50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

}
