package forms

import (
	"fmt"
	"testing"

	"github.com/gothite/forms/fields"
)

func TestForm(test *testing.T) {
	factory := NewFormFactory(
		func(form *Form) error {
			if form.Data["id"].(int) == 0 {
				return fmt.Errorf("")
			}

			return nil

		},
		&fields.Integer{Name: "id", Required: true},
		&fields.String{Name: "username"},
	)
	form := factory(map[string]interface{}{"id": 1})

	if !form.IsValid() {
		test.Errorf("Form must be valid. Form error: %v", form.Error)

		for name, field := range form.Fields {
			test.Errorf("%v error: %v", name, field.Error)
		}
	}

	// Check required fields
	form = factory(map[string]interface{}{})

	if form.IsValid() {
		test.Errorf("Form must be invalid")
	}

	// Check invalid value for field
	form = factory(map[string]interface{}{"id": "hex"})

	if form.IsValid() {
		test.Errorf("Form must be invalid")
	}

	// Check user-defined Form.Clean method.
	form = factory(map[string]interface{}{"id": 0})

	if form.IsValid() {
		test.Errorf("Form must be invalid")
	}
}
