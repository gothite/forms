package fields

import (
	"testing"

	"github.com/gothite/forms/validators"
)

func TestArrayAsField(test *testing.T) {
	var _ Field = (*Array)(nil)
}

func TestArrayIsRequired(test *testing.T) {
	var field = Array{Required: false}

	if field.IsRequired() {
		test.Fatal("Must be false!")
	}
}

func TestArrayGetDefault(test *testing.T) {
	var field = Array{Default: true}

	if value, ok := field.GetDefault().(bool); !ok || value != field.Default {
		test.Errorf("Incorrect default!")
		test.Errorf("Expected: %t", field.Default)
		test.Errorf("Got: %t", value)
	}
}

func TestArrayGetName(test *testing.T) {
	var field = Array{Name: "test"}

	if got := field.GetName(); got != field.Name {
		test.Errorf("Incorrect name!")
		test.Errorf("Expected: %s", field.Name)
		test.Errorf("Got: %s", got)
	}
}

func TestArrayValidate(test *testing.T) {
	var field = Array{}
	var value = []int{1}

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got.([]int)[0] != value[0] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", value[0])
		test.Errorf("Got: %d", got.([]int)[0])
	}
}

func TestArrayValidateInvalidValue(test *testing.T) {
	var field = Array{}

	if _, err := field.Validate(nil); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestArrayValidateEmpty(test *testing.T) {
	var field = Array{}

	if _, err := field.Validate([]int{}); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestArrayValidateField(test *testing.T) {
	var field = Array{Field: &Integer{}}
	var value = []int{1}

	if got, err := field.Validate(value); err != nil {
		test.Fatalf("Error: %s", err)
	} else if got.([]int)[0] != value[0] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %d", value[0])
		test.Errorf("Got: %d", got.([]int)[0])
	}
}

func TestArrayValidateFieldIncorrectValue(test *testing.T) {
	var field = Array{Field: &Integer{}}
	var value = []string{"s"}

	if _, err := field.Validate(value); err == nil {
		test.Fatalf("Must fail!")
	}
}

func TestArrayValidateValidators(test *testing.T) {
	var field = Array{
		Validators: []validators.Validator{
			validators.ArrayMinLength(1),
			validators.ArrayMaxLength(2),
		},
	}

	if _, err := field.Validate([]int{}); err == nil {
		test.Fatalf("Must fail!")
	}

	if _, err := field.Validate([]int{0, 0, 0}); err == nil {
		test.Fatalf("Must fail!")
	}
}
