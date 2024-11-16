// internal/config/config.go
package config

import (
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
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("DB_MAX_OPEN_CONNS", 25)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 10)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", 15)   // minutes
	viper.SetDefault("DB_CONN_MAX_IDLE_TIME", 5)   // minutes
	viper.SetDefault("DB_HEALTH_CHECK_PERIOD", 30) // seconds
	viper.SetDefault("DB_CONNECT_RETRIES", 5)
	viper.SetDefault("DB_CONNECT_RETRY_DELAY", 5) // seconds

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
