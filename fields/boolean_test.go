package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestBooleanIsRequired(test *testing.T) {
	field := Boolean{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestBooleanGetDefault(test *testing.T) {
	field := Boolean{Default: true}

	if value, ok := field.GetDefault().(bool); !ok || !value {
		test.Error("Returned invalid default value")
		return
	}
}

func TestBooleanGetName(test *testing.T) {
	field := Boolean{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestBooleanGetValidators(test *testing.T) {
	field := Boolean{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestBooleanValidate(test *testing.T) {
	field := Boolean{}

	value, err := field.Validate(true)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(bool); !ok || !value {
		test.Error("Returned invalid boolean")
		return
	}

	value, err = field.Validate(false)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(bool); !ok || value {
		test.Error("Returned invalid boolean")
		return
	}

	_, err = field.Validate(nil)

	if err == nil {
		test.Errorf("Finished without error on wrong boolean")
		return
	}
}

func TestBooleanValidateAsString(test *testing.T) {
	field := Boolean{AllowStrings: true}

	value, err := field.Validate("t")

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(bool); !ok || !value {
		test.Error("Returned invalid boolean")
		return
	}

	value, err = field.Validate("f")

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(bool); !ok || value {
		test.Error("Returned invalid boolean")
		return
	}
}

func TestBooleanValidateAsNumber(test *testing.T) {
	field := Boolean{AllowNumbers: true}

	value, err := field.Validate(1)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(bool); !ok || !value {
		test.Error("Returned invalid boolean")
		return
	}

	value, err = field.Validate(0)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if value, ok := value.(bool); !ok || value {
		test.Error("Returned invalid boolean")
		return
	}
}
