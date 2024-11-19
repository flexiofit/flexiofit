// internal/config/config.go
package config

import (
	"backend/internal/logging"
	"github.com/spf13/viper"
)

type Config struct {
	// Database settings
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	DBSSLMode           string `mapstructure:"DB_SSLMODE"`
	DBMaxOpenConns      int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns      int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBConnMaxLifetime   int    `mapstructure:"DB_CONN_MAX_LIFETIME"`
	DBConnMaxIdleTime   int    `mapstructure:"DB_CONN_MAX_IDLE_TIME"`
	DBHealthCheckPeriod int    `mapstructure:"DB_HEALTH_CHECK_PERIOD"`
	DBConnectRetries    int    `mapstructure:"DB_CONNECT_RETRIES"`
	DBConnectRetryDelay int    `mapstructure:"DB_CONNECT_RETRY_DELAY"`

	// Server settings
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerHost string `mapstructure:"SERVER_HOST"`

	// JWT settings
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	JWTExpirationHours int    `mapstructure:"JWT_EXPIRATION_HOURS"`

	// Logger settings
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	LogFilePath    string `mapstructure:"LOG_FILE_PATH"`
	LogMaxSize     int    `mapstructure:"LOG_MAX_SIZE"`
	LogMaxBackups  int    `mapstructure:"LOG_MAX_BACKUPS"`
	LogMaxAge      int    `mapstructure:"LOG_MAX_AGE"`
	LogCompress    bool   `mapstructure:"LOG_COMPRESS"`
	LogDevelopment bool   `mapstructure:"LOG_DEVELOPMENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set default values for database
	viper.SetDefault("DB_MAX_OPEN_CONNS", 25)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 10)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", 15)
	viper.SetDefault("DB_CONN_MAX_IDLE_TIME", 5)
	viper.SetDefault("DB_HEALTH_CHECK_PERIOD", 30)
	viper.SetDefault("DB_CONNECT_RETRIES", 5)
	viper.SetDefault("DB_CONNECT_RETRY_DELAY", 5)

	// Set default values for logger
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_FILE_PATH", "./logs/app.log")
	viper.SetDefault("LOG_MAX_SIZE", 100)
	viper.SetDefault("LOG_MAX_BACKUPS", 3)
	viper.SetDefault("LOG_MAX_AGE", 28)
	viper.SetDefault("LOG_COMPRESS", true)
	viper.SetDefault("LOG_DEVELOPMENT", false)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// Convert config to LoggerConfig for zap logger
func (c *Config) ToLoggerConfig() logging.LoggerConfig {
	return logging.LoggerConfig{
		Level:       c.LogLevel,
		Filepath:    c.LogFilePath,
		MaxSize:     c.LogMaxSize,
		MaxBackups:  c.LogMaxBackups,
		MaxAge:      c.LogMaxAge,
		Compress:    c.LogCompress,
		Development: c.LogDevelopment,
	}
}
