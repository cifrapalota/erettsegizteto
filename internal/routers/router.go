package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hu.erettsegizteto/internal/handlers"
)

// NewRouter creates a new http handler
func NewRouter(handler *handlers.Handler) http.Handler {
	router := gin.Default()

	// Add the New Relic middleware
	router.Use(handlers.NewRelicMiddleware(handler.NewRelic))

	// Add the logging middleware
	router.Use(handlers.LoggingMiddleware(handler.Logger))

	// Endpoints
	router.GET("/question/random", handler.GetRandomQuestion)
	router.POST("/question/:questionID/check-answers", handler.CheckAnswers)

	// Static assets
	router.StaticFile("/", "frontend/templates/index.html")
	router.Static("/assets/css", "frontend/assets/css")
	router.Static("/assets/js", "frontend/assets/js")
	router.Static("/assets/img", "frontend/assets/img")

	return router
}
