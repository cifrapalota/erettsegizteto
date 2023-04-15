package models

import "github.com/google/uuid"

type Question struct {
	ID            uuid.UUID      `json:"id"`
	Question      string         `json:"question"`
	Generated     bool           `json:"generated"`
	Year          *int           `json:"year,omitempty"`
	Semester      *int           `json:"semester,omitempty"`
	Number        *int           `json:"number,omitempty"`
	AnswerHolders []AnswerHolder `json:"answerHolders"`
}
