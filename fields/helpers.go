package fields

import (
	"fmt"

	"github.com/gothite/forms/codes"
)

func getError(code uint, value interface{}, custom, builtin map[uint]string, errorFunc ErrorFunc, parameters ...interface{}) error {
	if message, ok := custom[code]; ok {
		return fmt.Errorf(message, parameters...)
	} else if message, ok := builtin[code]; ok {
		return fmt.Errorf(message, parameters...)
	} else if errorFunc != nil {
		return errorFunc(code, value, parameters)
	}

	return fmt.Errorf(builtin[codes.Unknown], parameters...)
}
