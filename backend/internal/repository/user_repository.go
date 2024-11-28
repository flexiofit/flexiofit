// internal/repository/user_repository.go
package repository

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

// UserRepositoryInterface defines the contract for user-related database operations
type UserRepositoryInterface interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	ListUsers() ([]models.User, error)
}

// UserRepository implements UserRepositoryInterface
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByID retrieves a user by their ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates the details of an existing user in the database
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete removes a user by their ID from the database
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// ListUsers retrieves all users from the database
func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Unscoped().Find(&users).Error
	return users, err
}
