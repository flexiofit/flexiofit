// internal/services/services.go
package services

import (
	"backend/internal/repository"
	"gorm.io/gorm"
)

type Services struct {
	UserService *UserService
}

func NewServices(gormDB *gorm.DB) *Services {
	userRepository := repository.NewUserRepository(gormDB)

	return &Services{
		UserService: NewUserService(userRepository),
	}
}
