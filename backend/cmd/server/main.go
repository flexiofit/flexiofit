// cmd/server/main.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"backend/internal/config"
	db "backend/internal/db/sqlc"
	"backend/internal/handlers"
	"backend/internal/logging"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	postgresGorm "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupGormDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBSSLMode,
	)

	db, err := gorm.Open(postgresGorm.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL database: %v", err)
	}

	sqlDB.SetMaxOpenConns(config.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(config.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.DBConnMaxLifetime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.DBConnMaxIdleTime) * time.Minute)

	return db, nil
}

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

	logging.Log.Info("Database migration completed",
		zap.Uint("version", version),
		zap.Bool("dirty", dirty),
	)
	return nil
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		logging.Log.Fatal("Cannot load config", zap.Error(err))
	}

	if err := logging.InitializeLogger(config.ToLoggerConfig()); err != nil {
		logging.Log.Fatal("Failed to initialize logger", zap.Error(err))
	}
	defer logging.Sync()

	// Setup GORM database connection
	gormDB, err := setupGormDB(config)
	if err != nil {
		logging.Log.Fatal("Failed to initialize GORM database connection", zap.Error(err))
	}

	// Get SQL database for SQLC
	sqlDB, err := gormDB.DB()
	if err != nil {
		logging.Log.Fatal("Failed to get SQL database", zap.Error(err))
	}

	// Run database migrations
	if config.EnableMigration {
		if err := runDatabaseMigrations(sqlDB, "./internal/db/migrations"); err != nil {
			logging.Log.Fatal("Failed to run database migrations", zap.Error(err))
		}
	} else {
		logging.Log.Info("Database migrations are disabled")
	}

	// Run GORM automatic migrations for registered models
	if err := models.AutoMigrateDB(gormDB); err != nil {
		logging.Log.Fatal("Failed to run GORM auto migration",
			zap.Error(err),
			zap.Strings("registered_models", models.GetRegisteredModels(gormDB)),
		)
	}

	logging.Log.Info("GORM auto migration completed successfully",
		zap.Strings("migrated_models", models.GetRegisteredModels(gormDB)),
	)

	// Database health check
	if config.DBHealthCheckPeriod > 0 {
		go func() {
			ticker := time.NewTicker(time.Duration(config.DBHealthCheckPeriod) * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
				if err := sqlDB.PingContext(ctx); err != nil {
					logging.Log.Error("Database health check failed", zap.Error(err))
				}
				cancel()
			}
		}()
	}

	// Setup SQLC queries
	queries := db.New(sqlDB)

	// Initialize services with both SQLC and GORM
	// allServices := services.NewServices(queries)
	allServices := services.NewServices(queries, gormDB)

	// Setup router
	router := handlers.SetupRouter(allServices)

	// Start server
	serverAddr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	logging.Log.Info("Starting server",
		zap.String("host", config.ServerHost),
		zap.String("port", config.ServerPort),
	)

	if err := router.Run(serverAddr); err != nil {
		logging.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
