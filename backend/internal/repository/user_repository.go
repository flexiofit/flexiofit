// internal/repository/user_repository.go
// package repository
//
// import (
// 	"backend/internal/models"
// 	"gorm.io/gorm"
// )
//
// // UserRepositoryInterface defines the contract for user-related database operations
// type UserRepositoryInterface interface {
// 	Create(user *models.User) error
// 	FindByID(id uint) (*models.User, error)
// 	FindByEmail(email string) (*models.User, error)
// 	Update(user *models.User) error
// 	Delete(id uint) error
// 	ListUsers() ([]models.User, error)
// }
//
// // UserRepository implements UserRepositoryInterface
// type UserRepository struct {
// 	db *gorm.DB
// }
//
// // NewUserRepository creates a new UserRepository instance
// func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
// 	return &UserRepository{db: db}
// }
//
// // Create inserts a new user into the database
// func (r *UserRepository) Create(user *models.User) error {
// 	return r.db.Create(user).Error
// }
//
// // FindByID retrieves a user by their ID
// func (r *UserRepository) FindByID(id uint) (*models.User, error) {
// 	var user models.User
// 	err := r.db.First(&user, id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
//
// // FindByEmail retrieves a user by their email
// func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
//
// // Update updates the details of an existing user in the database
// func (r *UserRepository) Update(user *models.User) error {
// 	return r.db.Save(user).Error
// }
//
// // Delete removes a user by their ID from the database
// func (r *UserRepository) Delete(id uint) error {
// 	return r.db.Delete(&models.User{}, id).Error
// }
//
// // ListUsers retrieves all users from the database
// func (r *UserRepository) ListUsers() ([]models.User, error) {
// 	var users []models.User
// 	// err := r.db.Unscoped().Find(&users).Error
// 	err := r.db.Find(&users).Error
//
// 	return users, err
// }
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
	// Add base repository methods you want to expose
	GetPagination(filter map[string]interface{}) Pagination
	BuildQuery(query *gorm.DB, conditions map[string]interface{}) *gorm.DB
}

// UserRepository implements UserRepositoryInterface
type UserRepository struct {
	*BaseRepository
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *models.User) error {
	return r.Save(user)
}

// FindByID retrieves a user by their ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB().First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	conditions := map[string]interface{}{
		"email": map[string]interface{}{
			"Op":    "eq",
			"value": email,
		},
	}
	
	query := r.DB().Model(&models.User{})
	query = r.BuildQuery(query, conditions)
	
	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates the details of an existing user in the database
func (r *UserRepository) Update(user *models.User) error {
	return r.Save(user)
}

// Delete removes a user by their ID from the database
func (r *UserRepository) Delete(id uint) error {
    _, err := r.DeleteByField(&models.User{}, []interface{}{id}, "id")
    return err
}

// ListUsers retrieves all users from the database with pagination
func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB().Find(&users).Error
	return users, err
}

// Example of a new method using base repository functionality
func (r *UserRepository) ListUsersWithFilters(filters map[string]interface{}) ([]models.User, error) {
	var users []models.User
	
	// Get pagination settings
	pagination := r.GetPagination(filters)
	
	// Build query with conditions
	query := r.DB().Model(&models.User{})
	query = r.BuildQuery(query, filters)
	
	// Apply pagination
	query = query.Limit(pagination.Limit).Offset(pagination.Offset)
	
	err := query.Find(&users).Error
	return users, err
}
