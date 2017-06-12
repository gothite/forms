package fields

import (
	"strings"
	"testing"
	"time"

	"github.com/gothite/forms/validators"
)

func TestDatetimeIsRequired(test *testing.T) {
	field := Datetime{Required: false}

	if field.IsRequired() {
		test.Error("Returned invalid required flag")
		return
	}
}

func TestDatetimeGetDefault(test *testing.T) {
	field := Datetime{Default: 1.0}

	if value, ok := field.GetDefault().(float64); !ok || value != field.Default {
		test.Error("Returned invalid default value")
		return
	}
}

func TestDatetimeGetName(test *testing.T) {
	field := Datetime{Name: "test"}

	if name := field.GetName(); name != field.Name {
		test.Error("Returned invalid name")
		return
	}
}

func TestDatetimeGetValidators(test *testing.T) {
	field := Datetime{Validators: []validators.Validator{validators.MaxLength(1)}}

	if validators := field.GetValidators(); validators[0] == nil {
		test.Error("Returned invalid validators")
		return
	}
}

func TestDatetimeValidate(test *testing.T) {
	field := Datetime{}

	value := strings.Replace(time.RFC3339, "Z", "-", 1)
	result, err := field.Validate(value)

	if err != nil {
		test.Errorf("Returned error: %v", err)
		return
	}

	if result, ok := result.(time.Time); !ok || result.Format(time.RFC3339) != value {
		test.Errorf("Returned invalid value: %v", result)
		return
	}

	_, err = field.Validate(nil)

	if err == nil {
		test.Errorf("Finished without error on wrong string")
		return
	}

	_, err = field.Validate("")

	if err == nil {
		test.Errorf("Finished without error on wrong datetime string")
		return
	}
}
