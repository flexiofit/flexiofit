// internal/services/user_service.go
// package services
//
// import (
// 	db "backend/internal/db/sqlc"
// 	"context"
// )
//
// type UserService struct {
// 	db db.Querier
// }
//
// func NewUserService(querier db.Querier) *UserService {
// 	return &UserService{db: querier}
// }
//
// func (s *UserService) CreateUser(ctx context.Context, username, email string) (db.User, error) {
// 	params := db.CreateUserParams{
// 		Username: username,
// 		Email:    email,
// 	}
// 	return s.db.CreateUser(ctx, params)
// }
//
// func (s *UserService) GetUserByID(ctx context.Context, id int32) (db.User, error) {
// 	return s.db.GetUserByID(ctx, id)
// }
//
// func (s *UserService) UpdateUser(ctx context.Context, id int32, username, email string) (db.User, error) {
// 	params := db.UpdateUserParams{
// 		ID:       id, // Make sure this field exists in your UpdateUserParams
// 		Username: username,
// 		Email:    email,
// 	}
// 	return s.db.UpdateUser(ctx, params)
// }
//
// func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
// 	return s.db.DeleteUser(ctx, id)
// }
//
// func (s *UserService) ListUsers(ctx context.Context) ([]db.User, error) {
// 	return s.db.ListUsers(ctx)
// }

package services

import (
	"context"
	"fmt"

	db "backend/internal/db/sqlc"
	"backend/internal/models"
	"backend/internal/repository"
)

type UserService struct {
	queries        *db.Queries
	userRepository *repository.UserRepository
}

func NewUserService(queries *db.Queries, userRepository *repository.UserRepository) *UserService {
	return &UserService{
		queries:        queries,
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username, email string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Email:    email,
	}

	if err := validateUser(user); err != nil {
		return nil, err
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (*models.User, error) {
	return s.userRepository.FindByID(uint(id))
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, username, email string) (*models.User, error) {
	user, err := s.userRepository.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email

	if err := validateUser(user); err != nil {
		return nil, err
	}

	if err := s.userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.userRepository.Delete(uint(id))
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepository.ListUsers()
}

func validateUser(user *models.User) error {
	if user.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	return nil
}

// package services
//
// import (
// 	"context"
// 	"fmt"
//
// 	db "backend/internal/db/sqlc"
// 	"backend/internal/models"
// 	"backend/internal/repository"
// )
//
// type UserService struct {
// 	queries        *db.Queries
// 	userRepository *repository.UserRepository
// }
//
// func NewUserService(queries *db.Queries, userRepository *repository.UserRepository) *UserService {
// 	return &UserService{
// 		queries:        queries,
// 		userRepository: userRepository,
// 	}
// }
//
// func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
// 	if err := validateUser(user); err != nil {
// 		return err
// 	}
//
// 	return s.userRepository.Create(user)
// }
//
// func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
// 	// Option 1: Use Gorm repository
// 	return s.userRepository.FindByID(id)
//
// 	// Option 2: If you want to use SQLC instead
// 	// dbUser, err := s.queries.GetUserByID(ctx, int32(id))
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("failed to get user: %w", err)
// 	// }
// 	// return &models.User{
// 	// 	Username: dbUser.Username,
// 	// 	Email:    dbUser.Email,
// 	// }, nil
// }
//
// func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
// 	return s.userRepository.FindByUsername(username)
// }
//
// func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
// 	if err := validateUser(user); err != nil {
// 		return err
// 	}
//
// 	return s.userRepository.Update(user)
// }
//
// func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
// 	return s.userRepository.Delete(id)
// }
//
// func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
// 	return s.userRepository.ListUsers()
// }
//
// func validateUser(user *models.User) error {
// 	if user.Username == "" {
// 		return fmt.Errorf("username cannot be empty")
// 	}
// 	if user.Email == "" {
// 		return fmt.Errorf("email cannot be empty")
// 	}
// 	return nil
// }
