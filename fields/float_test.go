package fields

import (
	"fmt"
	"testing"

	"github.com/gothite/forms/validators"
)

func TestFloatAsField(test *testing.T) {
	var _ Field = (*Float)(nil)
}

func TestFloatIsRequired(test *testing.T) {
	field := Float{Required: false}

	if field.IsRequired() {
		test.Fatal("Must be false!")
	}
}

func TestFloatGetDefault(test *testing.T) {
	field := Float{Default: 1.0}

	if value, ok := field.GetDefault().(float64); !ok || value != field.Default {
		test.Errorf("Incorrect default!")
		test.Errorf("Expected: %f", field.Default)
		test.Errorf("Got: %f", value)
	}
}

func TestFloatGetName(test *testing.T) {
	field := Float{Name: "test"}

	if got := field.GetName(); got != field.Name {
		test.Errorf("Incorrect name!")
		test.Errorf("Expected: %s", field.Name)
		test.Errorf("Got: %s", got)
	}
}

func TestFloatValidate(test *testing.T) {
	var field = Float{}
	var value = 5.0

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %f", value)
		test.Errorf("Got: %f", got)
	}
}

func TestFloatValidateInvalidValue(test *testing.T) {
	var field = Float{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestFloatValidateString(test *testing.T) {
	var field = Float{AllowStrings: true}
	var value = 2.0

	if got, err := field.Validate(fmt.Sprintf("%f", value)); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %f", value)
		test.Errorf("Got: %f", got)
	}
}

func TestFloatValidateInvalidString(test *testing.T) {
	var field = Float{AllowStrings: true}

	if _, err := field.Validate(nil); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestFloatValidateIncorrectString(test *testing.T) {
	var field = Float{AllowStrings: true}

	if _, err := field.Validate("2.a"); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestFloatValidateValidators(test *testing.T) {
	var field = Float{
		Validators: []validators.FloatValidator{
			validators.FloatMinValue(2),
			validators.FloatMaxValue(4),
		},
	}

	if _, err := field.Validate(1.0); err == nil {
		test.Fatalf("Must fail!")
	}

	if _, err := field.Validate(5.0); err == nil {
		test.Fatalf("Must fail!")
	}
}
