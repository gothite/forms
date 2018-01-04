package fields

import (
	"fmt"
	"strconv"

	"github.com/gothite/forms/validators"
)

// IntegerErrors is a code-error mapping for Integer field.
var IntegerErrors = map[string]string{
	"Required": "This field is required.",
	"Invalid":  "Ensure this value is valid integer.",
	"MinValue": "Ensure this value is greater than or equal to %v.",
	"MaxValue": "Ensure this value is less than or equal to %v.",
}

// Integer is integer field.
type Integer struct {
	Name       string
	Validators []validators.IntegerValidator
	Required   bool
	Default    int
	Errors     map[string]string // Overrides default errors

	AllowStrings bool
}

// IsRequired returns true if field is required.
func (field *Integer) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Integer) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Integer) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Integer) GetError(code string, parameters ...interface{}) error {
	message, ok := field.Errors[code]

	if !ok {
		message = IntegerErrors[code]
	}

	return fmt.Errorf(message, parameters...)
}

// Validate check and clean an input value.
func (field *Integer) Validate(v interface{}) (interface{}, error) {
	if field.AllowStrings {
		if value, ok := v.(string); ok {
			value, err := strconv.Atoi(value)

			if err != nil {
				return nil, field.GetError("Invalid")
			}

			return value, nil
		}
	}

	value, ok := v.(int)

	if !ok {
		return nil, field.GetError("Invalid")
	}

	return value, nil
}
