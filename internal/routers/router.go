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
	mux.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("frontend/assets/css/"))))
	mux.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("frontend/assets/js/"))))
	return mux
}
