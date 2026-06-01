package app

import (
	"github.com/Mpayy/toko-api/helper"
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("valid_price", helper.ValidPrice)
	return validate
}
