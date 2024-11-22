// internal/repository/user_repository.go
// package repository
//
// import (
// 	"context"
// 	"errors"
//
// 	"gorm.io/gorm"
// )
//
// type User struct {
// 	gorm.Model
// 	Username string `gorm:"unique;not null"`
// 	Email    string `gorm:"unique;not null"`
// 	Password string `gorm:"not null"`
// 	// Add other fields as needed
// }
//
// type UserRepository struct {
// 	db *gorm.DB
// }
//
// func NewUserRepository(db *gorm.DB) *UserRepository {
// 	return &UserRepository{db: db}
// }
//
// func (r *UserRepository) Create(ctx context.Context, user *User) error {
// 	return r.db.WithContext(ctx).Create(user).Error
// }
//
// func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
// 	var user User
// 	result := r.db.WithContext(ctx).Where("username = ?", username).First(&user)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }
//
// func (r *UserRepository) Update(ctx context.Context, user *User) error {
// 	return r.db.WithContext(ctx).Save(user).Error
// }
//
// func (r *UserRepository) Delete(ctx context.Context, id uint) error {
// 	return r.db.WithContext(ctx).Delete(&User{}, id).Error
// }
//
// func (r *UserRepository) List(ctx context.Context, page, pageSize int) ([]User, int64, error) {
// 	var users []User
// 	var total int64
//
// 	// Count total records
// 	r.db.WithContext(ctx).Model(&User{}).Count(&total)
//
// 	// Paginate results
// 	result := r.db.WithContext(ctx).
// 		Offset((page - 1) * pageSize).
// 		Limit(pageSize).
// 		Find(&users)
//
// 	return users, total, result.Error
// }

package repository

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

//	func (r *UserRepository) ListUsers() ([]models.User, error) {
//		var users []models.User
//		err := r.db.Find(&users).Error
//		return users, err
//	}
// func (r *UserRepository) ListUsers() ([]models.User, error) {
// 	var users []models.User
// 	err := r.db.Unscoped().Where("deleted_at IS NULL").Find(&users).Error
// 	return users, err
// }

func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Unscoped().Find(&users).Error
	return users, err
}
