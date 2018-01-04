package fields

import (
	"fmt"
	"strconv"

	"github.com/gothite/forms/validators"
)

// FloatErrors is a code-error mapping for Float field.
var FloatErrors = map[string]string{
	"Required": "This field is required.",
	"Invalid":  "Ensure this value is a valid float.",
}

// Float is float field.
type Float struct {
	Name       string
	Validators []validators.FloatValidator
	Required   bool
	Default    float64
	Errors     map[string]string // Overrides default errors

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
func (field *Float) GetError(code string, parameters ...interface{}) error {
	message, ok := field.Errors[code]

	if !ok {
		message = FloatErrors[code]
	}

	return fmt.Errorf(message, parameters...)
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
