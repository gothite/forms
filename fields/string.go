package fields

import (
	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// StringErrors is a code-error mapping for String field.
var StringErrors = map[uint]string{
	codes.Required:  "This field is required.",
	codes.Invalid:   "Ensure this value is valid string.",
	codes.Blank:     "Blank strings aren't allowed.",
	codes.MinLength: "Ensure this value has at least %d characters.",
	codes.MaxLength: "Ensure this value has at most %d characters.",
}

// String is boolean field.
type String struct {
	Name       string
	Validators []validators.StringValidator
	Required   bool
	Default    string
	Errors     map[uint]string
	ErrorFunc  ErrorFunc

	AllowBlank bool
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

// GetError returns error by code.
func (field *String) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(code, value, field.Errors, StringErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *String) Validate(v interface{}) (interface{}, error) {
	value, ok := v.(string)

	if !ok {
		return nil, field.GetError(codes.Invalid, v)
	}

	if len(value) == 0 && !field.AllowBlank {
		return nil, field.GetError(codes.Blank, v)
	}

	for _, validator := range field.Validators {
		var err *validators.Error

		value, err = validator.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, v, err.Parameters...)
		}
	}

	return value, nil
}
