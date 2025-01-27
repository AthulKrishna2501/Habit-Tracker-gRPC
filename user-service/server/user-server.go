package server

import (
	"context"
	"fmt"
	"user-service/user-service/db"
	pb "user-service/user-service/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userID, err := s.repo.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		return nil, fmt.Errorf("cannot create user %v", err)
	}

	return &pb.CreateUserResponse{UserId: userID}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetUser(ctx, req.UserId)

	if err != nil {
		return nil, fmt.Errorf("cannot get user %v", err)
	}

	return &pb.GetUserResponse{
		UserId: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}
