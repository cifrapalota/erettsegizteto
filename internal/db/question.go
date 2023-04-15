package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"hu.erettsegizteto/internal/db/models"
)

// GetQuestionByID gets a Question by id
func (db *DB) GetQuestionByID(ctx context.Context, id uuid.UUID) (*models.Question, error) {
	var question models.Question

	err := db.gormDB.WithContext(ctx).First(&question, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("question with ID %d not found", id)
		}
		return nil, err
	}

	return &question, nil
}

// GetQuestionByID gets a random valid question
func (db *DB) GetRandomQuestion(ctx context.Context) (*models.Question, error) {
	var question models.Question

	if err := db.gormDB.WithContext(ctx).Preload("AnswerHolders").Where("valid = TRUE").Limit(1).Order("RANDOM()").Find(&question).Error; err != nil {
		return nil, fmt.Errorf("couldn't fetch a random question")
	}

	return &question, nil
}
