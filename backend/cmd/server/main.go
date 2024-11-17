// cmd/server/main.go
package main

import (
	"backend/internal/config"
	db "backend/internal/db/sqlc"
	"backend/internal/handlers"
	"backend/internal/services"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

// runDatabaseMigrations handles the database migration process
func runDatabaseMigrations(db *sql.DB, migrationPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run migrations: %v", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("could not get migration version: %v", err)
	}

	log.Printf("Database migration completed. Version: %d, Dirty: %v", version, dirty)
	return nil
}

// setupDBPool configures and validates the database connection pool
func setupDBPool(db *sql.DB, config config.Config) error {
	db.SetMaxOpenConns(config.DBMaxOpenConns)
	db.SetMaxIdleConns(config.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(config.DBConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(config.DBConnMaxIdleTime) * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	dbConn, err := waitForDB(config, config.DBConnectRetries, time.Duration(config.DBConnectRetryDelay)*time.Second)
	if err != nil {
		log.Fatal("Failed to initialize database connection:", err)
	}
	defer dbConn.Close()

	// Run database migrations
	if err := runDatabaseMigrations(dbConn, "./internal/db/migrations"); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

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

	queries := db.New(dbConn)
	userService := services.NewUserService(queries)
	router := gin.Default()
	handlers.NewUserHandler(router, userService)

	serverAddr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
