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

func (a *answersRepository) AddAnswerToQuestion(questionID int, userID, text string) (*models.Answer, error) {
	answer := &models.Answer{
		Question_ID: questionID,
		UserID:      userID,
		Text:        text,
	}

	if err := a.db.Create(answer).Error; err != nil {
		return nil, err
	}

	return answer, nil
}

func (a *answersRepository) DeleteAnswerByID(id int) error {
	if err := a.db.Delete(&models.Answer{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *answersRepository) GetAnswerByID(id int) (*models.Answer, error) {
	var answer models.Answer

	if err := a.db.First(&answer, id).Error; err != nil {
		return nil, err
	}

	return &answer, nil
}

func (a *answersRepository) GetAnswersByQuestionID(questionID int) ([]models.Answer, error) {
	var answers []models.Answer

	if err := a.db.Where("question_id = ?", questionID).Order("created_at ASC").Find(&answers).Error; err != nil {
		return nil, err
	}

	return answers, nil
}
