// internal/services/user_service.go
package services

import (
	db "backend/internal/db/sqlc"
	"context"
)

type UserService struct {
	db db.Querier
}

func NewUserService(querier db.Querier) *UserService {
	return &UserService{db: querier}
}

func (s *UserService) CreateUser(ctx context.Context, username, email string) (db.User, error) {
	params := db.CreateUserParams{
		Username: username,
		Email:    email,
	}
	return s.db.CreateUser(ctx, params)
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return s.db.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, username, email string) (db.User, error) {
	params := db.UpdateUserParams{
		ID:       id, // Make sure this field exists in your UpdateUserParams
		Username: username,
		Email:    email,
	}
	return s.db.UpdateUser(ctx, params)
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.db.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]db.User, error) {
	return s.db.ListUsers(ctx)
}
