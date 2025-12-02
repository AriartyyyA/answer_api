package service

import (
	"github/Ariartyyy/answer_api/internal/models"
	"github/Ariartyyy/answer_api/internal/repository"
)

type answerService struct {
	repo *repository.Repository
}

func NewAnswerService(repo *repository.Repository) AnswerService {
	return &answerService{
		repo: repo,
	}
}

func (a *answerService) AddAnswerToQuestion(questionID int, text string) (models.Answer, error) {
	panic("unimplemented")
}

func (a *answerService) DeleteAnswerByID(id int) error {
	panic("unimplemented")
}

func (a *answerService) GetAnswerByID(id int) (models.Answer, error) {
	panic("unimplemented")
}
