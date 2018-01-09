package validators

import (
	"reflect"

	"github.com/gothite/forms/codes"
)

// ArrayMinLengthValidator controls minimum length.
type ArrayMinLengthValidator struct {
	Length int
}

// Validate do validation.
func (validator ArrayMinLengthValidator) Validate(value interface{}) (interface{}, *Error) {
	if reflect.ValueOf(value).Len() < validator.Length {
		return value, NewError(codes.MinLength, validator.Length)
	}

	return value, nil
}

// ArrayMinLength initializes MinLengthValidator instance.
func ArrayMinLength(length int) *ArrayMinLengthValidator {
	return &ArrayMinLengthValidator{length}
}

// ArrayMaxLengthValidator controls maximum length.
type ArrayMaxLengthValidator struct {
	Length int
}

// Validate do validation.
func (validator ArrayMaxLengthValidator) Validate(value interface{}) (interface{}, *Error) {
	if reflect.ValueOf(value).Len() > validator.Length {
		return value, NewError(codes.MaxLength, validator.Length)
	}

	return value, nil
}

// ArrayMaxLength initializes MaxLengthValidator instance.
func ArrayMaxLength(length int) *ArrayMaxLengthValidator {
	return &ArrayMaxLengthValidator{length}
}
