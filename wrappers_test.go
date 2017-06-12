package forms

import (
	"testing"

	"github.com/gothite/forms/fields"
	"github.com/gothite/forms/validators"
)

func TestFieldWrapper(test *testing.T) {
	field := &fields.Integer{
		Validators: []validators.Validator{validators.MaxValue(5)},
	}
	wrapper := FieldWrapper{field, nil, nil}

	if _, err := wrapper.Validate(4); err != nil {
		test.Errorf("Error occured: %v", err)
	}

	if _, err := wrapper.Validate("s"); err == nil {
		test.Errorf("Error must occured")
	}

	if _, err := wrapper.Validate(6); err == nil {
		test.Errorf("Error must occured")
	}
}
