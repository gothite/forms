package validators

import "testing"

func TestMinLengthValidator(test *testing.T) {
	validator := MinLength(2)

	if _, err := validator.Validate(1); err == nil {
		test.Error("Error must occurred")
		return
	}

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
	validator := MaxLength(2)

	if _, err := validator.Validate(1); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate("123"); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate("12"); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}
}

func TestMinValueValidator(test *testing.T) {
	validator := MinValue(2.0)

	if _, err := validator.Validate("s"); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate(1); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate(2); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}

	if _, err := validator.Validate(2.0); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}
}

func TestMaxValueValidator(test *testing.T) {
	validator := MaxValue(2.0)

	if _, err := validator.Validate("s"); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate(3.5); err == nil {
		test.Error("Error must occurred")
		return
	}

	if _, err := validator.Validate(2); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}

	if _, err := validator.Validate(1.8); err != nil {
		test.Errorf("Error returned: %v", err)
		return
	}
}
