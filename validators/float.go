package validators

import "github.com/gothite/forms/codes"

type FloatValidator interface {
	Validate(value float64) (float64, *Error)
}

// FloatMinValueValidator controls minimum value.
type FloatMinValueValidator struct {
	Value float64
}

// Validate do validation.
func (validator *FloatMinValueValidator) Validate(value float64) (float64, *Error) {
	if value < validator.Value {
		return value, NewError(codes.MinValue, validator.Value)
	}

	return value, nil
}

// FloatMinValue initializes FloatMinValueValidator instance.
func FloatMinValue(value float64) *FloatMinValueValidator {
	return &FloatMinValueValidator{value}
}

// FloatMaxValueValidator controls maximum value.
type FloatMaxValueValidator struct {
	Value float64
}

// Validate do validation.
func (validator FloatMaxValueValidator) Validate(value float64) (float64, *Error) {
	if value > validator.Value {
		return value, NewError(codes.MaxValue, validator.Value)
	}

	return value, nil
}

// FloatMaxValue initializes FloatMaxValueValidator instance.
func FloatMaxValue(value float64) *FloatMaxValueValidator {
	return &FloatMaxValueValidator{value}
}
