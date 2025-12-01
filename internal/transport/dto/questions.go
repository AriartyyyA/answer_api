package dto

import (
	"github/Ariartyyy/answer_api/internal/models"
	"time"
)

// GET /questions/
type GetQuestionsResponse struct {
	Questions []models.Question
}

// POST /questions/
type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type CreateQuestionResponse struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	Created_at time.Time `json:"created_at"`
}

// GET /questions/{id}
type GetQuestionByIDResponse struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	Created_at string `json:"created_at"`
	Answers    []models.Answer
}
