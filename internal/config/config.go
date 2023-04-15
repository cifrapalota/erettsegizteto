package config

import (
	"os"
)

// Config represents a service configuration
type Config struct {
	DatabaseDSN string
	Port        string
}

// LoadConfig loads the configuration of the service
func LoadConfig() (*Config, error) {
	databaseDSN := os.Getenv("DATABASE_URL")
	if databaseDSN == "" {
		return nil, os.ErrInvalid
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, os.ErrInvalid
	}

	return &Config{
		DatabaseDSN: databaseDSN,
		Port:        port,
	}, nil
}
