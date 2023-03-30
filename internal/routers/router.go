package routers

import (
	"net/http"

	"hu.erettsegizteto/internal/handlers"
	"hu.erettsegizteto/internal/storage"
)

func NewRouter(storage *storage.Storage) http.Handler {
	questionHandler := handlers.NewQuestionHandler(storage)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/question", questionHandler.GetQuestionByID)

	return mux
}
