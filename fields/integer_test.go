package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

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

func TestIntegerGetValidators(test *testing.T) {
	field := Integer{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestIntegerValidate(test *testing.T) {
	field := Integer{}

	value, err := field.Validate(5)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(int); !ok || value != 5 {
		test.Error("Returned invalid string")
		return
	}

	_, err = field.Validate(nil)

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}
}

func TestIntegerValidateAsString(test *testing.T) {
	field := Integer{AllowStrings: true}

	value, err := field.Validate("2")

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(int); !ok || value != 2 {
		test.Error("Returned invalid integer")
		return
	}

	_, err = field.Validate("2a")

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}
}

func TestIntegerValidateMinMax(test *testing.T) {
	field := Integer{
		MinValue: validators.MinValue(2),
		MaxValue: validators.MaxValue(4),
	}

	if _, err := field.Validate(1); err == nil {
		test.Errorf("Error must occured")
		return
	}

	if _, err := field.Validate(5); err == nil {
		test.Errorf("Error must occured")
		return
	}
}
