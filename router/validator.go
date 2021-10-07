package router

import "errors"

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	// validator := validator.New()
	// validator.RegisterValidation("multiple_of", validation.MultipleOf)
	// validator.RegisterValidation("valid_subject", validation.ValidSubject)

	return &Validator{}
}

// Validator struct
type Validator struct{}

type Request interface {
	Validate() error
}

// Validate function exposes the validator to Echo
// In this case we want Echo to use the Struct validator
func (v *Validator) Validate(i interface{}) error {
	request, valid := i.(Request)
	if valid != true {
		return errors.New("not a valid request type")
	}
	return request.Validate()
	// return v.validator.Struct(i)
}
