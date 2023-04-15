package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apiModels "hu.erettsegizteto/internal/api/models"
	dbModels "hu.erettsegizteto/internal/db/models"
)

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

	var answers []apiModels.UserAnswer
	err = json.NewDecoder(c.Request.Body).Decode(&answers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	answerHolders, err := h.db.GetAnswerHoldersByQuestionID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	checkedAnswers, err := checkAnswers(answerHolders, answers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	question, err := h.db.GetQuestionByID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := apiModels.Solution{
		CheckedAnswers: checkedAnswers,
		Solution:       question.Solution,
	}

	c.JSON(http.StatusOK, response)
}

func checkAnswers(answerHolders []dbModels.AnswerHolder, answers []apiModels.UserAnswer) ([]apiModels.CheckedAnswerHolder, error) {
	checkedAnswers := []apiModels.CheckedAnswerHolder{}

	answerHolderMap := make(map[string]dbModels.AnswerHolder)
	for _, ah := range answerHolders {
		answerHolderMap[ah.ID.String()] = ah
	}

	for _, answer := range answers {
		answer.Answer = strings.ReplaceAll(answer.Answer, " ", "")

		ah, ok := answerHolderMap[answer.AnswerHolderID]
		if !ok {
			return nil, fmt.Errorf("answer holder with ID %s not found", answer.AnswerHolderID)
		}

		found := false
		for _, a := range ah.Answers {
			if a.Answer == answer.Answer {
				found = true
				break
			}
		}

		checkedAnswer := apiModels.CheckedAnswerHolder{
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
