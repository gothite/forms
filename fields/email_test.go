package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestEmailAsField(test *testing.T) {
	var _ Field = (*Email)(nil)
}

func TestEmailIsRequired(test *testing.T) {
	var field = Email{Required: false}

	if field.IsRequired() {
		test.Fatal("Must be false!")
	}
}

func TestEmailGetDefault(test *testing.T) {
	var field = Email{Default: "https://github.com/gothite/forms"}

	if value, ok := field.GetDefault().(string); !ok || value != field.Default {
		test.Errorf("Incorrect default!")
		test.Errorf("Expected: %s", field.Default)
		test.Errorf("Got: %s", value)
	}
}

func TestEmailGetName(test *testing.T) {
	field := Email{Name: "test"}

	if got := field.GetName(); got != field.Name {
		test.Errorf("Incorrect name!")
		test.Errorf("Expected: %s", field.Name)
		test.Errorf("Got: %s", got)
	}
}

func TestEmailValidate(test *testing.T) {
	var field = Email{}
	var value = "hello@bindlock.me"

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Errorf("Got: %s", got)
	}
}

func TestEmailValidateInvalidValue(test *testing.T) {
	var field = Email{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestEmailValidateIncorrectValue(test *testing.T) {
	var field = Email{}

	if _, err := field.Validate("hello@bindlockme"); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestEmailValidateValidators(test *testing.T) {
	var field = Email{
		Validators: []validators.StringValidator{
			validators.StringMaxLength(2),
		},
	}

	if _, err := field.Validate("hello@bindlock.me"); err == nil {
		test.Fatal("Must fail!")
	}
}
