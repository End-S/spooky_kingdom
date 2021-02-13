package router

import (
	"github.com/End-S/spooky_kingdom/router/validation"
	"github.com/go-playground/validator/v10"
)

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	validator := validator.New()
	validator.RegisterValidation("multiple_of", validation.MultipleOf)
	validator.RegisterValidation("valid_subject", validation.ValidSubject)

	return &Validator{
		validator: validator,
	}
}

// Validator struct
type Validator struct {
	validator *validator.Validate
}

// Validate function exposes the validator to Echo
// In this case we want Echo to use the Struct validator
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
