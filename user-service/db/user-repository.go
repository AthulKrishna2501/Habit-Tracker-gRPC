package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(dbURL string) (*UserRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	return &UserRepository{conn: conn}, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, name, email string) (string, error) {

	id := uuid.NewString()

	if r.conn == nil {
		return "", fmt.Errorf("r.conn is missing")
	}

	newUser := User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	fmt.Printf("Creating user with ID: %s, Name: %s, Email: %s\n", newUser.ID, newUser.Name, newUser.Email)
	result := r.conn.Create(&newUser)

	if result.Error != nil {
		fmt.Printf("Error details: %v\n", result.Error)
		return "", fmt.Errorf("failed to create user: %v", result.Error)
	}
	return id, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*User, error) {
	var user User
	if id == "" {
		return nil, fmt.Errorf("id is null")
	}
	if err := r.conn.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &user, nil
}
