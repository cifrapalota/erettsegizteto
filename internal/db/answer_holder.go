package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"hu.erettsegizteto/internal/db/models"
)

// GetAnswerHoldersByQuestionID gets all AnswerHolders by questionID
func (db *DB) GetAnswerHoldersByQuestionID(ctx context.Context, questionID uuid.UUID) ([]models.AnswerHolder, error) {
	var answerHolders []models.AnswerHolder

	if err := db.gormDB.WithContext(ctx).Preload("Answers").Where("question_id = ?", questionID).Order("number").Find(&answerHolders).Error; err != nil {
		return nil, fmt.Errorf("couldn't fetch AnswerHolders for the given question ID")
	}

	return answerHolders, nil
}
