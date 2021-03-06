package forms

import "fmt"

type FieldError struct {
	Code       uint          `json:"code"`
	Field      string        `json:"field"`
	Message    string        `json:"message"`
	Parameters []interface{} `json:"parameters"`
}

func NewFieldError(code uint, field, message string, parameters []interface{}) *FieldError {
	return &FieldError{code, field, message, parameters}
}

func (err *FieldError) Format() string {
	return fmt.Sprintf(err.Message, err.Parameters...)
}

func (err *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", err.Field, err.Format())
}

type FormError struct {
	Code    uint    `json:"code"`
	Message string  `json:"message"`
	Errors  []error `json:"errors"`
}

func NewFormError(code uint, message string, errors ...error) *FormError {
	return &FormError{code, message, errors}
}

func (err *FormError) Error() string {
	return fmt.Sprintf("%s (code: %d)", err.Message, err.Code)
}

const (
	Unknown uint = iota
	InvalidJSON
	Required
	Invalid
)

var Messages = map[uint]string{
	Unknown:     "Unknown error.",
	Invalid:     "Ensure that all values are valid.",
	InvalidJSON: "Unable to parse JSON.",
	Required:    "Value is required.",
}
