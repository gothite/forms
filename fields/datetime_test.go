package fields

import (
	"testing"
	"time"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

type TestDatetimeValidator struct{}

func (validator *TestDatetimeValidator) Validate(value time.Time) (time.Time, *validators.Error) {
	return value, validators.NewError(codes.Required)
}

func TestDatetimeAsField(test *testing.T) {
	var _ Field = (*Datetime)(nil)
}

func TestDatetimeIsRequired(test *testing.T) {
	var field = Datetime{Required: false}

	if field.IsRequired() {
		test.Fatal("Must be false!")
	}
}

func TestDatetimeGetDefault(test *testing.T) {
	var field = Datetime{Default: time.Now()}

	if value, ok := field.GetDefault().(time.Time); !ok || value != field.Default {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", field.Default)
		test.Errorf("Got: %s", value)
	}
}

func TestDatetimeGetName(test *testing.T) {
	var field = Datetime{Name: "test"}

	if got := field.GetName(); got != field.Name {
		test.Errorf("Incorrect name!")
		test.Errorf("Expected: %s", field.Name)
		test.Errorf("Got: %s", got)
	}
}

func TestDatetimeValidate(test *testing.T) {
	var field = Datetime{}
	var value = time.Now()

	if got, err := field.Validate(value.Format(time.RFC3339)); err != nil {
		test.Fatalf("Error: %s", err)
	} else if value.Equal(got.(time.Time)) {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Errorf("Got: %s", got)
	}
}

func TestDatetimeValidateInvalidValue(test *testing.T) {
	var field = Datetime{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestDatetimeValidateIncorrectValue(test *testing.T) {
	var field = Datetime{}

	if _, err := field.Validate(time.Now().Format(time.RFC1123)); err == nil {
		test.Fatal("Must fail!")
	}
}

func TestDatetimeValidateValidators(test *testing.T) {
	var field = Datetime{
		Validators: []validators.DatetimeValidator{
			&TestDatetimeValidator{},
		},
	}

	if _, err := field.Validate(time.Now().Format(time.RFC3339)); err == nil {
		test.Fatalf("Must fail!")
	}
}
