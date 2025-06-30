package repository

import (
	"server/internal/model"

	"gorm.io/gorm"
)

type IQuizRepository interface {
	GetQuizById(quiz *model.Quiz, quizId string) error
	CreateQuiz(quiz *model.Quiz) error
	UpdateQuiz(quiz *model.Quiz, quizId string) error
	DeleteQuiz(quizId string) error
	GetQuizOptionsByQuizId(options *[]model.QuizOption, quizId string) error
	CreateQuizOption(option *model.QuizOption) error
	UpdateQuizOption(option *model.QuizOption, quizId string, number uint) error
	DeleteQuizOption(quizId string, number uint) error
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) IQuizRepository {
	return &quizRepository{db}
}
func (qr *quizRepository) GetQuizById(quiz *model.Quiz, quizId string) error {
	if err := qr.db.Where("id = ?", quizId).First(quiz).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) CreateQuiz(quiz *model.Quiz) error {
	if err := qr.db.Create(quiz).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) UpdateQuiz(quiz *model.Quiz, quizId string) error {
	if err := qr.db.Model(&model.Quiz{}).Where("id = ?", quizId).Updates(quiz).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) DeleteQuiz(quizId string) error {
	if err := qr.db.Where("id = ?", quizId).Delete(&model.Quiz{}).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) GetQuizOptionsByQuizId(options *[]model.QuizOption, quizId string) error {
	if err := qr.db.Where("quiz_id = ?", quizId).Find(options).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) CreateQuizOption(option *model.QuizOption) error {
	if err := qr.db.Create(option).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) UpdateQuizOption(option *model.QuizOption, quizId string, number uint) error {
	if err := qr.db.Model(&model.QuizOption{}).Where("quiz_id = ? AND number = ?", quizId, number).Updates(option).Error; err != nil {
		return err
	}
	return nil
}
func (qr *quizRepository) DeleteQuizOption(quizId string, number uint) error {
	if err := qr.db.Where("quiz_id = ? AND number = ?", quizId, number).Delete(&model.QuizOption{}).Error; err != nil {
		return err
	}
	return nil
}