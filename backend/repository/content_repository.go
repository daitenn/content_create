package repository

import (
	"fmt"
	"go-restapi/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IContentRepository interface {
	GetAllContents(contents *[]model.Content, userId uint) error
	GetContentById(content *model.Content, userId uint, contentId uint) error
	CreateContent(content *model.Content) error
	UpdateContent(content *model.Content, userId uint, contentId uint) error
	DeleteContent(userId uint, contentId uint) error
}

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) IContentRepository {
	return &contentRepository{db}
}

func (cr *contentRepository) GetAllContents(contents *[]model.Content, userId uint) error {
	if err := cr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(contents).Error; err != nil {
		return err
	}
	return nil
}

func (cr *contentRepository) GetContentById(content *model.Content, userId uint, contentId uint) error {
	if err := cr.db.Joins("User").Where("user_id=?", userId).First(content, contentId).Error; err != nil {
		return err
	}
	return nil
}

func (cr *contentRepository) CreateContent(content *model.Content) error {
	if err := cr.db.Create(content).Error; err != nil {
		return err
	}
	return nil
}

func (cr *contentRepository) UpdateContent(content *model.Content, userId uint, contentId uint) error {
	result := cr.db.Model(content).Clauses(clause.Returning{}).Where("id=? AND user_id=?", contentId, userId).Update("title", content.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (cr *contentRepository) DeleteContent(userId uint, contentId uint) error {
	result := cr.db.Where("id=? AND user_id=?", contentId, userId).Delete(&model.Content{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
