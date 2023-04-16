package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		// After the request is processed
		statusCode := c.Writer.Status()
		latency := time.Since(startTime)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   clientIP,
			"method":      method,
			"path":        path,
		}).Info(fmt.Sprintf("%s %s", method, path))
	}
}
