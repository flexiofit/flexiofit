// internal/services/user_service.go
package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository // Keep Gorm repository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, firstName, middleName, lastName, email, password string) (*models.User, error) {
	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &models.User{
		FirstName:    firstName,
		MiddleName:   middleName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	if err := validateUser(user); err != nil {
		return nil, err
	}

	// Use the Gorm repository to create the user
	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (*models.User, error) {
	return s.userRepository.FindByID(uint(id))
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, firstName, middleName, lastName, email, password string) (*models.User, error) {
	// Retrieve existing user
	user, err := s.userRepository.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	// Update user details
	user.FirstName = firstName
	user.MiddleName = middleName
	user.LastName = lastName
	user.Email = email

	// Update password if provided
	if password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %v", err)
		}
		user.PasswordHash = string(passwordHash)
	}

	if err := validateUser(user); err != nil {
		return nil, err
	}

	// Use the Gorm repository to update the user
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
	if user.FirstName == "" {
		return fmt.Errorf("first name cannot be empty")
	}
	if user.LastName == "" {
		return fmt.Errorf("last name cannot be empty")
	}
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if user.PasswordHash == "" {
		return fmt.Errorf("password cannot be empty")
	}
	return nil
}
