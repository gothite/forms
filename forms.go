package forms

import (
	"github.com/gothite/forms/fields"
)

// FormFactory implements form factory.
type FormFactory func(map[string]interface{}) *Form

// NewFormFactory initializes factory to make forms with specified fields.
func NewFormFactory(clean func(form *Form) error, fields ...fields.Field) FormFactory {
	return func(data map[string]interface{}) *Form {
		form := &Form{Data: data, Clean: clean}
		form.Fields = make(map[string]*FieldWrapper, len(fields))

		for _, field := range fields {
			form.Fields[field.GetName()] = &FieldWrapper{field, nil, nil}
		}

		return form
	}

}

// Form describes a form validator.
type Form struct {
	Fields map[string]*FieldWrapper
	Data   map[string]interface{}
	Error  error
	Clean  func(form *Form) error
}

// IsValid checks data and returns true if data is valid.
func (form *Form) IsValid() bool {
	valid := true

	for name, field := range form.Fields {
		value, ok := form.Data[name]

		if !ok {
			if field.IsRequired() {
				field.Error = field.GetError("Required")
				valid = false
				continue
			} else {
				field.Value = field.GetDefault()
				continue
			}
		}

		value, err := field.Validate(value)

		if err != nil {
			field.Error = err
			valid = false
			continue
		}

		field.Value = value
	}

	if valid && form.Clean != nil {
		err := form.Clean(form)

		if err != nil {
			form.Error = err
			valid = false
		}
	}

	return valid
}
