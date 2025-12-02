package repository

import "gorm.io/gorm"

type questionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) QuestionsRepository {
	return &questionsRepository{
		db: db,
	}
}
