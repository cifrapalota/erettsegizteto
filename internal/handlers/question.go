package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hu.erettsegizteto/internal/api/conversion"
)

// GetQuestionByID returns a question by id in json
func (h *Handler) GetQuestionByID(c *gin.Context) {
	idStr := c.Param("questionID")
	if idStr == "" {
		h.Logger.Error("Missing 'questionID' path parameter")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing 'questionID' path parameter"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Logger.Errorf("Invalid 'questionID' path parameter: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid 'questionID' path parameter"})
		return
	}

	ctx := c.Request.Context()
	question, err := h.db.GetQuestionByID(ctx, id)
	if err != nil {
		h.Logger.Errorf("Error fetching question by ID: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, conversion.DBQuestionToAPIQuestion(question))
}

// GetRandomQuestion returns a random question in json
func (h *Handler) GetRandomQuestion(c *gin.Context) {
	ctx := c.Request.Context()

	question, err := h.db.GetRandomQuestion(ctx)
	if err != nil {
		h.Logger.Errorf("Error fetching random question: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, conversion.DBQuestionToAPIQuestion(question))
}
