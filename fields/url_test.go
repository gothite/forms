package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

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

func TestURLGetValidators(test *testing.T) {
	field := URL{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestURLValidate(test *testing.T) {
	field := URL{}

	value := "https://github.com/gothite/forms"
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

	_, err = field.Validate("://")

	if err == nil {
		test.Errorf("Finished without error on wrong URL string")
		return
	}
}
