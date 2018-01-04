package forms

import "fmt"

const (
	ErrorCodeUnknown uint = iota
	ErrorCodeInvalid
)

var FormErrors = map[uint]string{
	ErrorCodeUnknown: "Unknown error.",
	ErrorCodeInvalid: "Ensure that all values are valid.",
}

type Error struct {
	Code    uint             `json:"code"`
	Message string           `json:"message"`
	Errors  map[string]error `json:"errors"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s (code: %d)", err.Message, err.Code)
}
