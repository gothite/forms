package validators

import "testing"

func TestArrayMinLengthValidator(test *testing.T) {
	var validator = ArrayMinLength(2)

	if _, err := validator.Validate([]int{1}); err == nil {
		test.Fatal("Must fail!")
	}

	if _, err := validator.Validate([]int{1, 2}); err != nil {
		test.Fatalf("Error: %v", err)
	}
}

func TestArrayMaxLengthValidator(test *testing.T) {
	var validator = ArrayMaxLength(2)

	if _, err := validator.Validate([]int{1, 2, 3}); err == nil {
		test.Fatal("Must fail!")
	}

	if _, err := validator.Validate([]int{1}); err != nil {
		test.Fatalf("Error: %v", err)
	}
}
