// internal/handlers/user_handler.go
package handlers

import (
	"backend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service *db.Queries // Using sqlc's generated Queries struct
}

// NewUserHandler initializes routes for user operations
// @Summary Initialize user routes
// @Description Sets up all the user-related routes in the application
// @Tags setup
func NewUserHandler(r *gin.Engine, s *db.Queries) {
	handler := &UserHandler{service: s}
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUserByID)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)
	r.GET("/users", handler.ListUsers)
}

// APIResponseUser maps the db.User struct for API responses
type APIResponseUser struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser handles POST /users
// @Summary Create a new user
// @Description Creates a new user with the provided username and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body struct{Username string "json:\"username\" binding:\"required\"" Email string "json:\"email\" binding:\"required,email\""} true "User creation request"
// @Success 201 {object} APIResponseUser "User created successfully"
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(c, db.CreateUserParams{
		Username: input.Username,
		Email:    input.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := APIResponseUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
	c.JSON(http.StatusCreated, response)
}

// GetUserByID handles GET /users/:id
// @Summary Get a user by ID
// @Description Retrieves a user's details by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} APIResponseUser "User found"
// @Failure 400 {object} gin.H "Invalid user ID"
// @Failure 404 {object} gin.H "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.service.GetUser(c, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := APIResponseUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
	c.JSON(http.StatusOK, response)
}

// UpdateUser handles PUT /users/:id
// @Summary Update a user
// @Description Updates an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body struct{Username string "json:\"username\" binding:\"required\"" Email string "json:\"email\" binding:\"required,email\""} true "User update request"
// @Success 200 {object} APIResponseUser "User updated successfully"
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateUser(c, db.UpdateUserParams{
		ID:       int32(id),
		Username: input.Username,
		Email:    input.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := APIResponseUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteUser handles DELETE /users/:id
// @Summary Delete a user
// @Description Deletes a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 400 {object} gin.H "Invalid user ID"
// @Failure 404 {object} gin.H "User not found"
// @Router /users/{id} [delete]
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

// ListUsers handles GET /users
// @Summary List all users
// @Description Retrieves a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} APIResponseUser "List of users"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []APIResponseUser
	for _, user := range users {
		response = append(response, APIResponseUser{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
		})
	}
	c.JSON(http.StatusOK, response)
}
