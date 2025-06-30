package repository

import (
	"server/internal/domain"
	"server/internal/model"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	GetArticlesByCriteria(articles *[]model.Article, criteria domain.GetArticlesCriteria, userId string) error
	GetArticleById(article *model.Article, userId string, articleId string) error
	CreateArticle(article *model.Article) error
	UpdateArticle(article *model.Article, userId string, articleId string) error
	DeleteArticle(userId string, articleId string) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) GetArticlesByCriteria(articles *[]model.Article, criteria domain.GetArticlesCriteria, userId string) error {
	query := ar.db.Model(&model.Article{}).Where("user_id = ?", userId)
	if criteria.KeyWord != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+criteria.KeyWord+"%", "%"+criteria.KeyWord+"%")
	}
	if criteria.Status != nil {
		query = query.Where("status = ?", *criteria.Status)
	}
	if criteria.Category != nil {
		query = query.Where("category = ?", *criteria.Category)
	}
	if err := query.Find(articles).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) CreateArticle(article *model.Article) error {
	if err := ar.db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) GetArticleById(article *model.Article, userId string, articleId string) error {
	if err := ar.db.Joins("Users").Where("user_id=?", userId).Find(article, articleId).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) UpdateArticle(article *model.Article, userId string, articleId string) error {
	if err := ar.db.Model(&model.Article{}).Where("user_id = ? AND id = ?", userId, articleId).Updates(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) DeleteArticle(userId string, articleId string) error {
	if err := ar.db.Where("user_id = ? AND id = ?", userId, articleId).Delete(&model.Article{}).Error; err != nil {
		return err
	}
	return nil
}
