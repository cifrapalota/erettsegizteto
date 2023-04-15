package models

import "github.com/google/uuid"

// Question represents a question in the database
type Question struct {
	ID            uuid.UUID
	Question      string
	Generated     bool
	Year          *int
	Semester      *int
	Number        *int
	Workings      string
	AnswerHolders []AnswerHolder `gorm:"ForeignKey:question_id"`
}
