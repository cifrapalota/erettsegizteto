package handlers

import "github.com/gin-gonic/gin"

func (qh *Handler) IndexHandler(c *gin.Context) {
	c.File("frontend/templates/index.html")
}
