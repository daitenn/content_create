package usecase

import (
	"go-restapi/model"
	"go-restapi/repository"
	"go-restapi/validator"
)

type IContentUsecase interface {
	GetAllContents(userId uint) ([]model.ContentResponse, error)
	GetContentById(userId uint, contentId uint) (model.ContentResponse, error)
	CreateContent(content model.Content) (model.ContentResponse, error)
	UpdateContent(content model.Content, userId uint, contentId uint) (model.ContentResponse, error)
	DeleteContent(userId uint, contentId uint) error
}

type contentUsecase struct {
	cr repository.IContentRepository
	cv validator.IContentValidator
}

func NewContentUsecase(cr repository.IContentRepository, cv validator.IContentValidator) IContentUsecase {
	return &contentUsecase{cr, cv}
}

func (cu *contentUsecase) GetAllContents(userId uint) ([]model.ContentResponse, error) {
	contents := []model.Content{}
	if err := cu.cr.GetAllContents(&contents, userId); err != nil {
		return nil, err
	}
	resContents := []model.ContentResponse{}
	for _, v := range contents {
		c := model.ContentResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resContents = append(resContents, c)
	}
	return resContents, nil
}

func (cu *contentUsecase) GetContentById(userId uint, contentId uint) (model.ContentResponse, error) {
	content := model.Content{}
	if err := cu.cr.GetContentById(&content, userId, contentId); err != nil {
		return model.ContentResponse{}, err
	}
	resContent := model.ContentResponse{
		ID:        content.ID,
		Title:     content.Title,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
	}
	return resContent, nil
}

func (cu *contentUsecase) CreateContent(content model.Content) (model.ContentResponse, error) {
	if err := cu.cv.ContentValidate(content); err != nil {
		return model.ContentResponse{}, err
	}
	if err := cu.cr.CreateContent(&content); err != nil {
		return model.ContentResponse{}, err
	}
	resContent := model.ContentResponse{
		ID:        content.ID,
		Title:     content.Title,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
	}
	return resContent, nil
}

func (cu *contentUsecase) UpdateContent(content model.Content, userId uint, contentId uint) (model.ContentResponse, error) {
	if err := cu.cv.ContentValidate(content); err != nil {
		return model.ContentResponse{}, err
	}
	if err := cu.cr.UpdateContent(&content, userId, contentId); err != nil {
		return model.ContentResponse{}, err
	}
	resContent := model.ContentResponse{
		ID:        content.ID,
		Title:     content.Title,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
	}
	return resContent, nil
}

func (cu *contentUsecase) DeleteContent(userId uint, contentId uint) error {
	if err := cu.cr.DeleteContent(userId, contentId); err != nil {
		return err
	}
	return nil
}
