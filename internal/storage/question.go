package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"hu.erettsegizteto/internal/models"
)

func (s *Storage) GetQuestionByID(ctx context.Context, id uuid.UUID) (*models.Question, error) {
	var question models.Question
	err := s.db.WithContext(ctx).First(&question, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("question with ID %d not found", id)
		}
		return nil, err
	}

	return &question, nil
}

func (s *Storage) GetRandomQuestion(ctx context.Context) (*models.Question, error) {

	// Get the question at the random index
	var question models.Question
	if err := s.db.WithContext(ctx).Preload("AnswerHolders").Where("valid = TRUE").Limit(1).Order("RANDOM()").Find(&question).Error; err != nil {
		return nil, fmt.Errorf("couldn't fetch a random question")
	}

	return &question, nil
}
