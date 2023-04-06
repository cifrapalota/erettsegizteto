package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (qh *Handler) GetQuestionByID(c *gin.Context) {
	idStr := c.Param("questionID")
	if idStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing 'questionID' path parameter"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' query parameter"})
		return
	}

	ctx := c.Request.Context()
	question, err := qh.storage.GetQuestionByID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	question.Question = strings.Replace(question.Question, "{image_link}", *question.ImageLink, 1)

	c.JSON(http.StatusOK, question)
}

func (qh *Handler) GetRandomQuestion(c *gin.Context) {
	ctx := c.Request.Context()

	count, err := qh.storage.GetQuestionCount(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	randID := rand.Intn(int(count)) + 1

	question, err := qh.storage.GetQuestionByID(ctx, randID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	question.Question = strings.Replace(question.Question, "{image_link}", *question.ImageLink, 1)

	c.JSON(http.StatusOK, question)
}
