package fields

import (
	"testing"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

type TestObjectValidator struct{}

func (validator *TestObjectValidator) Validate(value interface{}) (interface{}, *validators.Error) {
	return value, validators.NewError(codes.Required)
}

func TestObjectAsField(test *testing.T) {
	var _ Field = (*Object)(nil)
}

func TestObjectIsRequired(test *testing.T) {
	var field = Object{Required: false}

	if field.IsRequired() {
		test.Fatal("Must be false!")
	}
}

func TestObjectGetDefault(test *testing.T) {
	var field = Object{Default: true}

	if value, ok := field.GetDefault().(bool); !ok || value != field.Default {
		test.Errorf("Incorrect default!")
		test.Errorf("Expected: %t", field.Default)
		test.Errorf("Got: %t", value)
	}
}

func TestObjectGetName(test *testing.T) {
	var field = Object{Name: "test"}

	if got := field.GetName(); got != field.Name {
		test.Errorf("Incorrect name!")
		test.Errorf("Expected: %s", field.Name)
		test.Errorf("Got: %s", got)
	}
}

func TestObjectValidate(test *testing.T) {
	var field = Object{Fields: []Field{}}
	var value = map[int]int{0: 1}

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got.(map[int]int)[0] != value[0] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", value[0])
		test.Errorf("Got: %d", got.(map[int]int)[0])
	}
}

func TestObjectValidateInvalidValue(test *testing.T) {
	var field = Object{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestObjectValidateEmpty(test *testing.T) {
	var field = Object{}

	if _, err := field.Validate(map[int]int{}); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestObjectValidateFields(test *testing.T) {
	var field = Object{Fields: []Field{&Integer{Name: "0"}}}
	var value = map[string]int{"0": 1}

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got.(map[string]interface{})["0"].(int) != value["0"] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", value["0"])
		test.Errorf("Got: %d", got.(map[string]interface{})["0"].(int))
	}
}

func TestObjectValidateFieldsDefaultValue(test *testing.T) {
	var field = Object{Fields: []Field{&Integer{Name: "0", Required: false, Default: 1}}}
	var value = map[string]int{"1": 1}

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got.(map[string]interface{})["0"].(int) != 1 {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", 1)
		test.Errorf("Got: %d", got.(map[string]interface{})["0"].(int))
	}
}

func TestObjectValidateFieldsRequiredValue(test *testing.T) {
	var field = Object{Fields: []Field{&Integer{Name: "0", Required: true}}}
	var value = map[string]string{"1": "s"}

	if _, err := field.Validate(value); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestObjectValidateFieldsIncorrectValue(test *testing.T) {
	var field = Object{Fields: []Field{&Integer{Name: "0"}}}
	var value = map[string]string{"0": "s"}

	if _, err := field.Validate(value); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestObjectValidateValidators(test *testing.T) {
	var field = Object{
		Validators: []validators.Validator{
			&TestObjectValidator{},
		},
		AllowEmpty: true,
	}

	if _, err := field.Validate(map[int]int{}); err == nil {
		test.Fatalf("Must fail!")
	}
}
