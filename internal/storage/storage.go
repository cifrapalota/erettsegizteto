package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"hu.erettsegizteto/internal/models"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "db_erettsegizteto.", // schema name
		},
	})
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

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

// func (s *Storage) GetRandomQuestion(ctx context.Context) (*models.Question, error) {

// 	// Get the question at the random index
// 	var question models.Question
// 	if err := s.db.WithContext(ctx).Limit(1).Order("RANDOM()").Find(&question).Error; err != nil {
// 		return nil, fmt.Errorf("couldn't fetch a random question")
// 	}

// 	return &question, nil
// }

// GetAnswerHoldersByQuestionID fetches answerholder objects by questionid
func GetAnswerHoldersByQuestionID(db *gorm.DB, questionID uuid.UUID) ([]*models.AnswerHolder, error) {
	var answerHolders []*models.AnswerHolder
	err := db.Where("questions_id = ?", questionID).Find(&answerHolders).Error
	return answerHolders, err
}

func (s *Storage) GetRandomQuestion(ctx context.Context) (*models.Question, error) {

	// Get the question at the random index
	var question models.Question
	if err := s.db.WithContext(ctx).Where("valid = TRUE").Limit(1).Order("RANDOM()").Find(&question).Error; err != nil {
		return nil, fmt.Errorf("couldn't fetch a random question")
	}

	// Get the AnswerHolders for the question
	var answerHolders []models.AnswerHolder
	if err := s.db.WithContext(ctx).Where("question_id = ?", question.ID).Order("number").Find(&answerHolders).Error; err != nil {
		return nil, fmt.Errorf("couldn't fetch answer holders for the question")
	}

	question.AnswerHolders = answerHolders
	return &question, nil
}
