package forms

import (
	"encoding/json"
	"io"
	"net/url"
	"reflect"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/fields"
)

type FormData interface {
	Clean(form *Form) error
}

// Form describes a form validator.
type Form struct {
	Fields []fields.Field
	Errors map[uint]error
}

// NewForm returns Form instance with passed fields.
func NewForm(errors map[uint]error, fields ...fields.Field) *Form {
	return &Form{Errors: errors, Fields: fields}
}

func (form *Form) GetError(code uint, errors map[string]error) error {
	if err, ok := form.Errors[code]; ok {
		return err
	} else if message, ok := Errors[code]; ok {
		return &Error{
			Code:    code,
			Message: message,
			Errors:  errors,
		}
	} else {
		return &Error{
			Code:    codes.Unknown,
			Message: Errors[codes.Unknown],
			Errors:  errors,
		}
	}
}

func (form *Form) ValidateJSON(target FormData, reader io.Reader) (error, map[string]error) {
	var data = make(map[string]interface{}, len(form.Fields))

	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return form.GetError(codes.InvalidJSON, nil), nil
	}

	return form.Validate(target, data)
}

func (form *Form) ValidateForm(target FormData, payload url.Values) (error, map[string]error) {
	var data = make(map[string]interface{}, len(form.Fields))

	for key, values := range payload {
		data[key] = values[0]
	}

	return form.Validate(target, data)
}

// Validate validates input data and map it to target.
func (form *Form) Validate(target FormData, data map[string]interface{}) (error, map[string]error) {
	var errors = make(map[string]error, len(form.Fields))

	for _, field := range form.Fields {
		value, ok := data[field.GetName()]

		if !ok {
			if field.IsRequired() {
				errors[field.GetName()] = field.GetError(codes.Required, nil)
			} else {
				data[field.GetName()] = field.GetDefault()
			}

			continue
		}

		value, err := field.Validate(value)

		if err != nil {
			errors[field.GetName()] = err
			continue
		}

		data[field.GetName()] = value
	}

	if len(errors) == 0 {
		Map(target, data)

		if err := target.Clean(form); err != nil {
			return err, errors
		}

		return nil, errors
	}

	return form.GetError(codes.Invalid, errors), errors
}

func Map(target FormData, data map[string]interface{}) {
	var targetValue = reflect.Indirect(reflect.ValueOf(target))
	var targetType = targetValue.Type()
	var numFields = targetValue.NumField()

	for i := 0; i < numFields; i++ {
		var fieldType = targetType.Field(i)
		var name = fieldType.Tag.Get("forms")

		if name == "" {
			name = fieldType.Name
		}

		if value, ok := data[name]; ok {
			targetValue.Field(i).Set(reflect.ValueOf(value))
		}
	}
}
