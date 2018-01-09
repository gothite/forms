package fields

import (
	"reflect"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// ArrayErrors is a code-error mapping for Array field.
var ArrayErrors = map[uint]string{
	codes.Unknown:       "Unknown error.",
	codes.Required:      "This field is required.",
	codes.Invalid:       "Ensure this value is a valid array.",
	codes.MinLength:     "Ensure this array has at least %d items.",
	codes.MaxLength:     "Ensure this array has at most %d items.",
	codes.IncorrectItem: "%d: %s",
	codes.Blank:         "Empty arrays aren't allowed.",
}

type Array struct {
	Name       string
	Validators []validators.Validator
	Required   bool
	Default    interface{}
	Errors     map[uint]string
	ErrorFunc  ErrorFunc

	Field      Field
	AllowEmpty bool
}

func (field *Array) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Array) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Array) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Array) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(field, code, value, field.Errors, ArrayErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Array) Validate(value interface{}) (interface{}, error) {
	var reflectValue = reflect.ValueOf(value)

	if reflectValue.Kind() != reflect.Slice {
		return nil, field.GetError(codes.Invalid, value)
	}

	var length = reflectValue.Len()

	if length == 0 && !field.AllowEmpty {
		return nil, field.GetError(codes.Blank, value)
	}

	if field.Field != nil {
		for i := 0; i < length; i++ {
			var indexValue = reflectValue.Index(i)

			v, err := field.Field.Validate(indexValue.Interface())

			if err != nil {
				return nil, field.GetError(codes.IncorrectItem, v, i)
			}

			indexValue.Set(reflect.ValueOf(v))
		}
	}

	for _, validator := range field.Validators {
		var err *validators.Error

		value, err = validator.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, value, err.Parameters...)
		}
	}

	return value, nil
}
