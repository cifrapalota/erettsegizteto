package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hu.erettsegizteto/internal/storage"
)

type QuestionHandler struct {
	storage *storage.Storage
}

func NewQuestionHandler(storage *storage.Storage) *QuestionHandler {
	return &QuestionHandler{storage: storage}
}

func (qh *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	question, err := qh.storage.GetQuestionByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(question)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
