package repository

import "gorm.io/gorm"

type answersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) *answersRepository {
	return &answersRepository{
		db: db,
	}
}
