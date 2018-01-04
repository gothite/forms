package validators

import "testing"

func TestStringMinLengthValidator(test *testing.T) {
	var validator = StringMinLength(2)

	if _, err := validator.Validate("1"); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate("12"); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}
}

func TestMaxLengthValidator(test *testing.T) {
	var validator = StringMaxLength(2)

	if _, err := validator.Validate("123"); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate("12"); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}
}
