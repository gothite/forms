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
		Map(target, data, []int{}, target)

		if err := target.Clean(form); err != nil {
			return err, errors
		}

		return nil, errors
	}

	return form.GetError(codes.Invalid, errors), errors
}

func Map(target interface{}, data map[string]interface{}, index []int, subtarget interface{}) {
	var targetValue = reflect.Indirect(reflect.ValueOf(target))
	var targetType = targetValue.Type()
	var numFields = reflect.Indirect(reflect.ValueOf(subtarget)).NumField()

	index = append(index, 0)

	for i := 0; i < numFields; i++ {
		index[len(index)-1] = i

		var fieldType = targetType.FieldByIndex(index)
		var fieldValue = targetValue.FieldByIndex(index)
		var name = fieldType.Tag.Get("forms")

		if name == "" {
			name = fieldType.Name
		}

		if value, ok := data[name]; ok && value != nil {
			var reflectValue = reflect.ValueOf(value)

			if reflectValue.Kind() == reflect.Slice {
				var length = reflectValue.Len()
				var elem = fieldValue.Addr().Elem()

				elem.Set(reflect.MakeSlice(fieldType.Type, length, length))

				for n := 0; n < length; n++ {
					set(fieldValue.Index(n), reflectValue.Index(n))
				}

			} else if reflectValue.Kind() == reflect.Map {
				Map(target, value.(map[string]interface{}), index, fieldValue.Interface())
			} else {
				fieldValue.Set(reflectValue)
			}
		}
	}
}

func set(target reflect.Value, value reflect.Value) {
	switch target.Kind() {
	case reflect.Bool:
		target.SetBool(value.Interface().(bool))
	case reflect.String:
		target.SetString(value.Interface().(string))
	case reflect.Float32, reflect.Float64:
		target.SetFloat(value.Interface().(float64))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		target.SetInt(int64(value.Interface().(int)))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		target.SetUint(uint64(value.Interface().(int)))
	}
}
