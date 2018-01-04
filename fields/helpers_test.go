package fields

import (
	"errors"
	"testing"

	"github.com/gothite/forms/codes"
)

func TestGetErrorFromCustom(test *testing.T) {
	var code = codes.Unknown
	var message = "Unknown"
	var messages = map[uint]string{code: message}

	if got := getError(nil, code, nil, messages, nil, nil); got.Error() != message {
		test.Errorf("Invalid error!")
		test.Errorf("Expected: %s", message)
		test.Errorf("Got: %s", got)
	}
}

func TestGetErrorFromBuiltin(test *testing.T) {
	var code = codes.Invalid
	var message = "Invalid"
	var messages = map[uint]string{code: message}

	if got := getError(nil, code, nil, nil, messages, nil); got.Error() != message {
		test.Errorf("Invalid error!")
		test.Errorf("Expected: %s", message)
		test.Errorf("Got: %s", got)
	}
}

func TestGetErrorFromErrorFunc(test *testing.T) {
	var field = &Email{}
	var code = codes.Invalid
	var message = "Invalid"
	var errorFunc = func(field Field, code uint, value interface{}, parameters ...interface{}) error {
		return errors.New(message)
	}

	if got := getError(field, code, nil, nil, nil, errorFunc); got.Error() != message {
		test.Errorf("Invalid error!")
		test.Errorf("Expected: %s", message)
		test.Errorf("Got: %s", got)
	}
}
