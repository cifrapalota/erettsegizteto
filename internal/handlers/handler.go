package handlers

import (
	"github.com/sirupsen/logrus"
	"hu.erettsegizteto/internal/db"
)

// Handler represents a request handler
type Handler struct {
	db     *db.DB
	Logger *logrus.Logger
}

// NewRouter creates a new Handler
func NewHandler(db *db.DB, logger *logrus.Logger) *Handler {
	return &Handler{
		db:     db,
		Logger: logger,
	}
}
