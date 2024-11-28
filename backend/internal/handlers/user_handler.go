// internal/handlers/user_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"backend/internal/dtos"
	"backend/internal/mappers"
	"backend/internal/services"
	"github.com/gin-gonic/gin"
	r "backend/internal/resources"

)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{service: userService}
}

// RegisterRoutes sets up routes for user-related operations.
func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("/:id", h.GetUserByID)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("", h.GetUsers)
	}
}

// CreateUser handles the creation of a user.
// func (h *UserHandler) CreateUser(c *gin.Context) {
// 	fmt.Println("12 12 12")
// 	var input dtos.CreateUserRequest
//
// 	fmt.Println("input", input)
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	// Call service to create a new user
// 	user, err := h.service.CreateUser(c,input)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	// Return created user as a DTO
// 	c.JSON(http.StatusCreated, mappers.ToUserDTO(user))
// }
func (h *UserHandler) CreateUser(c *gin.Context) {
    var input dtos.CreateUserRequest
    
    if err := c.ShouldBindJSON(&input); err != nil {
        r.BadRequestError(c, err.Error())
        return
    }

    user, err := h.service.CreateUser(c, input)
    if err != nil {
        r.InternalServerError(c, err)
        return
    }

    r.SendSuccessResponse(c, "User created successfully", mappers.ToUserDTO(user))
}
// GetUserByID handles retrieving a user by ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch user from service
	user, err := h.service.GetUserByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user as a DTO
	c.JSON(http.StatusOK, mappers.ToUserDTO(user))
}

// UpdateUser handles updating a user by ID.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input dtos.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to update the user
	user, err := h.service.UpdateUser(c, int32(id), input.FirstName, "", input.LastName, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return updated user as a DTO
	c.JSON(http.StatusOK, mappers.ToUserDTO(user))
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call service to delete the user
	err = h.service.DeleteUser(c, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetUsers handles retrieving all users.
func (h *UserHandler) GetUsers(c *gin.Context) {
	// Call service to get all users
	users, err := h.service.ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Return the list of users as DTOs
	c.JSON(http.StatusOK, mappers.ToUserDTOs(users))
}
