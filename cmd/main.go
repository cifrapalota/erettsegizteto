package main

import (
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"hu.erettsegizteto/internal/config"
	"hu.erettsegizteto/internal/db"
	"hu.erettsegizteto/internal/handlers"
	"hu.erettsegizteto/internal/routers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := db.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize db: %v", err)
	}

	logger := logrus.New()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.AppName),
		newrelic.ConfigLicense(cfg.NewRelicLicenseKey),
	)
	if err != nil {
		logger.Errorf("Failed to initialize newRelic: %v", err)
	}

	handler := handlers.NewHandler(db, logger, app)

	router := routers.NewRouter(handler)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
