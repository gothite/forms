package fields

import (
	"fmt"
	"strings"

	"github.com/gothite/forms/validators"
)

// BooleanErrors is a code-error mapping for Boolean field.
var BooleanErrors = map[string]string{
	"Required": "This field is required.",
	"Invalid":  "Ensure this value is a valid boolean.",
}

// Boolean is boolean field.
type Boolean struct {
	Name       string
	Validators []validators.BooleanValidator
	Required   bool
	Default    bool
	Errors     map[string]string

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
func (field *Boolean) GetError(code string, parameters ...interface{}) error {
	message, ok := field.Errors[code]

	if !ok {
		message = IntegerErrors[code]
	}

	return fmt.Errorf(message, parameters...)
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
				return nil, field.GetError("Invalid")
			}
		}
	case int:
		if field.AllowNumbers {
			if v == 1 {
				value = true
			} else if v == 0 {
				value = false
			} else {
				return nil, field.GetError("Invalid")
			}
		}
	default:
		return nil, field.GetError("Invalid")
	}

	for _, validator := range field.Validators {
		var err *validators.Error

		value, err = validator.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, err.Parameters...)
		}
	}

	return value, nil
}
