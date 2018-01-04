package validators

import "testing"

func TestIntegerMinValueValidator(test *testing.T) {
	var validator = IntegerMinValue(2)

	if _, err := validator.Validate(1); err == nil {
		test.Fatal("Must fail!")
	}

	if _, err := validator.Validate(2); err != nil {
		test.Fatalf("Error: %v", err)
	}
}

func TestIntegerMaxValueValidator(test *testing.T) {
	var validator = IntegerMaxValue(2)

	if _, err := validator.Validate(3); err == nil {
		test.Fatal("Must fail!")
	}

	if _, err := validator.Validate(1); err != nil {
		test.Fatalf("Error: %v", err)
	}
}
