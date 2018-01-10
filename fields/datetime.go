package fields

import (
	"time"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// DatetimeErrors is a code-error mapping for Datetime field.
var DatetimeErrors = map[uint]string{
	codes.Required: "This field is required.",
	codes.Invalid:  "Ensure this value is a valid datetime string.",
}

// Datetime is integer field.
type Datetime struct {
	Name       string
	Validators []validators.DatetimeValidator
	Required   bool
	Default    time.Time
	Errors     map[uint]string
	ErrorFunc  ErrorFunc
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

// GetError returns error by code.
func (field *Datetime) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(field, code, value, field.Errors, DatetimeErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Datetime) Validate(v interface{}) (interface{}, error) {
	var value time.Time

	if result, ok := v.(string); !ok {
		return nil, field.GetError(codes.Invalid, v)
	} else {
		var err error

		value, err = time.Parse(time.RFC3339Nano, result)

		if err != nil {
			return nil, field.GetError(codes.Invalid, v)
		}
	}

	for _, validator := range field.Validators {
		v, err := validator.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, value, err.Parameters...)
		}

		value = v
	}

	return value, nil
}
