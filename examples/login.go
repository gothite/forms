package main

import (
	"log"
	"net/url"

	"github.com/gothite/forms"
)

var InvalidCredentials = forms.NewFormError(101, "Incorrect username or password!", nil)

type LoginForm struct {
	forms.BaseForm

	Username string
	Password string
}

func (form *LoginForm) Schema() []forms.Clean {
	return []forms.Clean{
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
		return form.GetError("username", forms.Required)
	} else if value, ok := value.(string); !ok {
		return form.GetError("username", forms.Invalid)
	} else {
		form.Username = value
	}

	return nil
}

func (form *LoginForm) CleanPassword(data forms.Data) error {
	if value := data.Get("password"); value == nil {
		return form.GetError("password", forms.Required)
	} else if value, ok := value.(string); !ok {
		return form.GetError("password", forms.Invalid)
	} else {
		form.Password = value
	}

	return nil
}

func main() {
	var form = &LoginForm{}
	var data = url.Values{"username": []string{"bindlock"}, "password": []string{"bindlock"}}

	if err := forms.ValidateForm(form, data); err != nil {
		log.Fatal(err.(*forms.FormError))
	}
}
