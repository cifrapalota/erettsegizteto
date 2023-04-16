package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
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

func NewRelicMiddleware(app *newrelic.Application) gin.HandlerFunc {
	if app == nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path)
		defer txn.End()

		txn.SetWebRequestHTTP(c.Request)
		c.Request = newrelic.RequestWithTransactionContext(c.Request, txn)

		c.Next()

		txn.SetWebResponse(c.Writer)
		txn.SetName(c.HandlerName())
	}
}
