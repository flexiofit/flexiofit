// cmd/server/main.go
package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	db "backend/internal/db/sqlc"
	"backend/internal/handlers"
	"backend/internal/services"
)

func main() {
	// Initialize database connection pool
	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres:postgres@localhost:5432/mydb")
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// Initialize database queries
	queries := db.New(dbpool)

	// Initialize service layer
	userService := services.NewUserService(queries)

	// Initialize Gin router
	router := gin.Default()

	// Initialize handlers
	handlers.NewUserHandler(router, userService)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
