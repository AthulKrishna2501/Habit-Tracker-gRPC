package server

import (
	"context"
	"fmt"
	"user-service/habit-service/db"
	pb "user-service/habit-service/proto"
)

type HabitServer struct {
	pb.UnimplementedHabitServiceServer
	repo *db.HabitRepository
}

func NewHabitServer(repo *db.HabitRepository) *HabitServer {
	return &HabitServer{repo: repo}
}

func (s *HabitServer) CreateHabit(ctx context.Context, req *pb.CreateHabitRequest) (*pb.CreateHabitResponse, error) {
	HabitID, err := s.repo.CreateHabit(ctx, req.UserId, req.Name, req.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create habit %v", err)
	}

	return &pb.CreateHabitResponse{HabitId: HabitID}, nil
}
func (s *HabitServer) GetHabits(ctx context.Context, req *pb.GetHabitsRequest) (*pb.GetHabitsResponse, error) {
	habits, err := s.repo.GetHabits(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get habits %v", err)
	}

	var pbHabits []*pb.Habit

	for _, habit := range habits {
		pbHabits = append(pbHabits, &pb.Habit{
			HabitId:     habit.ID,
			Name:        habit.Name,
			UserId:      habit.ID,
			Description: habit.Description,
		})
	}

	return &pb.GetHabitsResponse{Habits: pbHabits}, nil
}
