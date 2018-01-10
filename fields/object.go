package fields

import (
	"reflect"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// ObjectErrors is a code-error mapping for Object field.
var ObjectErrors = map[uint]string{
	codes.Unknown:       "Unknown error.",
	codes.Required:      "This field is required.",
	codes.Invalid:       "Ensure this value is a valid object.",
	codes.IncorrectItem: "%s: %s",
	codes.Blank:         "Empty objects aren't allowed.",
}

type Object struct {
	Name       string
	Validators []validators.Validator
	Required   bool
	Default    interface{}
	Errors     map[uint]string
	ErrorFunc  ErrorFunc

	Fields     []Field
	AllowEmpty bool
}

func (field *Object) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Object) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Object) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Object) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(field, code, value, field.Errors, ObjectErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Object) Validate(value interface{}) (interface{}, error) {
	var reflectValue = reflect.ValueOf(value)

	if reflectValue.Kind() != reflect.Map {
		return nil, field.GetError(codes.Invalid, value)
	}

	var length = reflectValue.Len()

	if length == 0 && !field.AllowEmpty {
		return nil, field.GetError(codes.Blank, value)
	}

	if len(field.Fields) != 0 {
		var result map[string]interface{}
		var resultValue reflect.Value
		var ok bool

		if result, ok = value.(map[string]interface{}); !ok {
			result = make(map[string]interface{}, length)
			resultValue = reflect.ValueOf(result)
		} else {
			resultValue = reflectValue
		}

		for _, subfield := range field.Fields {
			var key = reflect.ValueOf(subfield.GetName())
			var item interface{}

			if v := reflectValue.MapIndex(key); !v.IsValid() {
				if subfield.IsRequired() {
					return nil, field.GetError(codes.IncorrectItem, value, subfield.GetName())
				}

				item = subfield.GetDefault()
			} else {
				var err error

				item, err = subfield.Validate(v.Interface())

				if err != nil {
					return nil, field.GetError(codes.IncorrectItem, value, subfield.GetName())
				}
			}

			resultValue.SetMapIndex(key, reflect.ValueOf(item))
		}

		value = result
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
