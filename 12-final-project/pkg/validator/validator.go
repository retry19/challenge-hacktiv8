package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorValidator struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Tag     string `json:"tag"`
}

func ValidateSchema(data interface{}) []ErrorValidator {
	errors := []ErrorValidator{}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			errors = append(errors, ErrorValidator{
				Field:   er.Field(),
				Message: strings.Split(er.Error(), "Error:")[1],
				Tag:     er.Tag(),
			})
		}
	}

	return errors
}
