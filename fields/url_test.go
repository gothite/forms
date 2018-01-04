package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestURLAsField(test *testing.T) {
	var _ Field = (*URL)(nil)
}

func TestURLIsRequired(test *testing.T) {
	field := URL{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestURLGetDefault(test *testing.T) {
	field := URL{Default: 1.0}

	if value, ok := field.GetDefault().(float64); !ok || value != field.Default {
		test.Error("Returned invalid default value")
		return
	}
}

func TestURLGetName(test *testing.T) {
	field := URL{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestURLValidate(test *testing.T) {
	var field = URL{}
	var value = "https://github.com/gothite/forms"

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Errorf("Got: %s", got)
	}
}

func TestURLValidateInvalidValue(test *testing.T) {
	var field = URL{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestURLValidateIncorrectValue(test *testing.T) {
	var field = URL{}

	if _, err := field.Validate(":/github.com/gothite/forms"); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestURLValidateValidators(test *testing.T) {
	field := URL{
		Validators: []validators.StringValidator{
			validators.StringMaxLength(2),
		},
	}

	if _, err := field.Validate("https://github.com/gothite/forms"); err == nil {
		test.Fatalf("Must fail!")
	}
}
