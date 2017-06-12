package fields

import (
	"fmt"

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
	Validators []validators.Validator
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

// GetValidators returns additional field validators.
func (field *Boolean) GetValidators() []validators.Validator {
	return field.Validators
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
	if field.AllowStrings {
		if value, ok := v.(string); ok {
			if value == "t" || value == "true" {
				return true, nil
			} else if value == "f" || value == "false" {
				return false, nil
			}
		}
	}

	if field.AllowNumbers {
		if value, ok := v.(int); ok {
			if value == 1 {
				return true, nil
			} else if value == 0 {
				return false, nil
			}
		}
	}

	value, ok := v.(bool)

	if !ok {
		return nil, field.GetError("Invalid")
	}

	return value, nil
}
