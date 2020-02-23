package validator

import (
	v "github.com/go-playground/validator/v10"
)

var validator *v.Validate

func init() {
	validator = v.New()
}

func Validate(s interface{}) error {
	return validator.Struct(s)
}
