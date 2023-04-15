package models

import "github.com/google/uuid"

// Answer represents a answer in the database
type Answer struct {
	ID             uuid.UUID
	AnswerHolderID uuid.UUID
	Answer         string
}
