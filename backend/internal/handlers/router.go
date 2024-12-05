// internal/handlers/router.go
package handlers

import (
	"backend/internal/services"
	"backend/internal/middleware"

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

	 // Apply the CORS middleware
	 router.Use(middleware.CORSMiddleware())

	api := router.Group("/api/v1")

	// Swagger route

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// Add AuthMiddleware to the protected group (e.g., user-specific routes)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware()) // Apply AuthMiddleware only to this group


	// List of all handlers
	handlers := []Handler{
		NewAuthHandler(*services.UserService),
		NewUserHandler(services.UserService),

		// Add new handlers here (e.g., NewAuthHandler, NewProductHandler, etc.)
	}

	// Register all routes dynamically
	for _, handler := range handlers {
		handler.RegisterRoutes(api)
	}


	// Add non-protected routes without authentication
	// Example: public routes that don't need auth (like login or registration)
	// public := api.Group("/public")
	// public.GET("/login", LoginHandler)
	//
	// Optional: Static file serving (e.g., frontend or assets)
	router.Static("/static", "./public")

	return router
}
