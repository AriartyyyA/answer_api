package repository

import "gorm.io/gorm"

type QuestionsRepository interface {
	//
}

type AnswersRepository interface {
}

type Repository struct {
	QuestionsRepository
	AnswersRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		QuestionsRepository: NewQuestionsRepository(db),
		AnswersRepository:   NewAnswersRepository(db),
	}
}
