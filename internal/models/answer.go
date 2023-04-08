package models

import "github.com/google/uuid"

// Answer represents a answer in the database
type Answer struct {
	ID     uuid.UUID `json:"id"`
	Answer string    `json:"question"`
}
