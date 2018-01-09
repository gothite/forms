package fields

import (
	"strconv"
	"testing"

	"github.com/gothite/forms/validators"
)

func TestIntegerAsField(test *testing.T) {
	var _ Field = (*Integer)(nil)
}

func TestIntegerIsRequired(test *testing.T) {
	field := Integer{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestIntegerGetDefault(test *testing.T) {
	field := Integer{Default: 1}

	if value, ok := field.GetDefault().(int); !ok || value != field.Default {
		test.Error("Returned invalid default value")
		return
	}
}

func TestIntegerGetName(test *testing.T) {
	field := Integer{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestIntegerValidate(test *testing.T) {
	var field = Integer{}
	var value = 5

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", value)
		test.Errorf("Got: %d", got)
	}
}

func TestIntegerValidateIncorrectValue(test *testing.T) {
	var field = Integer{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestIntegerValidateStringNotAllowed(test *testing.T) {
	var field = Integer{}

	if _, err := field.Validate("s"); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestIntegerValidateAsString(test *testing.T) {
	var field = Integer{AllowStrings: true}
	var value = 2

	if got, err := field.Validate(strconv.Itoa(value)); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", value)
		test.Errorf("Got: %d", got)
		return
	}
}

func TestIntegerValidateInvalidString(test *testing.T) {
	var field = Integer{AllowStrings: true}

	if _, err := field.Validate(nil); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestIntegerValidateIncorrectString(test *testing.T) {
	var field = Integer{AllowStrings: true}

	if _, err := field.Validate("2a"); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestIntegerValidateValidators(test *testing.T) {
	var field = Integer{
		Validators: []validators.IntegerValidator{
			validators.IntegerMinValue(2),
			validators.IntegerMaxValue(4),
		},
	}

	if _, err := field.Validate(1.0); err == nil {
		test.Fatalf("Must fail!")
	}

	if _, err := field.Validate(5.0); err == nil {
		test.Fatalf("Must fail!")
	}
}
