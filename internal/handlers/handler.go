package handlers

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"hu.erettsegizteto/internal/db"
)

// Handler represents a request handler
type Handler struct {
	db       *db.DB
	Logger   *logrus.Logger
	NewRelic *newrelic.Application
}

// NewRouter creates a new Handler
func NewHandler(db *db.DB, logger *logrus.Logger, newRelic *newrelic.Application) *Handler {
	return &Handler{
		db:       db,
		Logger:   logger,
		NewRelic: newRelic,
	}
}
