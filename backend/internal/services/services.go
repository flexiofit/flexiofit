// internal/services/services.go
// package services
//
// import (
//
//	db "backend/internal/db/sqlc"
//
// )
//
//	type Services struct {
//		UserService *UserService
//		// AuthService *AuthService
//		// Add other services as needed
//	}
//
//	func NewServices(querier db.Querier) *Services {
//		return &Services{
//			UserService: NewUserService(querier), // Pass the querier here
//			// AuthService: NewAuthService(),
//		}
//	}
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
