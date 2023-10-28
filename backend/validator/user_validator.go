package validator

import (
	"go-restapi/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(
			&user.Name,
			validation.Required.Error("Name is required"),
			validation.RuneLength(1,30).Error("limited max 30"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6,20).Error("limited between 6 chars and 20 chars"),
		),
	)
}
