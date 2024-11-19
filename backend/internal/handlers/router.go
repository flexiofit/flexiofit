// internal/handlers/router.go
package handlers

import (
	"backend/internal/services"
	"github.com/gin-gonic/gin"
)

// Handler interface defines a contract for all handlers
type Handler interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

// SetupRouter sets up the router and registers all handlers
func SetupRouter(services *services.Services) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")

	// List of all handlers
	handlers := []Handler{
		NewUserHandler(services.UserService),
		// Add new handlers here (e.g., NewAuthHandler, NewProductHandler, etc.)
	}

	// Register all routes dynamically
	for _, handler := range handlers {
		handler.RegisterRoutes(api)
	}

	return router
}
