package fields

import (
	"strings"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// BooleanErrors is a code-error mapping for Boolean field.
var BooleanErrors = map[uint]string{
	codes.Unknown:  "Unknown error.",
	codes.Required: "This field is required.",
	codes.Invalid:  "Ensure this value is a valid boolean.",
}

// Boolean is boolean field.
type Boolean struct {
	Name       string
	Validators []validators.BooleanValidator
	Required   bool
	Default    bool
	Errors     map[uint]string
	ErrorFunc  ErrorFunc

	AllowStrings bool // Allow pass strings "t", "true", "f", "false" as valid boolean.
	AllowNumbers bool // Allow pass numbers 0 or 1 as valid boolean.
}

// IsRequired returns true if field is required.
func (field *Boolean) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Boolean) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Boolean) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Boolean) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(code, value, field.Errors, BooleanErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Boolean) Validate(v interface{}) (interface{}, error) {
	var value bool

	switch v := v.(type) {
	case bool:
		value = v
	case string:
		if field.AllowStrings {
			v = strings.ToLower(v)

			if v == "t" || v == "true" {
				value = true
			} else if v == "f" || v == "false" {
				value = false
			} else {
				return nil, field.GetError(codes.Invalid, v)
			}
		}
	case int:
		if field.AllowNumbers {
			if v == 1 {
				value = true
			} else if v == 0 {
				value = false
			} else {
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
