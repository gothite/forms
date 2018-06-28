package forms

import (
	"bytes"
	"net/url"
	"testing"
)

var InvalidCredentials = NewFormError(101, "Incorrect username or password!", nil)

type LoginForm struct {
	BaseForm

	Username string
	Password string

	schema []Clean
}

func (form *LoginForm) Schema() []Clean {
	return []Clean{
		form.CleanUsername,
		form.CleanPassword,
	}
}

func (form *LoginForm) Clean(_ Data) error {
	if form.Username != form.Password {
		return InvalidCredentials
	}

	return nil
}

func (form *LoginForm) CleanUsername(data Data) error {
	if value := data.Get("username"); value == nil {
		return form.GetError("username", Required)
	} else if value, ok := value.(string); !ok {
		return form.GetError("username", Invalid)
	} else {
		form.Username = value
	}

	return nil
}

func (form *LoginForm) CleanPassword(data Data) error {
	if value := data.Get("password"); value == nil {
		return form.GetError("password", Required)
	} else if value, ok := value.(string); !ok {
		return form.GetError("password", Invalid)
	} else {
		form.Password = value
	}

	return nil
}

func TestValidateJSON(test *testing.T) {
	var form = &LoginForm{}
	var data = bytes.NewBufferString(`{"username": "bindlock", "password": "bindlock"}`)

	if err := ValidateJSON(form, data); err != nil {
		test.Fatal(err)
	} else if form.Username != "bindlock" || form.Username != form.Password {
		test.Errorf("Incorrect values!")
		test.Errorf("Expected: bindlock")
		test.Fatalf("Actual: %s", form.Username)
	}
}

func TestValidateForm(test *testing.T) {
	var form = &LoginForm{}
	var data = url.Values{"username": []string{"bindlock"}, "password": []string{"bindlock"}}

	if err := ValidateForm(form, data); err != nil {
		test.Fatal(err)
	} else if form.Username != "bindlock" || form.Username != form.Password {
		test.Errorf("Incorrect values!")
		test.Errorf("Expected: bindlock")
		test.Fatalf("Actual: %s", form.Username)
	}
}

func BenchmarkValidate(benchmark *testing.B) {
	var data = JSON(map[string]interface{}{"username": "bindlock", "password": "bindlock"})

	for i := 0; i < benchmark.N; i++ {
		Validate(&LoginForm{}, data)
	}
}
