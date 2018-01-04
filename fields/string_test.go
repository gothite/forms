package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestStringAsField(test *testing.T) {
	var _ Field = (*String)(nil)
}

func TestStringIsRequired(test *testing.T) {
	field := String{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestStringGetDefault(test *testing.T) {
	field := String{Default: "test"}

	if value, ok := field.GetDefault().(string); !ok || value != field.Default {
		test.Error("Returned invalid default value")
		return
	}
}

func TestStringGetName(test *testing.T) {
	field := String{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestStringValidate(test *testing.T) {
	var field = String{}
	var value = "string"

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Errorf("Got: %s", got)
	}
}

func TestStringValidateInvalidValue(test *testing.T) {
	var field = String{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestStringValidateBlank(test *testing.T) {
	var field = String{}
	var value = ""

	if _, err := field.Validate(value); err == nil {
		test.Fatal("Must fail!")
	}

	field.AllowBlank = true

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Errorf("Got: %s", got)
	}
}

func TestStringValidateValidators(test *testing.T) {
	field := String{
		Validators: []validators.StringValidator{
			validators.StringMinLength(2),
			validators.StringMaxLength(4),
		},
	}

	if _, err := field.Validate("12345"); err == nil {
		test.Fatalf("Must fail!")
	}

	if _, err := field.Validate("1"); err == nil {
		test.Fatalf("Must fail!")
	}
}
