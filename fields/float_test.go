package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestFloatIsRequired(test *testing.T) {
	field := Float{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestFloatGetDefault(test *testing.T) {
	field := Float{Default: 1.0}

	if value, ok := field.GetDefault().(float64); !ok || value != field.Default {
		test.Error("Returned invalid default value")
		return
	}
}

func TestFloatGetName(test *testing.T) {
	field := Float{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestFloatGetValidators(test *testing.T) {
	field := Float{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestFloatValidate(test *testing.T) {
	field := Float{}

	value, err := field.Validate(5.0)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(float64); !ok || value != 5.0 {
		test.Error("Returned invalid string")
		return
	}

	_, err = field.Validate(nil)

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}
}

func TestFloatValidateAsString(test *testing.T) {
	field := Float{AllowStrings: true}

	value, err := field.Validate("2")

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(float64); !ok || value != 2 {
		test.Error("Returned invalid float")
		return
	}

	_, err = field.Validate("2a")

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}
}

func TestFloatValidateMinMax(test *testing.T) {
	field := Float{
		MinValue: validators.MinValue(2),
		MaxValue: validators.MaxValue(4),
	}

	if _, err := field.Validate(1.0); err == nil {
		test.Errorf("Error must occured")
		return
	}

	if _, err := field.Validate(5.0); err == nil {
		test.Errorf("Error must occured")
		return
	}
}
