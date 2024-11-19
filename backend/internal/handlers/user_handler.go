// internal/handlers/user_handler.go
// package handlers
//
// import (
//
//	"net/http"
//	"strconv"
//
//	"backend/internal/services"
//	"github.com/gin-gonic/gin"
//
// )
//
//	type UserHandler struct {
//		service *services.UserService
//	}
//
//	func NewUserHandler(userService *services.UserService) *UserHandler {
//		return &UserHandler{service: userService}
//	}
//
// // Register updated routes
//
//	func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
//		userGroup := rg.Group("/users")
//		{
//			userGroup.POST("", h.CreateUser)
//			userGroup.GET("/:id", h.GetUserByID)
//			userGroup.PUT("/:id", h.UpdateUser)
//			userGroup.DELETE("/:id", h.DeleteUser)
//			userGroup.GET("", h.ListUsers)
//		}
//	}
//
// // CreateUser godoc
// // @Summary Create a new user
// // @Description Create a new user with username and email
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param user body CreateUserRequest true "User Create Request"
// // @Success 201 {object} User
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /users [post]
//
//	func (h *UserHandler) CreateUser(c *gin.Context) {
//		var input struct {
//			Username string `json:"username" binding:"required"`
//			Email    string `json:"email" binding:"required,email"`
//		}
//		if err := c.ShouldBindJSON(&input); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		user, err := h.service.CreateUser(c, input.Username, input.Email)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusCreated, user)
//	}
//
// // GetUserByID godoc
// // @Summary Get a user by ID
// // @Description Retrieve a user's details by their ID
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param id path int true "User ID"
// // @Success 200 {object} User
// // @Failure 400 {object} map[string]string
// // @Failure 404 {object} map[string]string
// // @Router /users/{id} [get]
//
//	func (h *UserHandler) GetUserByID(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//			return
//		}
//		user, err := h.service.GetUserByID(c, int32(id))
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//			return
//		}
//		c.JSON(http.StatusOK, user)
//	}
//
// // UpdateUser godoc
// // @Summary Update an existing user
// // @Description Update user details by user ID
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param id path int true "User ID"
// // @Param user body UpdateUserRequest true "User Update Request"
// // @Success 200 {object} User
// // @Failure 400 {object} map[string]string "Invalid user ID or request"
// // @Failure 500 {object} map[string]string "Internal server error"
// // @Router /users/{id} [put]
//
//	func (h *UserHandler) UpdateUser(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//			return
//		}
//		var input struct {
//			Username string `json:"username" binding:"required"`
//			Email    string `json:"email" binding:"required,email"`
//		}
//		if err := c.ShouldBindJSON(&input); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		user, err := h.service.UpdateUser(c, int32(id), input.Username, input.Email)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, user)
//	}
//
// // DeleteUser godoc
// // @Summary Delete a user
// // @Description Delete user by user ID
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param id path int true "User ID"
// // @Success 204 "User successfully deleted"
// // @Failure 400 {object} map[string]string "Invalid user ID"
// // @Failure 404 {object} map[string]string "User not found"
// // @Router /users/{id} [delete]
//
//	func (h *UserHandler) DeleteUser(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//			return
//		}
//		if err := h.service.DeleteUser(c, int32(id)); err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//			return
//		}
//		c.Status(http.StatusNoContent)
//	}
//
// // ListUsers godoc
// // @Summary List all users
// // @Description Retrieve a list of all users
// // @Tags users
// // @Accept json
// // @Produce json
// // @Success 200 {array} ListUsersResponse
// // @Failure 500 {object} map[string]string "Internal server error"
// // @Router /users [get]
//
//	func (h *UserHandler) ListUsers(c *gin.Context) {
//		users, err := h.service.ListUsers(c)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, users)
//	}
package handlers

import (
	"net/http"
	"strconv"

	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

// User represents the user model for Swagger documentation
type User struct {
	ID       int32  `json:"id" example:"1"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@example.com"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe_updated"`
	Email    string `json:"email" binding:"required,email" example:"john_updated@example.com"`
}

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	service *services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{service: userService}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with username and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User Create Request"
// @Success 201 {object} User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.CreateUser(c, input.Username, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieve a user's details by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUserByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UpdateUserRequest true "User Update Request"
// @Success 200 {object} User
// @Failure 400 {object} map[string]string "Invalid user ID or request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var input UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.UpdateUser(c, int32(id), input.Username, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete user by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User successfully deleted"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 404 {object} map[string]string "User not found"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := h.service.DeleteUser(c, int32(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// ListUsers godoc
// @Summary List all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// RegisterRoutes sets up user-related routes
func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	{
		userGroup.POST("", h.CreateUser)
		userGroup.GET("/:id", h.GetUserByID)
		userGroup.PUT("/:id", h.UpdateUser)
		userGroup.DELETE("/:id", h.DeleteUser)
		userGroup.GET("", h.ListUsers)
	}
}
