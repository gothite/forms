package validators

import "github.com/gothite/forms/codes"

type IntegerValidator interface {
	Validate(value int) (int, *Error)
}

// IntegerMinValueValidator controls minimum value.
type IntegerMinValueValidator struct {
	Value int
}

// Validate do validation.
func (validator IntegerMinValueValidator) Validate(value int) (int, *Error) {
	if value < validator.Value {
		return value, NewError(codes.MinValue, validator.Value)
	}

	return value, nil
}

// IntegerMinValue initializes IntegerMinValueValidator instance.
func IntegerMinValue(value int) *IntegerMinValueValidator {
	return &IntegerMinValueValidator{value}
}

// IntegerMaxValueValidator controls maximum value.
type IntegerMaxValueValidator struct {
	Value int
}

// Validate do validation.
func (validator IntegerMaxValueValidator) Validate(value int) (int, *Error) {
	if value > validator.Value {
		return value, NewError(codes.MaxValue, validator.Value)
	}

	return value, nil
}

// IntegerMaxValue initializes IntegerMaxValueValidator instance.
func IntegerMaxValue(value int) *IntegerMaxValueValidator {
	return &IntegerMaxValueValidator{value}
}
