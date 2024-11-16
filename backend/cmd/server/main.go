// cmd/server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	// "os"
	"time"

	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"

	"backend/internal/config"
	db "backend/internal/db/sqlc"
	"backend/internal/handlers"
	"backend/internal/services"
)

// setupDBPool configures and validates the database connection pool
func setupDBPool(db *sql.DB, config config.Config) error {
	// Set maximum number of open connections
	db.SetMaxOpenConns(config.DBMaxOpenConns)

	// Set maximum number of idle connections
	db.SetMaxIdleConns(config.DBMaxIdleConns)

	// Set maximum lifetime of a connection
	db.SetConnMaxLifetime(time.Duration(config.DBConnMaxLifetime) * time.Minute)

	// Set maximum idle time for connections
	db.SetConnMaxIdleTime(time.Duration(config.DBConnMaxIdleTime) * time.Minute)

	// Create a context with timeout for connection validation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify connection pool
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to verify database connection: %v", err)
	}

	return nil
}

// waitForDB attempts to connect to the database with retries
func waitForDB(config config.Config, maxAttempts int, retryDelay time.Duration) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)

	// Create basic connection string without pool parameters
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = sql.Open("pgx", dsn)
		if err != nil {
			log.Printf("Failed to open database (attempt %d/%d): %v", attempt, maxAttempts, err)
			time.Sleep(retryDelay)
			continue
		}

		if err = setupDBPool(db, config); err != nil {
			log.Printf("Failed to setup database pool (attempt %d/%d): %v", attempt, maxAttempts, err)
			db.Close()
			time.Sleep(retryDelay)
			continue
		}

		log.Printf("Successfully connected to database (attempt %d/%d)", attempt, maxAttempts)
		return db, nil
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxAttempts)
}

func main() {
	// Load configuration
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Attempt to connect to the database with retries
	dbConn, err := waitForDB(config, config.DBConnectRetries, time.Duration(config.DBConnectRetryDelay)*time.Second)
	if err != nil {
		log.Fatal("Failed to initialize database connection:", err)
	}
	defer dbConn.Close()

	// Start a background health check if enabled
	if config.DBHealthCheckPeriod > 0 {
		go func() {
			ticker := time.NewTicker(time.Duration(config.DBHealthCheckPeriod) * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
					if err := dbConn.PingContext(ctx); err != nil {
						log.Printf("Database health check failed: %v", err)
					}
					cancel()
				}
			}
		}()
	}

	// Initialize database queries
	queries := db.New(dbConn)

	// Initialize service layer
	userService := services.NewUserService(queries)

	// Initialize Gin router
	router := gin.Default()

	// Initialize handlers
	handlers.NewUserHandler(router, userService)

	// Start server with configured host and port
	serverAddr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
