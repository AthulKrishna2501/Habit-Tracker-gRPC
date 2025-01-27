package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Habit struct {
	ID          string
	UserID      string
	Name        string
	Description string
	Streak      int32
}

type HabitRepository struct {
	conn *gorm.DB
}

func NewHabitRepository(dbURL string) (*HabitRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	return &HabitRepository{conn: conn}, nil
}

func (r *HabitRepository) CreateHabit(ctx context.Context, UserID, name, description string) (string, error) {
	id := uuid.NewString()

	newHabit := Habit{
		ID:          id,
		UserID:      UserID,
		Name:        name,
		Description: description,
	}

	err := r.conn.Create(&newHabit).Error
	if err != nil {
		return "", fmt.Errorf("cannot create habit %v", err)
	}

	log.Printf("Creating habit for userID: %s", UserID)
	log.Print("NewHabit:", newHabit)

	return id, nil

}

func (r *HabitRepository) GetHabits(ctx context.Context, UserID string) ([]Habit, error) {
	var habit []Habit

	err := r.conn.Where("user_id = ?", UserID).Find(&habit).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get habits %v", err)
	}

	log.Printf("Retriving habits for user; %s", UserID)
	log.Print("Habits: ", habit)
	return habit, nil

}

func (r *HabitRepository) UpdateHabit(ctx context.Context, id, name, description string, streak int32) error {
	var habit Habit

	updateHabit := Habit{
		Name:        name,
		Description: description,
		Streak:      streak,
	}
	err := r.conn.Model(&habit).Where("id=?", id).Updates(&updateHabit)
	if err != nil {
		return fmt.Errorf("failed to update habit %v", err)
	}

	return nil
}

func (r *HabitRepository) DeleteHabit(ctx context.Context, id string) error {
	var habit Habit

	err := r.conn.Where("id=?", id).Delete(&habit).Error
	if err != nil {
		return fmt.Errorf("failed to delete habit %v", err)
	}

	return nil
}
