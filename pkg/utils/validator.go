package utils

import (
	"fmt"

	"github.com/andreylm/bmlabs/pkg/utils/search"
	v "github.com/go-playground/validator/v10"
)

var validator *v.Validate

func init() {
	validator = v.New()
}

// Validate - validate struct
func Validate(s interface{}) map[string]string {
	errors := make(map[string]string)
	err := validator.Struct(s)
	if errs, ok := err.(v.ValidationErrors); ok {
		for _, e := range errs {
			errors[e.Field()] = e.ActualTag()
		}
	}

	return errors
}

// ValidateSearch - validates search form
func ValidateSearch(s *search.Search) error {
	for _, f := range s.Filters {
		if !f.IsValidType() {
			return fmt.Errorf("invalid filter type: %v", f.Type)
		}
		if !f.IsValidOperator() {
			return fmt.Errorf("invalid operator: %v", f.Operator)
		}
	}
	return nil
}
