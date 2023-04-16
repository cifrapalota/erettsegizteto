package config

import (
	"os"
)

// Config represents a service configuration
type Config struct {
	AppName            string
	DatabaseURL        string
	Port               string
	NewRelicLicenseKey string
}

// LoadConfig loads the configuration of the service
func LoadConfig() (*Config, error) {
	DatabaseURL := os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		return nil, os.ErrInvalid
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, os.ErrInvalid
	}

	newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")

	return &Config{
		AppName:            "erettsegizteto",
		DatabaseURL:        DatabaseURL,
		Port:               port,
		NewRelicLicenseKey: newRelicLicenseKey,
	}, nil
}
