package transport

import (
	"encoding/json"
	"github/Ariartyyy/answer_api/internal/transport/dto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *HTTPHandlers) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.service.GetQuestions()
	if err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQuestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	question, err := h.service.CreateQuestion(req.Text)
	if err != nil {
		WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.CreateQuestionResponse{
		ID:         question.ID,
		Text:       question.Text,
		Created_at: question.Created_at,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) GetQuestionByIDWithAnswers(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		WriteJSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	question, err := h.service.GetQuestionByID(id)
	if err != nil {
		WriteJSONError(w, "Question not found", http.StatusNotFound)
		return
	}

	answers, err := h.service.GetAnswersByQuestionID(id)
	if err != nil {
		WriteJSONError(w, "Failed to get answers", http.StatusInternalServerError)
		return
	}

	resp := dto.GetQuestionByIDResponse{
		ID:         question.ID,
		Text:       question.Text,
		Created_at: question.Created_at.Format("2006-01-02T15:04:05Z07:00"),
		Answers:    answers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) DeleteQuestionByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		WriteJSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteQuestionByIDWithAnswers(id); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getIDFromPath(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	return strconv.Atoi(idStr)
}
