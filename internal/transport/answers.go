package transport

import (
	"encoding/json"
	"github/Ariartyyy/answer_api/internal/transport/dto"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (h *HTTPHandlers) AddAnswerToQuestionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionIDStr := vars["question_id"]
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		WriteJSONError(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var req dto.AddAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Text) == "" {
		WriteJSONError(w, "Answer text cannot be empty", http.StatusBadRequest)
		return
	}

	if req.UserID <= "" {
		WriteJSONError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	answer, err := h.service.AddAnswerToQuestion(questionID, req.UserID, req.Text)
	if err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.AnswerResponse{
		ID:          answer.ID,
		Question_ID: answer.Question_ID,
		User_ID:     answer.UserID,
		Text:        answer.Text,
		Created_at:  answer.Created_at,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) GetAnswerByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteJSONError(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	answer, err := h.service.GetAnswerByID(id)
	if err != nil {
		WriteJSONError(w, "Answer not found", http.StatusNotFound)
		return
	}

	resp := dto.AnswerResponse{
		ID:          answer.ID,
		Question_ID: answer.Question_ID,
		User_ID:     answer.UserID,
		Text:        answer.Text,
		Created_at:  answer.Created_at,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) DeleteAnswerByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteJSONError(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteAnswerByID(id); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
