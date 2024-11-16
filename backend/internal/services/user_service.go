// internal/services/user_service.go
package services

import (
	"context"
	"errors"
	"backend/internal/db"
	"backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserInput struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Password  string `json:"password" binding:"required,min=6"`
	FullName  string `json:"full_name" binding:"required"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}

func (s *UserService) CreateUser(ctx context.Context, input CreateUserInput) (db.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, err
	}

	// Create user
	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Email:        input.Email,
		Username:     input.Username,
		FullName:     input.FullName,
		PasswordHash: string(hashedPassword),
		Bio:          sql.NullString{String: input.Bio, Valid: input.Bio != ""},
		AvatarURL:    sql.NullString{String: input.AvatarURL, Valid: input.AvatarURL != ""},
	})

	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (db.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error) {
	return s.repo.ListUsers(ctx, limit, offset)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, input db.UpdateUserParams) (db.User, error) {
	return s.repo.UpdateUser(ctx, input)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
