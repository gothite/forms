package fields

import (
	"strconv"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// FloatErrors is a code-error mapping for Float field.
var FloatErrors = map[uint]string{
	codes.Required: "This field is required.",
	codes.Invalid:  "Ensure this value is a valid float.",
	codes.MinValue: "Ensure this value is greater than or equal to %f.",
	codes.MaxValue: "Ensure this value is less than or equal to %f.",
}

// Float is float field.
type Float struct {
	Name       string
	Validators []validators.FloatValidator
	Required   bool
	Default    float64
	Errors     map[uint]string
	ErrorFunc  ErrorFunc

	AllowStrings bool
}

// IsRequired returns true if field is required.
func (field *Float) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Float) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Float) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Float) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(field, code, value, field.Errors, FloatErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Float) Validate(v interface{}) (interface{}, error) {
	var value float64

	switch v := v.(type) {
	case float64:
		value = v
	case string:
		if field.AllowStrings {
			var err error

			value, err = strconv.ParseFloat(v, 64)

			if err != nil {
				return nil, field.GetError(codes.Invalid, v)
			}
		}
	default:
		return nil, field.GetError(codes.Invalid, v)
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
