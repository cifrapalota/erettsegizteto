package models

import "github.com/google/uuid"

// Question represents a question in the database
type AnswerHolder struct {
	ID         uuid.UUID `json:"id"`
	QuestionID uuid.UUID `json:"question_id,omitempty"`
	Prefix     string    `json:"prefix,omitempty"`
	Suffix     string    `json:"suffix,omitempty"`
	Number     int       `json:"number"`
	Help       string    `json:"help,omitempty"`
	Answers    []Answer  `json:"answers,omitempty" gorm:"ForeignKey:answer_holder_id"`
}
