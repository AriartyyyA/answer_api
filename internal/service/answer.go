package service

import "github/Ariartyyy/answer_api/internal/models"

type answerService struct {
	//
}

func NewAnswerService() AnswerService {
	return &answerService{
		//
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
