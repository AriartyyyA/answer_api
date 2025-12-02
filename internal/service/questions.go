package service

import "github/Ariartyyy/answer_api/internal/models"

type questionsService struct {
	//
}

func NewQuestionsService() QuestionsService {
	return &questionsService{
		//
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
