package handlers

import (
	"hu.erettsegizteto/internal/db"
)

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	return &Handler{db: db}
}
