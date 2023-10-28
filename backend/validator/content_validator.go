package validator

import (
	"go-restapi/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IContentValidator interface {
	ContentValidate(content model.Content) error
}

type contentValidator struct{}

func NewContentValidate() IContentValidator {
	return &contentValidator{}
}

func (cv *contentValidator) ContentValidate(content model.Content) error {
	return validation.ValidateStruct(
		&content,
		validation.Field(
			&content.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}
