// internal/mappers/user_mapper.go
package mappers

import (
	"backend/internal/dtos"
	"backend/internal/models"
)

// ToUserDTO - Converts a domain user model to a user DTO.
func ToUserDTO(user *models.User) dtos.UserDTO {
	// Convert user.ID from uint to int32
	return dtos.UserDTO{
		ID:        int32(user.ID), // Type conversion here
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

// ToUserDTOs - Converts a slice of domain user models to user DTOs.
func ToUserDTOs(users []models.User) []dtos.UserDTO {
	var userDTOs []dtos.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, ToUserDTO(&user))
	}
	return userDTOs
}
