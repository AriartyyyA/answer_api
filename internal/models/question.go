package models

import "time"

type Question struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	Created_at time.Time `json:"created_at"`
}
