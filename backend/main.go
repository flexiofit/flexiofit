package main

import (
    "context"
    "database/sql"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/flexiofit/flexiofit/backend/api"
    "github.com/flexiofit/flexiofit/backend/config"  // Ensure correct import
    "github.com/flexiofit/flexiofit/backend/logger"
    "github.com/flexiofit/flexiofit/backend/search/typesense"
    "github.com/joho/godotenv"
)

type serverConfig struct {
    log        logger.Logger
    config     config.Config  // Change to config.Config
    dbConn     *sql.DB
    store      db.Store
    tsHandler  *typesense.TypesenseHandler
}

func setupLogger(env string) logger.Logger {
    var logger *zap.Logger
    var err error

    switch env {
    case "test":
        atom := zap.NewAtomicLevel()
        atom.SetLevel(zapcore.ErrorLevel)
        zapConfig := zap.NewDevelopmentConfig()
        zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
        zapConfig.Level = atom
        logger, err = zapConfig.Build()
    case "dev", "":
        logger, err = zap.NewDevelopment()
    default:
        logger, err = zap.NewProduction()
    }

    if err != nil {
        panic(fmt.Sprintf("failed to initialize logger: %v", err))
    }
    
    return logger.Sugar() // Ensure you return the SugaredLogger
}

func buildDBUrl(config config.Config, env string) string {  // Change to config.Config
    dbName := config.DBDbName
    if env == "test" {
        dbName = config.DBDbNameTest
    }
    
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        config.DB_USERNAME,
        config.DB_PASSWORD,
        config.DB_HOSTNAME,
        config.DB_PORT,
        dbName,  // Use dbName from config
    )
}

func initializeServer(env string) *serverConfig {
    // Set default environment to "dev" if not provided
    if env == "" {
        env = "dev"
    }

    // Load environment variables from the correct file in the `env` folder
    envFile := fmt.Sprintf("./env/%s.env", env)
    if err := godotenv.Load(envFile); err != nil {
        fmt.Printf("Warning: Error loading %s file: %v\n", envFile, err)
    } else {
        fmt.Println("Environment variables loaded successfully from:", envFile)
    }

    // Initialize components
    log := setupLogger(env)
    config := config.LoadConfig(env, "./env")  // Ensure correct config package
    log.Infow("Loaded config", "environment", env, "config", config)

    // Set up database connection
    config.DBUrl = buildDBUrl(config, env)
    log.Infow("Database configuration", "url", config.DBUrl)
    
    fmt.Println("Connecting to the database...")
    dbConn := db.Connect(config)
    fmt.Println("Connected to the database successfully")
    if dbConn == nil {
        log.Fatal("Failed to connect to database")
    }

    log.Info("Successfully connected to database")

    if err := db.AutoMigrate(config); err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    log.Info("Database migrations completed successfully")

    // Initialize store and typesense
    store := db.NewConduitStore(dbConn)
    tsClient := typesense.NewClient(&config)
    tsHandler := typesense.NewTypesenseHandler(tsClient, "articles")
    
    if err := tsHandler.CreateCollection(); err != nil {
        log.Fatalf("Failed to create typesense collection: %v", err)
    }

    log.Info("Typesense collection created successfully")

    return &serverConfig{
        log:       log,
        config:    config,
        dbConn:    dbConn,
        store:     store,
        tsHandler: tsHandler,
    }
}

func setupServer(sc *serverConfig) *http.Server {
    server := api.NewServer(
        sc.config,
        sc.store,
        sc.tsHandler,
        sc.log,
    )

    server.MountHandlers()
    server.MountSwaggerHandlers()

    addr := fmt.Sprintf(":%s", sc.config.Port)
    sc.log.Infow("Server configuration", "address", addr)
    
    return &http.Server{
        Addr:    addr,
        Handler: server.Router(),
    }
}

func runServer(srv *http.Server, log logger.Logger) {
    // Launch server in a goroutine
    go func() {
        log.Infof("Starting server on %s", srv.Addr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error starting server: %v", err)
        }
    }()

    // Wait for interrupt signal to gracefully shutdown the server
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Info("Shutdown Server ...")

    // Graceful shutdown with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Increased timeout
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Info("Server exiting")
}

func main() {
    env := os.Getenv("ENVIRONMENT")
    sc := initializeServer(env)
    defer func() {
        if err := sc.dbConn.Close(); err != nil {
            sc.log.Errorf("Error closing database connection: %v", err)
        }
    }()
    
    srv := setupServer(sc)
    sc.log.Info("Server setup complete, starting...")
    runServer(srv, sc.log)
}
