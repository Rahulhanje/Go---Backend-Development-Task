package models

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

// ValidateStruct validates any struct using go-playground/validator
func ValidateStruct(s interface{}) error {
	return Validate.Struct(s)
}
