package repository

import "backend/internal/models"

// UserRepositoryInterface defines the contract for UserRepository.
type UserRepositoryInterface interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	ListUsers() ([]models.User, error)
}
