package services

import (
	"backend/internal/repository"
)

// ProvideUserService creates an instance of UserService.
func ProvideUserService(userRepo repository.UserRepositoryInterface) *UserService {
	if repo, ok := userRepo.(*repository.UserRepository); ok {
		return NewUserService(repo)
	}
	panic("userRepo is not of type *repository.UserRepository")
}
