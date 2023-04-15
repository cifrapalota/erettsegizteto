package models

import "github.com/google/uuid"

// AnswerHolder represents an answerHolder in the database
type AnswerHolder struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	Prefix     string
	Suffix     string
	Number     int
	Help       string
	Answers    []Answer `gorm:"ForeignKey:answer_holder_id"`
}
