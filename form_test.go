package forms

import (
	"testing"

	"github.com/gothite/forms/fields"
)

func TestForm(test *testing.T) {
	var form struct {
		ID       int `forms:"id"`
		Username string
	}
	var CustomForm = NewForm(
		&fields.Integer{Name: "id", Required: true},
		&fields.String{Name: "Username"},
	)
	var data = map[string]interface{}{"id": 1}

	if valid, errors := CustomForm.Validate(&form, data); !valid {
		test.Fatalf("Errors: %s", errors)
	}

	if form.ID != data["id"].(int) {
		test.Errorf("ID incorrect!")
		test.Errorf("Expected: %d", data["id"].(int))
		test.Errorf("Got: %d", form.ID)
	}
}

func TestFormIncorrectData(test *testing.T) {
	var form struct {
		ID       int `forms:"id"`
		Username string
	}
	var CustomForm = NewForm(
		&fields.Integer{Name: "id", Required: true},
		&fields.String{Name: "Username"},
	)
	var data = map[string]interface{}{"Username": 1}

	if valid, errors := CustomForm.Validate(&form, data); valid {
		test.Fatal("Must fail!")
	} else if _, ok := errors["id"]; !ok {
		test.Fatal("Errors must contains 'id'!")
	} else if _, ok := errors["Username"]; !ok {
		test.Fatal("Errors must contains 'Username'!")
	}
}

func BenchmarkForm(benchmark *testing.B) {
	type CustomFormData struct {
		ID       int    `forms:"id"`
		Username string `forms:"username"`
	}

	var CustomForm = NewForm(
		&fields.Integer{Name: "id", Required: true},
		&fields.String{Name: "username"},
	)
	var data = map[string]interface{}{"id": 1}

	for i := 0; i < benchmark.N; i++ {
		var form CustomFormData

		CustomForm.Validate(&form, data)
	}
}
