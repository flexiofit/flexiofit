// internal/services/services.go
package services

import (
	"backend/internal/repository"
	"gorm.io/gorm"
)

type Services struct {
	UserService *UserService
	// OtherService    *OtherService  // Add more services if needed
}

func NewServices(gormDB *gorm.DB) *Services {
	// Instantiate multiple repositories
	userRepository := repository.NewUserRepository(gormDB)
	// otherRepository := repository.NewOtherRepository(gormDB) // Another repository instance

	// Pass multiple repositories into the services
	return &Services{
		UserService: NewUserService(userRepository),
		// OtherService: NewOtherService(otherRepository),
	}
}
