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

func (a *answerService) AddAnswerToQuestion(questionID int, userID, text string) (models.Answer, error) {
	answer, err := a.repo.AddAnswerToQuestion(questionID, userID, text)
	if err != nil {
		return models.Answer{}, err
	}

	return *answer, nil
}

func (a *answerService) DeleteAnswerByID(id int) error {
	return a.repo.DeleteAnswerByID(id)
}

func (a *answerService) GetAnswerByID(id int) (models.Answer, error) {
	answer, err := a.repo.GetAnswerByID(id)
	if err != nil {
		return models.Answer{}, err
	}

	return *answer, nil
}

func (a *answerService) GetAnswersByQuestionID(questionID int) ([]models.Answer, error) {
	answers, err := a.repo.GetAnswersByQuestionID(questionID)
	if err != nil {
		return nil, err
	}

	return answers, nil
}
