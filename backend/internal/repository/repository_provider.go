package repository

import "gorm.io/gorm"

// ProvideUserRepository creates an instance of UserRepositoryInterface (dependency abstraction).
func ProvideUserRepository(db *gorm.DB) UserRepositoryInterface {
	return NewUserRepository(db)
}
