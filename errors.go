package forms

import (
	"fmt"

	"github.com/gothite/forms/codes"
)

var Errors = map[uint]string{
	codes.Unknown: "Unknown error.",
	codes.Invalid: "Ensure that all values are valid.",
}

type Error struct {
	Code    uint             `json:"code"`
	Message string           `json:"message"`
	Errors  map[string]error `json:"errors"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s (code: %d)", err.Message, err.Code)
}
