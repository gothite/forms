package fields

import (
	"fmt"
	"time"

	"github.com/gothite/forms/validators"
)

// DatetimeErrors is a code-error mapping for Datetime field.
var DatetimeErrors = map[string]string{
	"Required": "This field is required.",
	"Invalid":  "Ensure this value is a valid datetime string.",
}

// Datetime is integer field.
type Datetime struct {
	Name       string
	Validators []validators.Validator
	Required   bool
	Default    float64
	Errors     map[string]string // Overrides default errors
}

// IsRequired returns true if field is required.
func (field *Datetime) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Datetime) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Datetime) GetName() string {
	return field.Name
}

// GetValidators returns additional field validators.
func (field *Datetime) GetValidators() []validators.Validator {
	return field.Validators
}

// GetError returns error by code.
func (field *Datetime) GetError(code string, parameters ...interface{}) error {
	message, ok := field.Errors[code]

	if !ok {
		message = DatetimeErrors[code]
	}

	return fmt.Errorf(message, parameters...)
}

// Validate check and clean an input value.
func (field *Datetime) Validate(v interface{}) (interface{}, error) {
	value, ok := v.(string)

	if !ok {
		return nil, field.GetError("Invalid")
	}

	datetime, err := time.Parse(time.RFC3339, value)

	if err != nil {
		return nil, err
	}

	return datetime, nil
}
