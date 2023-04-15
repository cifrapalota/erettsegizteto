package models

import "github.com/google/uuid"

// Answer represents a answer in the database
type Answer struct {
	ID             uuid.UUID `json:"id"`
	AnswerHolderID uuid.UUID `json:"answer_holder_id,omitempty"`
	Answer         string    `json:"answer"`
}
