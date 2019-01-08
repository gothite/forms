package main

import (
	"log"
	"net/url"

	"github.com/govenant/forms"
)

var InvalidCredentials = forms.NewFormError(101, "Incorrect username or password!", nil)

type LoginForm struct {
	forms.BaseForm

	Username string
	Password string
}

func (form *LoginForm) Schema() forms.Schema {
	return forms.Schema{
		form.CleanUsername,
		form.CleanPassword,
	}
}

func (form *LoginForm) Clean(_ forms.Data) error {
	if form.Username != form.Password {
		return InvalidCredentials
	}

	return nil
}

func (form *LoginForm) CleanUsername(data forms.Data) error {
	if value := data.Get("username"); value == nil {
		return form.FieldError("username", forms.Required)
	} else if value, ok := value.(string); !ok {
		return form.FieldError("username", forms.Invalid)
	} else {
		form.Username = value
	}

	return nil
}

func (form *LoginForm) CleanPassword(data forms.Data) error {
	if value := data.Get("password"); value == nil {
		return form.FieldError("password", forms.Required)
	} else if value, ok := value.(string); !ok {
		return form.FieldError("password", forms.Invalid)
	} else {
		form.Password = value
	}

	return nil
}

func main() {
	var form = &LoginForm{}
	var data = url.Values{"username": {"govenant"}, "password": {"covenant"}}

	if err := forms.ValidateForm(form, data); err != nil {
		log.Fatal(err.(*forms.FormError))
	}
}
