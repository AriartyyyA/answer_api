package repository

import (
	"github/Ariartyyy/answer_api/internal/models"

	"gorm.io/gorm"
)

type QuestionsRepository interface {
	GetQuestions() ([]models.Question, error)
	CreateQuestion(text string) (models.Question, error)
	GetQuestionByID(id int) (models.Question, error)
	DeleteQuestionByIDWithAnswers(id int) error
}

type AnswersRepository interface {
	AddAnswerToQuestion(questionID int, text string) (models.Answer, error)
	GetAnswerByID(id int) (models.Answer, error)
	DeleteAnswerByID(id int) error
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
