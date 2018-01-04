package forms

import (
	"reflect"

	"github.com/gothite/forms/fields"
)

// NewForm returns Form instance with passed fields.
func NewForm(fields ...fields.Field) *Form {
	return &Form{Fields: fields}
}

// Form describes a form validator.
type Form struct {
	Fields []fields.Field
}

// Validate validates input data and map it to target.
func (form *Form) Validate(target interface{}, data map[string]interface{}) (bool, map[string]error) {
	var errors = make(map[string]error)

	for _, field := range form.Fields {
		value, ok := data[field.GetName()]

		if !ok {
			if field.IsRequired() {
				errors[field.GetName()] = field.GetError("Required")
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
		return true, errors
	}

	return false, errors
}

func Map(target interface{}, data map[string]interface{}) {
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
