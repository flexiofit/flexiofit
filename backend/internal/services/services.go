// internal/services/services.go
package services

import (
	db "backend/internal/db/sqlc"
)

type Services struct {
	UserService *UserService
	// AuthService *AuthService
	// Add other services as needed
}

func NewServices(querier db.Querier) *Services {
	return &Services{
		UserService: NewUserService(querier), // Pass the querier here
		// AuthService: NewAuthService(),
	}
}
