// internal/handlers/auth_handler.go
package handlers

import (
	"net/http"
	"backend/internal/middleware"
	"backend/internal/services"
	"backend/internal/resources/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService services.UserService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(userService services.UserService) Handler {
	return &AuthHandler{
		userService: userService,
	}
}


func (h *AuthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	// Public routes
	rg.Group("/auth")
	rg.POST("/login", h.Login)
	rg.POST("/refresh-token", h.RefreshToken)

	// Protected routes (after login)
	authenticated := rg.Group("/")
	authenticated.Use(middleware.AuthMiddleware())
	{
		authenticated.GET("/getUserInfo", h.GetUserInfo)
		authenticated.POST("/logout", h.Logout)
	}
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	// Replace with actual authentication logic
	if req.Username != "Soybean" || req.Password != "123456" {
		response.BadRequestError(c, "invalid credentials")
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := middleware.GenerateTokens(req.Username)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"msg":  "success",
		"data": gin.H{
			"token":        accessToken,
			"refreshToken": refreshToken,
		},
	})
}

// GetUserInfo returns user information for authenticated users
func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	// Extract username from context (set by AuthMiddleware)
	username, exists := c.Get("user")
	if !exists {
		response.BadRequestError(c, "User not found")
		// response.UnauthorizedError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"msg":  "success",
		"data": gin.H{
			"userId":   "0",
			"userName": username,
			"roles":    []string{"R_SUPER"},
			"buttons": []string{
				"B_CODE1",
				"B_CODE2",
				"B_CODE3",
			},
		},
	})
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// Get the refresh token from the request
	var refreshRequest struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	// Validate the refresh token
	claims, err := middleware.ValidateRefreshToken(refreshRequest.RefreshToken)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	// Extract username from claims
	username := claims.Data[0]["userName"]

	// Generate new access and refresh tokens
	accessToken, refreshToken, err := middleware.GenerateTokens(username)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"msg":  "success",
		"data": gin.H{
			"token":        accessToken,
			"refreshToken": refreshToken,
		},
	})
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *gin.Context) {
	// In a stateless JWT system, logout is typically handled client-side
	// You might want to implement token blacklisting or other logout mechanisms
	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"msg":  "logout successful",
	})
}
