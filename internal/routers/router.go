package routers

import (
	"net/http"

	"hu.erettsegizteto/internal/handlers"
)

func NewRouter(handler *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	//Endpoints
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/question", handler.GetQuestionByID)

	//Static assets
	mux.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("frontend/assets/css/"))))
	mux.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("frontend/assets/js/"))))

	return mux
}
