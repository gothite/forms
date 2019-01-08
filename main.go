package forms

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
)

type Clean func(data Data) error
type Schema []Clean

type Form interface {
	Schema() Schema
	Messages() map[uint]string
	FieldError(field string, code uint, parameters ...interface{}) error
	FormError(code uint, errors ...error) error
	Clean(data Data) error
	OnError(err *FormError) error
}

type BaseForm struct{}

func (form *BaseForm) Messages() map[uint]string {
	return Messages
}

func (form *BaseForm) FieldError(field string, code uint, parameters ...interface{}) error {
	return NewFieldError(code, field, Messages[Unknown], parameters)
}

func (form *BaseForm) FormError(code uint, errors ...error) error {
	return NewFormError(code, Messages[Invalid], errors...)
}

func (form *BaseForm) Schema() Schema {
	return nil
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
		return form.OnError(NewFormError(InvalidJSON, Messages[InvalidJSON]))
	}

	return Validate(form, JSON(data))
}

func ValidateForm(form Form, data url.Values) error {
	return Validate(form, FormData(data))
}

func ValidateMultipartForm(form Form, data *multipart.Form) error {
	return Validate(form, MultipartFormData{data})
}

func getMessage(code uint, empty string, messages ...map[uint]string) string {
	for _, messages := range messages {
		if message, ok := messages[code]; ok {
			return message
		}
	}

	return empty
}

func Validate(form Form, data Data) error {
	var errors []error
	var messages = form.Messages()

	for _, clean := range form.Schema() {
		if err := clean(data); err != nil {
			if err, ok := err.(*FieldError); ok {
				err.Message = getMessage(err.Code, err.Message, messages, Messages)
			}

			errors = append(errors, err)
		}
	}

	if errors != nil {
		return form.OnError(NewFormError(Invalid, getMessage(Invalid, Messages[Unknown], messages, Messages), errors...))
	}

	err := form.Clean(data)

	if err, ok := err.(*FormError); ok {
		err.Message = getMessage(err.Code, err.Message, messages, Messages)
	}

	if err, ok := err.(*FieldError); ok {
		err.Message = getMessage(err.Code, err.Message, messages, Messages)
	}

	return err
}
