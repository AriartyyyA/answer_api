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

func (s *questionsService) CreateQuestion(text string) (*models.Question, error) {
	question, err := s.repo.CreateQuestion(text)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s *questionsService) DeleteQuestionByIDWithAnswers(id int) error {
	return s.repo.DeleteQuestionByIDWithAnswers(id)
}

func (s *questionsService) GetQuestionByID(id int) (*models.Question, error) {
	question, err := s.repo.GetQuestionByID(id)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s *questionsService) GetQuestions() ([]models.Question, error) {
	questions, err := s.repo.GetQuestions()
	if err != nil {
		return nil, err
	}

	return questions, nil
}
