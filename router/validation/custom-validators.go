package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// MultipleOf function checks the current fields value is a multiple of the params value
func MultipleOf(fl validator.FieldLevel) bool {

	field := fl.Field()
	kind := field.Kind()

	currentField, currentKind, ok := fl.GetStructFieldOK()
	if !ok || currentKind != kind {
		return false
	}

	// avoid devision by zero
	if currentField.Int() == 0 {
		return true
	}

	switch kind {

	case reflect.Int:

		return (field.Int() % currentField.Int()) == 0

	}

	// default
	return false
}

// ValidSubject function checks the subject is one we recognise
func ValidSubject(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	validSubjects := [3]string{"ghost", "ufo", "weird"}

	var v string
	switch kind {
	case reflect.String:
		v = field.String()
	default:
		return false
	}

	for i := 0; i < len(validSubjects); i++ {
		if validSubjects[i] == v {
			return true
		}
	}

	return false
}
