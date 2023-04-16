package routers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"hu.erettsegizteto/internal/handlers"
)

// NewRouter creates a new http handler
func NewRouter(handler *handlers.Handler) http.Handler {
	router := gin.Default()

	license := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if license == "" {
		return nil
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("erettsegizteto"),
		newrelic.ConfigLicense(license),
	)
	if err != nil {
		panic(err)
	}

	// Add the New Relic middleware
	router.Use(NewRelicMiddleware(app))

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

func NewRelicMiddleware(app *newrelic.Application) gin.HandlerFunc {
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
