package forms

import (
	"encoding/json"
	"io"
	"net/url"
)

type Clean func(data Data) error

type Form interface {
	Schema() []Clean
	GetError(field string, code uint, parameters ...interface{}) error
	Clean(data Data) error
	OnError(err *FormError) error
}

type BaseForm struct{}

func (form *BaseForm) GetError(field string, code uint, parameters ...interface{}) error {
	var message string
	var ok bool

	if message, ok = Errors[code]; !ok {
		message = Errors[Unknown]
	}

	return NewFieldError(code, field, message, parameters)
}

func (form *BaseForm) Clean(_ Data) error {
	return nil
}

func (form *BaseForm) OnError(err *FormError) error {
	return err
}

func ValidateJSON(form Form, reader io.Reader) error {
	var data map[string]interface{}

	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return form.OnError(NewFormError(InvalidJSON, Errors[InvalidJSON]))
	}

	return Validate(form, JSON(data))
}

func ValidateForm(form Form, data url.Values) error {
	return Validate(form, FormData(data))
}

func Validate(form Form, data Data) error {
	var errors []error

	for _, clean := range form.Schema() {
		if err := clean(data); err != nil {
			errors = append(errors, err)
		}
	}

	if errors != nil {
		return form.OnError(NewFormError(Invalid, Errors[InvalidJSON], errors...))
	}

	return form.Clean(data)
}
