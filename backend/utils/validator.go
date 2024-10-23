package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetValidationErrors(err error) []ErrorMsg {
	var errors []ErrorMsg

	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorMsg
		element.Field = err.Field()

		switch err.Tag() {
		case "required":
			element.Message = fmt.Sprintf("%s is required", err.Field())
		case "email":
			element.Message = "Invalid email format"
		case "min":
			element.Message = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
		case "max":
			element.Message = fmt.Sprintf("%s must not exceed %s characters", err.Field(), err.Param())
		default:
			element.Message = fmt.Sprintf("Validation failed on %s with tag %s", err.Field(), err.Tag())
		}

		errors = append(errors, element)
	}

	return errors
}
