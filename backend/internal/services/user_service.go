// internal/services/user_service.go
package services

import (
	"backend/internal/models"
	"backend/internal/dtos"
	"backend/internal/repository"
	. "backend/internal/resources/constants"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	// Change to use the interface instead of concrete repository type
	userRepository repository.UserRepositoryInterface
}

// Update constructor to accept the interface
func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, input dtos.CreateUserRequest) (*models.User, error) {
	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

userType := GetUserType(input.UserType)
if userType == INVALID {
    return nil, fmt.Errorf("invalid user type: %s", input.UserType)
}

	user := &models.User{
		FirstName:    input.FirstName,
		MiddleName:   input.MiddleName,
		LastName:     input.LastName,
		Email:        input.Email,
		PasswordHash: string(passwordHash),
		UserType:     userType,
	}

	if err := validateUser(user); err != nil {
		return nil, err
	}

	// Use the repository interface method
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

	// Use the repository interface method
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

// ListUsersWithFilters demonstrates using the enhanced repository features
func (s *UserService) ListUsersWithFilters(ctx context.Context, page, limit int, filters map[string]interface{}) ([]models.User, error) {
	// Add pagination to filters
	filters["offset"] = page
	filters["limit"] = limit

	// Using the repository's new method
	if repo, ok := s.userRepository.(*repository.UserRepository); ok {
		return repo.ListUsersWithFilters(filters)
	}
	
	// Fallback to regular list if repository doesn't support filters
	return s.userRepository.ListUsers()
}
