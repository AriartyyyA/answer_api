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

func (r *questionsRepository) CreateQuestion(text string) (*models.Question, error) {
	question := &models.Question{
		Text: text,
	}

	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *questionsRepository) DeleteQuestionByIDWithAnswers(id int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("question_id = ?", id).Delete(&models.Answer{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Question{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *questionsRepository) GetQuestionByID(id int) (*models.Question, error) {
	var question models.Question

	if err := r.db.First(&question, id).Error; err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *questionsRepository) GetQuestions() ([]models.Question, error) {
	var questions []models.Question

	if err := r.db.Order("created_at DESC").Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}
