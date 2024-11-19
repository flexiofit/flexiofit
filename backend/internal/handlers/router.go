// internal/handlers/router.go
package handlers

import (
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler interface defines a contract for all handlers
type Handler interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

// SetupRouter sets up the router and registers all handlers
func SetupRouter(services *services.Services) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")

	// Swagger route

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
