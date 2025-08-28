package config

import (
	"os"
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBName      string
	DBHost      string
	DBPort      string
	HTTPAddress string
}

func Load() (*Config, error) {
	httpAddress := os.Getenv("HTTP_ADDRESS")
	if httpAddress == "" {
		httpAddress = ":8080"
	}

	cfg := &Config{
		DBUser:      os.Getenv("POSTGRES_USER"),
		DBPassword:  os.Getenv("POSTGRES_PASSWORD"),
		DBName:      os.Getenv("POSTGRES_DB"),
		DBHost:      os.Getenv("POSTGRES_HOST"),
		DBPort:      os.Getenv("POSTGRES_PORT"),
		HTTPAddress: httpAddress,
	}

	return cfg, nil
}
