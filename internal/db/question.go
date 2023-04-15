package db

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

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
	var count int64

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Count the total number of valid questions
	if err := db.gormDB.WithContext(ctx).Model(&models.Question{}).Where("valid = TRUE").Count(&count).Error; err != nil {
		return nil, fmt.Errorf("couldn't count the valid questions: %v", err)
	}

	// If there are no valid questions, return an error
	if count == 0 {
		return nil, fmt.Errorf("no valid questions found")
	}

	// Generate a random offset within the range of the total count
	offset := rand.Int63n(count)

	// Retrieve the question at the random offset
	if err := db.gormDB.WithContext(ctx).Preload("AnswerHolders").Where("valid = TRUE").Offset(int(offset)).Limit(1).Find(&question).Error; err != nil {
		return nil, fmt.Errorf("couldn't fetch a random question: %v", err)
	}

	return &question, nil
}
