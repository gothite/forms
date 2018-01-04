package validators

type StringValidator interface {
	Validate(value string) (string, *Error)
}

// StringMinLengthValidator controls minimum length.
type StringMinLengthValidator struct {
	Length int
}

// Validate do validation.
func (validator StringMinLengthValidator) Validate(value string) (string, *Error) {
	if len(value) < validator.Length {
		return value, NewError("MinLength", validator.Length)
	}

	return value, nil
}

// StringMinLength initializes MinLengthValidator instance.
func StringMinLength(length int) *StringMinLengthValidator {
	return &StringMinLengthValidator{length}
}

// StringMaxLengthValidator controls minimum length.
type StringMaxLengthValidator struct {
	Length int
}

// Validate do validation.
func (validator StringMaxLengthValidator) Validate(value string) (string, *Error) {
	if len(value) > validator.Length {
		return value, NewError("MaxLength", validator.Length)
	}

	return value, nil
}

// StringMaxLength initializes MaxLengthValidator instance.
func StringMaxLength(length int) *StringMaxLengthValidator {
	return &StringMaxLengthValidator{length}
}
