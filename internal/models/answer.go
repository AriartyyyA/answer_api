package models

import "time"

type Answer struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	Question_ID int       `gorm:"column:question_id" json:"question_id"`
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	Text        string    `gorm:"column:text" json:"text"`
	Created_at  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
