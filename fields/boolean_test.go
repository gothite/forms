package fields

import (
	"testing"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

type TestBooleanValidator struct{}

func (validator *TestBooleanValidator) Validate(value bool) (bool, *validators.Error) {
	return value, validators.NewError(codes.Required)
}

func TestBooleanAsField(test *testing.T) {
	var _ Field = (*Boolean)(nil)
}

func TestBooleanIsRequired(test *testing.T) {
	var field = Boolean{Required: false}

	if field.IsRequired() {
		test.Fatal("Must be false!")
	}
}

func TestBooleanGetDefault(test *testing.T) {
	var field = Boolean{Default: true}

	if value, ok := field.GetDefault().(bool); !ok || value != field.Default {
		test.Errorf("Incorrect default!")
		test.Errorf("Expected: %t", field.Default)
		test.Errorf("Got: %t", value)
	}
}

func TestBooleanGetName(test *testing.T) {
	var field = Boolean{Name: "test"}

	if got := field.GetName(); got != field.Name {
		test.Errorf("Incorrect name!")
		test.Errorf("Expected: %s", field.Name)
		test.Errorf("Got: %s", got)
	}
}

func TestBooleanValidate(test *testing.T) {
	var field = Boolean{}
	var value = true

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %t", value)
		test.Errorf("Got: %t", got)
	}
}

func TestBooleanValidateInvalidValue(test *testing.T) {
	var field = Boolean{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestBooleanValidateStringNotAllowed(test *testing.T) {
	var field = Boolean{}

	if _, err := field.Validate("s"); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestBooleanValidateString(test *testing.T) {
	var field = Boolean{AllowStrings: true}
	var values = map[string]bool{
		"t":     true,
		"true":  true,
		"f":     false,
		"false": false,
	}

	for value, result := range values {
		if got, err := field.Validate(value); err != nil {
			test.Fatalf("Error for %s: %s", value, err)
		} else if got.(bool) != result {
			test.Errorf("Incorrect value for %s!", value)
			test.Errorf("Expected: %t", result)
			test.Errorf("Got: %t", got)
			return
		}
	}
}

func TestBooleanValidateIncorrectString(test *testing.T) {
	var field = Boolean{AllowStrings: true}

	if _, err := field.Validate("tru"); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestBooleanValidateNumber(test *testing.T) {
	var field = Boolean{AllowNumbers: true}

	var values = map[int]bool{
		1: true,
		0: false,
	}

	for value, result := range values {
		if got, err := field.Validate(value); err != nil {
			test.Fatalf("Error for %d: %s", value, err)
		} else if got.(bool) != result {
			test.Errorf("Incorrect result for %d!", value)
			test.Errorf("Expected: %t", result)
			test.Errorf("Got: %t", got)
			return
		}
	}
}

func TestBooleanValidateIncorrectNumber(test *testing.T) {
	var field = Boolean{AllowNumbers: true}

	if _, err := field.Validate(10); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestBooleanValidateFlagNil(test *testing.T) {
	var field = Boolean{Flag: true}

	if got, err := field.Validate(nil); err != nil {
		test.Fatalf("Error: %s", err)
	} else if !got.(bool) {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %t", true)
		test.Errorf("Got: %t", got)
	}
}

func TestBooleanValidateFlagEmptyString(test *testing.T) {
	var field = Boolean{Flag: true}

	if got, err := field.Validate(""); err != nil {
		test.Fatalf("Error: %s", err)
	} else if !got.(bool) {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %t", true)
		test.Errorf("Got: %t", got)
	}
}

func TestBooleanValidateValidators(test *testing.T) {
	var field = Boolean{
		Validators: []validators.BooleanValidator{
			&TestBooleanValidator{},
		},
	}

	if _, err := field.Validate(true); err == nil {
		test.Fatalf("Must fail!")
	}
}
