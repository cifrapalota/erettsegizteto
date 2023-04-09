package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hu.erettsegizteto/internal/models"
)

func (h *Handler) GetQuestionByID(c *gin.Context) {
	idStr := c.Param("questionID")
	if idStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing 'questionID' path parameter"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid 'questionID' path parameter"})
		return
	}

	ctx := c.Request.Context()
	question, err := h.storage.GetQuestionByID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}

func (h *Handler) GetRandomQuestion(c *gin.Context) {
	ctx := c.Request.Context()

	question, err := h.storage.GetRandomQuestion(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}

type Answer struct {
	AnswerHolderID string `json:"answer_holder_id"`
	Answer         string `json:"answer"`
}

type CheckedAnswer struct {
	AnswerHolderID string   `json:"answer_holder_id"`
	Correct        bool     `json:"correct"`
	Answers        []string `json:"answers"`
}

func (h *Handler) CheckAnswers(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("questionID")
	if idStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing 'questionID' path parameter"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid 'questionID' path parameter"})
		return
	}

	var answers []Answer
	err = json.NewDecoder(c.Request.Body).Decode(&answers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	answerHolders, err := h.storage.GetAnswerHoldersByQuestionID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	checkedAnswers, err := checkAnswers(answerHolders, answers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, checkedAnswers)
}

func checkAnswers(answerHolders []models.AnswerHolder, answers []Answer) ([]CheckedAnswer, error) {
	checkedAnswers := []CheckedAnswer{}

	// Create a map of AnswerHolderID to AnswerHolder for faster lookup
	answerHolderMap := make(map[string]models.AnswerHolder)
	for _, ah := range answerHolders {
		answerHolderMap[ah.ID.String()] = ah
	}

	// Iterate through the Answers and check if they are in the AnswerHolder's Answers list
	for _, answer := range answers {
		answer.Answer = strings.ReplaceAll(answer.Answer, " ", "") // remove spaces from Answer field

		ah, ok := answerHolderMap[answer.AnswerHolderID]
		if !ok {
			return nil, fmt.Errorf("answer holder with ID %s not found", answer.AnswerHolderID)
		}

		// Check if the answer is in the AnswerHolder's Answers list
		found := false
		for _, a := range ah.Answers {
			if a.Answer == answer.Answer {
				found = true
				break
			}
		}

		checkedAnswer := CheckedAnswer{
			AnswerHolderID: ah.ID.String(),
			Correct:        found,
			Answers:        []string{},
		}

		for _, a := range ah.Answers {
			checkedAnswer.Answers = append(checkedAnswer.Answers, a.Answer)
		}

		checkedAnswers = append(checkedAnswers, checkedAnswer)
	}

	return checkedAnswers, nil
}
