package repository

import (
	"github/Ariartyyy/answer_api/internal/models"

	"gorm.io/gorm"
)

type answersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) AnswersRepository {
	return &answersRepository{
		db: db,
	}
}

func (a *answersRepository) AddAnswerToQuestion(questionID int, text string) (models.Answer, error) {
	panic("unimplemented")
}

func (a *answersRepository) DeleteAnswerByID(id int) error {
	panic("unimplemented")
}

func (a *answersRepository) GetAnswerByID(id int) (models.Answer, error) {
	panic("unimplemented")
}
