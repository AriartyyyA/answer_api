package models

import "time"

type Question struct {
	ID         int       `gorm:"primaryKey;column:id" json:"id"`
	Text       string    `gorm:"column:text" json:"text"`
	Created_at time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
