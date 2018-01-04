package forms

import (
	"errors"
	"testing"

	"github.com/gothite/forms/fields"
)

type CustomFormData struct {
	ID       int `forms:"id"`
	Username string
}

func (form *CustomFormData) Clean() error {
	if form.ID == 0 {
		return errors.New("ID == 0!")
	}

	return nil
}

var CustomForm = NewForm(
	&fields.Integer{Name: "id", Required: true},
	&fields.String{Name: "Username"},
)

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

func BenchmarkForm(benchmark *testing.B) {
	var data = map[string]interface{}{"id": 1}

	for i := 0; i < benchmark.N; i++ {
		var form CustomFormData

		CustomForm.Validate(&form, data)
	}
}
