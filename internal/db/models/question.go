package models

import "github.com/google/uuid"

// Question represents a question in the database
type Question struct {
	ID            uuid.UUID      `json:"id"`
	Question      string         `json:"question"`
	Generated     bool           `json:"generated"`
	Year          *int           `json:"year,omitempty"`
	Semester      *int           `json:"semester,omitempty"`
	Number        *int           `json:"number,omitempty"`
	ImageLink     string         `json:"imageLink,omitempty"`
	Solution      string         `json:"solution,omitempty"`
	AnswerHolders []AnswerHolder `json:"answerHolders" gorm:"ForeignKey:question_id"`
}
