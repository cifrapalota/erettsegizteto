package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (qh *Handler) GetQuestionByID(c *gin.Context) {
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
	question, err := qh.storage.GetQuestionByID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}

func (qh *Handler) GetRandomQuestion(c *gin.Context) {
	ctx := c.Request.Context()

	question, err := qh.storage.GetRandomQuestion(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}
