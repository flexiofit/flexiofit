package dtos

type UserDTO struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserType  string `json:"user_type"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	MiddleName string `json:"middle_name"`
	LastName  string `json:"last_name" binding:"required"`
	UserType  string    `json:"user_type" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"min=6"`
}
