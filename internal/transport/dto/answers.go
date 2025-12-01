package dto

import "time"

// POST /questions/{question_id}/answers
type AddAnswerRequest struct {
	Text string `json:"text"`
}

// Используем во многих ответах
type AnswerResponse struct {
	ID          int       `json:"id"`
	Question_ID int       `json:"question_id"`
	User_ID     string    `json:"user_id"`
	Text        string    `json:"text"`
	Created_at  time.Time `json:"created_at"`
}
