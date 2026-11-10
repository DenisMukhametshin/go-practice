package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port string
	Db   PostgresConfig
}

type PostgresConfig struct {
	Password string
	Port     int
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) != 0 {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: getEnv("PORT", "8080"),
		Db: PostgresConfig{
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
	}
	return cfg, nil
}
