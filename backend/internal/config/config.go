// internal/config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// Database settings
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`

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

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
