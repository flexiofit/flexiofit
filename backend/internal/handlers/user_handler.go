// internal/handlers/user_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"backend/internal/dtos"
	"backend/internal/mappers"
	"backend/internal/services"
	"github.com/gin-gonic/gin"
	. "backend/internal/resources/response"
	. "backend/internal/resources/constants"
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

func (h *UserHandler) CreateUser(c *gin.Context) {
    var input dtos.CreateUserRequest
    
    if err := c.ShouldBindJSON(&input); err != nil {
        BadRequestError(c, err.Error())
        return
    }

    user, err := h.service.CreateUser(c, input)
    if err != nil {
        InternalServerError(c, err)
        return
    }

    SendSuccessResponse(c, USER_CREATED, mappers.ToUserDTO(user))
}
// GetUserByID handles retrieving a user by ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		BadRequestError(c, INVALID_USER_INPUT)
		return
	}

	// Fetch user from service
	user, err := h.service.GetUserByID(c, int32(id))
	if err != nil {
		SendErrorResponse(c, STATUS_BAD_REQUEST, USER_NOT_FOUND, err.Error())
		return
	}

	// Return the user as a DTO
  SendSuccessResponse(c, SUCCESS, mappers.ToUserDTO(user))
}

// UpdateUser handles updating a user by ID.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendErrorResponse(c, STATUS_BAD_REQUEST, USER_NOT_FOUND, err.Error())
		return
	}

	var input dtos.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		SendErrorResponse(c, STATUS_BAD_REQUEST, "", err.Error())
		return
	}

	// Call service to update the user
	user, err := h.service.UpdateUser(c, int32(id), input.FirstName, "", input.LastName, input.Email, input.Password)
	if err != nil {
		SendErrorResponse(c, STATUS_BAD_REQUEST, INVALID_USER_INPUT, err.Error())
		return
	}
  SendSuccessResponse(c, USER_CREATED, mappers.ToUserDTO(user))
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendErrorResponse(c, STATUS_BAD_REQUEST, USER_NOT_FOUND, err.Error())
		return
	}

	// Call service to delete the user
	err = h.service.DeleteUser(c, int32(id))
	if err != nil {
		SendErrorResponse(c, STATUS_BAD_REQUEST, USER_NOT_FOUND, err.Error())
		return
	}

  SendSuccessResponse(c, USER_DELETED_SUCCESSFULLY, nil)
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
  SendSuccessResponse(c, SUCCESS, mappers.ToUserDTOs(users))
}
