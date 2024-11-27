package services

import (
	"backend/internal/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeServices(db *gorm.DB) *Services {
	wire.Build(
		repository.ProvideUserRepository, // Provide UserRepository
		ProvideUserService,               // Provide UserService
		wire.Struct(new(Services), "*"),  // Inject into Services struct
	)
	return &Services{}
}
