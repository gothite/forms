package forms

import (
	"github.com/gothite/forms/fields"
	"github.com/gothite/forms/validators"
)

// FieldWrapper wraps Field instance to keep current form value and possible error.
type FieldWrapper struct {
	fields.Field

	Value interface{}
	Error error
}

// Validate checks an input value using Field.Validate and field validators.
func (wrapper *FieldWrapper) Validate(v interface{}) (interface{}, error) {
	value, err := wrapper.Field.Validate(v)

	if err != nil {
		return nil, err
	}

	for _, validator := range wrapper.GetValidators() {
		var err *validators.Error

		value, err = validator.Validate(v)

		if err != nil {
			return nil, wrapper.GetError(err.Code, err.Parameters...)
		}
	}

	return value, nil
}
