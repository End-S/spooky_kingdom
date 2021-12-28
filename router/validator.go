package router

import "errors"

// NewValidator creates a new validator instance
func NewValidator() *Validator {
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
}
