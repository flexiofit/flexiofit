// internal/services/services.go
package services

import (
	db "backend/internal/db/sqlc"
	"backend/internal/repository"
	"gorm.io/gorm"
)

type Services struct {
	UserService *UserService
}

func NewServices(sqlcQueries *db.Queries, gormDB *gorm.DB) *Services {
	userRepository := repository.NewUserRepository(gormDB)

	return &Services{
		UserService: NewUserService(sqlcQueries, userRepository),
	}
}
