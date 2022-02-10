package validations

import (
	"github.com/MET-DEV/api-project/models"
	"gopkg.in/go-playground/validator.v9"
)

func ProductValidator(data models.Product) validator.FieldError {
	v := validator.New()

	err := v.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return e
		}
	}

	return nil
}
func CategoryValidator(data models.Category) validator.FieldError {
	v := validator.New()

	err := v.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return e
		}
	}

	return nil
}

func UserValidator(data models.User) validator.FieldError {
	v := validator.New()

	err := v.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return e
		}
	}

	return nil
}
