package models

type Answer struct {
	ID          int    `json:"id"`
	Question_ID int    `json:"question_id"`
	UserID      string `json:"user_id"`
	Text        string `json:"text"`
	Created_at  string `json:"created_at"`
}
