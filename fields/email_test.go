package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestEmailIsRequired(test *testing.T) {
	field := Email{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestEmailGetDefault(test *testing.T) {
	field := Email{Default: 1.0}

	if value, ok := field.GetDefault().(float64); !ok || value != field.Default {
		test.Error("Returned invalid default value")
		return
	}
}

func TestEmailGetName(test *testing.T) {
	field := Email{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestEmailGetValidators(test *testing.T) {
	field := Email{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestEmailValidate(test *testing.T) {
	field := Email{}

	value := "me@pyvim.com"
	result, err := field.Validate(value)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if result, ok := result.(string); !ok || result != value {
		test.Errorf("Returned invalid value: %v", result)
		return
	}

	_, err = field.Validate(nil)

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}

	_, err = field.Validate("me@pyvimcom")

	if err == nil {
		test.Errorf("Finished without error on wrong email string")
		return
	}
}
