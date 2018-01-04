package validators

import "testing"

func TestFloatMinValueValidator(test *testing.T) {
	var validator = FloatMinValue(2.0)

	if _, err := validator.Validate(1.9); err == nil {
		test.Fatal("Must fail!")
	}

	if _, err := validator.Validate(2.0); err != nil {
		test.Fatalf("Error: %v", err)
	}
}

func TestFloatMaxValueValidator(test *testing.T) {
	var validator = FloatMaxValue(2.0)

	if _, err := validator.Validate(3.5); err == nil {
		test.Fatal("Must fail!")
	}

	if _, err := validator.Validate(1.8); err != nil {
		test.Fatalf("Error: %v", err)
	}
}
