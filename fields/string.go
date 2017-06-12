package fields

import (
	"fmt"

	"github.com/gothite/forms/validators"
)

// StringErrors is a code-error mapping for String field.
var StringErrors = map[string]string{
	"Required":  "This field is required.",
	"Invalid":   "Ensure this value is valid string.",
	"Blank":     "Blank strings aren't allowed.",
	"MinLength": "Ensure this value has at least %v characters.",
	"MaxLength": "Ensure this value has at most %v characters.",
}

// String is boolean field.
type String struct {
	Name       string
	Validators []validators.Validator
	Required   bool
	Default    string
	Errors     map[string]string

	AllowBlank bool
	MinLength  *validators.MinLengthValidator
	MaxLength  *validators.MaxLengthValidator
}

// IsRequired returns true if field is required.
func (field *String) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *String) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *String) GetName() string {
	return field.Name
}

// GetValidators returns additional field validators.
func (field *String) GetValidators() []validators.Validator {
	return field.Validators
}

// GetError returns error by code.
func (field *String) GetError(code string, parameters ...interface{}) error {
	message, ok := field.Errors[code]

	if !ok {
		message = StringErrors[code]
	}

	return fmt.Errorf(message, parameters...)
}

// Validate check and clean an input value.
func (field *String) Validate(v interface{}) (interface{}, error) {
	value, ok := v.(string)

	if !ok {
		return nil, field.GetError("Invalid")
	}

	if len(value) == 0 && !field.AllowBlank {
		return nil, field.GetError("Blank")
	}

	if field.MinLength != nil {
		_, err := field.MinLength.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, err.Parameters...)
		}
	}

	if field.MaxLength != nil {
		_, err := field.MaxLength.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, err.Parameters...)
		}
	}

	return value, nil
}
