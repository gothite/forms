package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

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

func TestStringGetValidators(test *testing.T) {
	field := String{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestStringValidate(test *testing.T) {
	field := String{}

	str := "test"
	value, err := field.Validate(str)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(string); !ok || value != str {
		test.Error("Returned invalid string")
		return
	}

	_, err = field.Validate(nil)

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}
}

func TestStringValidateBlank(test *testing.T) {
	field := String{}

	_, err := field.Validate("")

	if err == nil {
		test.Errorf("Finished without error on blank string")
		return
	}

	field.AllowBlank = true

	value, err := field.Validate("")

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(string); !ok || value != "" {
		test.Error("Returned invalid string")
		return
	}
}

func TestStringValidateLength(test *testing.T) {
	field := String{
		MinLength: validators.MinLength(2),
		MaxLength: validators.MaxLength(4),
	}

	_, err := field.Validate("12345")

	if err == nil {
		test.Errorf("Finished without error on too long string")
		return
	}

	_, err = field.Validate("1")

	if err == nil {
		test.Errorf("Finished without error on too small string")
		return
	}
}
