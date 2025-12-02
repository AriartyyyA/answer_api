package service

import (
	"github/Ariartyyy/answer_api/internal/models"
	"github/Ariartyyy/answer_api/internal/repository"
)

type questionsService struct {
	repo *repository.Repository
}

func NewQuestionsService(repo *repository.Repository) QuestionsService {
	return &questionsService{
		repo: repo,
	}
}

func (q *questionsService) CreateQuestion(text string) (models.Question, error) {
	panic("unimplemented")
}

func (q *questionsService) DeleteQuestionByIDWithAnswers(id int) error {
	panic("unimplemented")
}

func (q *questionsService) GetQuestionByID(id int) (models.Question, error) {
	panic("unimplemented")
}

func (q *questionsService) GetQuestions() ([]models.Question, error) {
	panic("unimplemented")
}
