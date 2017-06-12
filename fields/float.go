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
	Validators []validators.Validator
	Required   bool
	Default    float64
	Errors     map[string]string // Overrides default errors

	AllowStrings bool
	MinValue     *validators.MinValueValidator
	MaxValue     *validators.MaxValueValidator
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

// GetValidators returns additional field validators.
func (field *Float) GetValidators() []validators.Validator {
	return field.Validators
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
	if field.AllowStrings {
		if value, ok := v.(string); ok {
			value, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return nil, field.GetError("Invalid")
			}

			return value, nil
		}
	}

	value, ok := v.(float64)

	if !ok {
		return nil, field.GetError("Invalid")
	}

	if field.MinValue != nil {
		_, err := field.MinValue.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, err.Parameters...)
		}
	}

	if field.MaxValue != nil {
		_, err := field.MaxValue.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, err.Parameters...)
		}
	}

	return value, nil
}
