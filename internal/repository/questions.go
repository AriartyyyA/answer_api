package repository

import (
	"github/Ariartyyy/answer_api/internal/models"

	"gorm.io/gorm"
)

type questionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) QuestionsRepository {
	return &questionsRepository{
		db: db,
	}
}

func (q *questionsRepository) CreateQuestion(text string) (models.Question, error) {
	panic("unimplemented")
}

func (q *questionsRepository) DeleteQuestionByIDWithAnswers(id int) error {
	panic("unimplemented")
}

func (q *questionsRepository) GetQuestionByID(id int) (models.Question, error) {
	panic("unimplemented")
}

func (q *questionsRepository) GetQuestions() ([]models.Question, error) {
	panic("unimplemented")
}
