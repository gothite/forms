package validators

// MinLengthValidator controls minimum length.
type MinLengthValidator struct {
	Length int
}

// Validate do validation.
func (validator MinLengthValidator) Validate(v interface{}) (interface{}, *Error) {
	value, ok := v.(string)

	if !ok {
		return nil, NewError("Invalid")
	}

	if len(value) < validator.Length {
		return nil, NewError("MinLength", validator.Length)
	}

	return value, nil
}

// MinLength initializes MinLengthValidator instance.
func MinLength(length int) *MinLengthValidator {
	return &MinLengthValidator{length}
}

// MaxLengthValidator controls minimum length.
type MaxLengthValidator struct {
	Length int
}

// Validate do validation.
func (validator MaxLengthValidator) Validate(v interface{}) (interface{}, *Error) {
	value, ok := v.(string)

	if !ok {
		return nil, NewError("Invalid")
	}

	if len(value) > validator.Length {
		return nil, NewError("MaxLength", validator.Length)
	}

	return value, nil
}

// MaxLength initializes MaxLengthValidator instance.
func MaxLength(length int) *MaxLengthValidator {
	return &MaxLengthValidator{length}
}

// MinValueValidator controls minimum value.
type MinValueValidator struct {
	Value float64
}

// Validate do validation.
func (validator MinValueValidator) Validate(v interface{}) (interface{}, *Error) {
	var value float64

	if float, ok := v.(float64); ok {
		value = float
	} else if integer, ok := v.(int); ok {
		value = float64(integer)
	} else {
		return nil, NewError("Invalid")
	}

	if value < validator.Value {
		return nil, NewError("MinValue", validator.Value)
	}

	return value, nil
}

// MinValue initializes MinValueValidator instance.
func MinValue(value float64) *MinValueValidator {
	return &MinValueValidator{value}
}

// MaxValueValidator controls maximum value.
type MaxValueValidator struct {
	Value float64
}

// Validate do validation.
func (validator MaxValueValidator) Validate(v interface{}) (interface{}, *Error) {
	var value float64

	if float, ok := v.(float64); ok {
		value = float
	} else if integer, ok := v.(int); ok {
		value = float64(integer)
	} else {
		return nil, NewError("Invalid")
	}

	if value > validator.Value {
		return nil, NewError("MaxValue", validator.Value)
	}

	return value, nil
}

// MaxValue initializes MaxValueValidator instance.
func MaxValue(value float64) *MaxValueValidator {
	return &MaxValueValidator{value}
}
