package forms

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/gothite/forms/fields"
)

const ErrorCodeTest uint = 777

var CustomForm = NewForm(
	map[uint]error{ErrorCodeTest: TestError{}},
	&fields.Integer{Name: "id", Required: true, AllowStrings: true},
	&fields.String{Name: "Username", AllowBlank: true},
)

type CustomFormData struct {
	ID       int `forms:"id"`
	Username string
}

func (data *CustomFormData) Clean(form *Form) error {
	if data.ID == 0 {
		return form.GetError(ErrorCodeTest, nil)
	}

	return nil
}

type TestError struct{}

func (err TestError) Error() string {
	return ""
}

func TestForm(test *testing.T) {
	var form CustomFormData
	var data = map[string]interface{}{"id": 1}

	if err, errors := CustomForm.Validate(&form, data); err != nil {
		test.Errorf("Clean error: %s", err)
		test.Errorf("Fields errors: %s", errors)
		return
	}

	if form.ID != data["id"].(int) {
		test.Errorf("ID incorrect!")
		test.Errorf("Expected: %d", data["id"].(int))
		test.Errorf("Got: %d", form.ID)
	}
}

func TestFormIncorrectData(test *testing.T) {
	var form CustomFormData
	var data = map[string]interface{}{"Username": 1}

	if err, errors := CustomForm.Validate(&form, data); err == nil {
		test.Fatal("Must fail!")
	} else if _, ok := errors["id"]; !ok {
		test.Fatal("Errors must contains 'id'!")
	} else if _, ok := errors["Username"]; !ok {
		test.Fatal("Errors must contains 'Username'!")
	}
}

func TestFormCleanFail(test *testing.T) {
	var form CustomFormData
	var data = map[string]interface{}{"id": 0}

	if err, _ := CustomForm.Validate(&form, data); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestFormGetErrorUnknownCode(test *testing.T) {
	if err := CustomForm.GetError(666, nil); err == nil {
		test.Fatal("Must return error!")
	} else if err, ok := err.(*Error); !ok {
		test.Fatal("Must be Error!")
	} else {
		_ = err.Error()
	}
}

func TestFormValidateJSON(test *testing.T) {
	var form CustomFormData
	var reader = bytes.NewReader([]byte(`{"id": "1"}`))

	if err, errors := CustomForm.ValidateJSON(&form, reader); err != nil {
		test.Errorf("Clean error: %s", err)
		test.Errorf("Fields errors: %s", errors)
	} else if form.ID != 1 {
		test.Errorf("ID incorrect!")
		test.Errorf("Expected: 1")
		test.Errorf("Got: %d", form.ID)
	}
}

func TestFormValidateJSONInvalid(test *testing.T) {
	var form CustomFormData
	var reader = bytes.NewReader([]byte(`{"id=1"}`))

	if err, _ := CustomForm.ValidateJSON(&form, reader); err == nil {
		test.Fatal("Must return error!")
	}
}

func TestFormValidateForm(test *testing.T) {
	var form CustomFormData
	var payload, _ = url.ParseQuery("id=1&Username")

	if err, errors := CustomForm.ValidateForm(&form, payload); err != nil {
		test.Errorf("Clean error: %s", err)
		test.Errorf("Fields errors: %s", errors)
	} else if form.ID != 1 {
		test.Errorf("ID incorrect!")
		test.Errorf("Expected: 1")
		test.Errorf("Got: %d", form.ID)
	}
}

func BenchmarkForm(benchmark *testing.B) {
	var data = map[string]interface{}{"id": 1}

	for i := 0; i < benchmark.N; i++ {
		var form CustomFormData

		CustomForm.Validate(&form, data)
	}
}
