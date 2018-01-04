package fields

import (
	"regexp"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// https://github.com/asaskevich/govalidator
var emailRE = regexp.MustCompile(
	"^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$",
)

// EmailErrors is a code-error mapping for Email field.
var EmailErrors = map[uint]string{
	codes.Required: "This field is required.",
	codes.Invalid:  "Ensure this value is valid email string.",
}

// Email is integer field.
type Email struct {
	Name       string
	Validators []validators.StringValidator
	Required   bool
	Default    string
	Errors     map[uint]string
	ErrorFunc  ErrorFunc
}

// IsRequired returns true if field is required.
func (field *Email) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Email) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Email) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Email) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(code, value, field.Errors, EmailErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Email) Validate(v interface{}) (interface{}, error) {
	value, ok := v.(string)

	if !ok {
		return nil, field.GetError(codes.Invalid, v)
	}

	if !emailRE.MatchString(value) {
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
