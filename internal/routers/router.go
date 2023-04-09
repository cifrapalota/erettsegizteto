package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hu.erettsegizteto/internal/handlers"
)

func NewRouter(handler *handlers.Handler) http.Handler {
	router := gin.Default()

	//Endpoints
	router.GET("/", handler.IndexHandler)
	router.GET("/question/:questionID", handler.GetQuestionByID)
	router.GET("/question/random", handler.GetRandomQuestion)
	router.POST("/question/:questionID/check_answers", handler.CheckAnswers)

	//Static assets
	router.Static("/assets/css", "frontend/assets/css")
	router.Static("/assets/js", "frontend/assets/js")
	router.Static("/assets/img", "frontend/assets/img")

	return router
}
