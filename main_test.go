package forms

import (
	"bytes"
	"mime/multipart"
	"net/url"
	"testing"
)

var IncorrectCredentials uint = 101
var messages = map[uint]string{IncorrectCredentials: "Incorrect username or password!"}

type LoginForm struct {
	BaseForm

	Username string
	Password string
}

func (form *LoginForm) Schema() Schema {
	return Schema{
		form.CleanUsername,
		form.CleanPassword,
	}
}

func (form *LoginForm) Messages() map[uint]string {
	return messages
}

func (form *LoginForm) Clean(_ Data) error {
	if form.Username != form.Password {
		return form.FormError(101)
	}

	return nil
}

func (form *LoginForm) CleanUsername(data Data) error {
	if value := data.Get("username"); value == nil {
		return form.FieldError("username", IncorrectCredentials)
	} else if value, ok := value.(string); !ok {
		return form.FieldError("username", Invalid)
	} else {
		form.Username = value
	}

	return nil
}

func (form *LoginForm) CleanPassword(data Data) error {
	if value := data.Get("password"); value == nil {
		return form.FieldError("password", Unknown)
	} else if value, ok := value.(string); !ok {
		return form.FieldError("password", Invalid)
	} else {
		form.Password = value
	}

	return nil
}

func TestValidateJSON(test *testing.T) {
	var form = &LoginForm{}
	var data = bytes.NewBufferString(`{"username": "bindlock",`)

	if err := ValidateJSON(form, data); err == nil {
		test.Fatal("Expected error!")
	} else if err := err.(*FormError); err.Code != InvalidJSON {
		test.Errorf("Incorrect code!")
		test.Errorf("Expected: %d", InvalidJSON)
		test.Fatalf("Actual: %d", err.Code)
	}

	data = bytes.NewBufferString(`{"username": "bindlock", "password": "bindlock"}`)

	if err := ValidateJSON(form, data); err != nil {
		test.Fatal(err)
	} else if form.Username != "bindlock" || form.Username != form.Password {
		test.Errorf("Incorrect values!")
		test.Errorf("Expected: bindlock")
		test.Fatalf("Actual: %s", form.Username)
	}

	data = bytes.NewBufferString(`{"username": "bindlock"}`)

	if err := ValidateJSON(form, data); err == nil {
		test.Fatal("Expected error!")
	} else if err := err.(*FormError); err.Code != Invalid {
		test.Errorf("Incorrect code!")
		test.Errorf("Expected: %d", Invalid)
		test.Fatalf("Actual: %d", err.Code)
	}
}

func TestValidateForm(test *testing.T) {
	var form = &LoginForm{}
	var data = url.Values{"username": {"bindlock"}, "password": {"bindlock"}}

	if err := ValidateForm(form, data); err != nil {
		test.Fatal(err)
	} else if form.Username != "bindlock" || form.Username != form.Password {
		test.Errorf("Incorrect values!")
		test.Errorf("Expected: bindlock")
		test.Fatalf("Actual: %s", form.Username)
	}

	data = url.Values{"username": {"bindlock"}, "password": {"lock"}}

	if err := ValidateForm(form, data); err == nil {
		test.Fatal("Expected error!")
	} else if message := err.(*FormError).Message; message != messages[101] {
		test.Errorf("Incorrect error message!")
		test.Errorf("Expected: %s", messages[101])
		test.Fatalf("Actual: %s", message)
	}

	data = url.Values{}

	if err := ValidateForm(form, data); err == nil {
		test.Fatal("Expected error!")
	} else if message := err.(*FormError).Message; message != Messages[Invalid] {
		test.Errorf("Incorrect error message!")
		test.Errorf("Expected: %s", Messages[Invalid])
		test.Fatalf("Actual: %s", message)
	}
}

func TestValidateMultipartForm(test *testing.T) {
	var form = &LoginForm{}
	var data = &multipart.Form{Value: map[string][]string{"username": {"bindlock"}, "password": {"bindlock"}}}

	if err := ValidateMultipartForm(form, data); err != nil {
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
