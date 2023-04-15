package config

import (
	"os"
)

type Config struct {
	DatabaseDSN string
	Port        string
}

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
