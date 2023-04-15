package handlers

import (
	"hu.erettsegizteto/internal/db"
)

// Handler represents a request handler
type Handler struct {
	db *db.DB
}

// NewRouter creates a new Handler
func NewHandler(db *db.DB) *Handler {
	return &Handler{db: db}
}
