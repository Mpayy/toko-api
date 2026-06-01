package helper

import "github.com/go-playground/validator/v10"

func ValidPrice(field validator.FieldLevel) bool {
	value := field.Field().Int()
	if value%100 == 0 {
		return true
	}
	return false
}