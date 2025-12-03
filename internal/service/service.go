package service

import (
	"github/Ariartyyy/answer_api/internal/models"
	"github/Ariartyyy/answer_api/internal/repository"
)

type QuestionsService interface {
	GetQuestions() ([]models.Question, error)
	CreateQuestion(text string) (*models.Question, error)
	GetQuestionByID(id int) (*models.Question, error)
	DeleteQuestionByIDWithAnswers(id int) error
}

type AnswerService interface {
	AddAnswerToQuestion(questionID int, userID, text string) (models.Answer, error)
	GetAnswerByID(id int) (models.Answer, error)
	GetAnswersByQuestionID(questionID int) ([]models.Answer, error)
	DeleteAnswerByID(id int) error
}

type Service struct {
	QuestionsService
	AnswerService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		QuestionsService: NewQuestionsService(repo),
		AnswerService:    NewAnswerService(repo),
	}
}
