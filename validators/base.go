package validators

// Validator describes validation function.
// Error must contains error code as message.
type Validator interface {
	Validate(value interface{}) (interface{}, *Error)
}

// Error describes error occured during validation.
type Error struct {
	Code       string
	Parameters []interface{}
}

// NewError initializes new Error instance.
func NewError(code string, parameters ...interface{}) *Error {
	return &Error{Code: code, Parameters: parameters}
}
