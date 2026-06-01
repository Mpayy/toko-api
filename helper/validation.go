package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationError(errors validator.ValidationErrors) map[string]string {
	fieldError := make(map[string]string)
	for _, e := range errors {
		var message string
		switch e.Tag() {
		case "required":
			message = "must be filled"
		case "email":
			message = "must be a valid email"
		case "min":
			message = "must be at least " + e.Param() + " characters long"
		case "max":
			message = "must be at most " + e.Param() + " characters long"
		case "numeric":
			message = "must be a number"
		case "gt":
			message = "must be greater than " + e.Param()
		case "valid_price":
			message = "price must be multiples of 100"
		default:
			message = "invalid input value"
		}
		fieldError[strings.ToLower(e.Field())] = message
	}
	return fieldError
}