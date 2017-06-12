package fields

import (
	"github.com/gothite/forms/validators"
)

// Field describes a field interface.
type Field interface {
	GetName() string
	GetValidators() []validators.Validator
	IsRequired() bool
	GetDefault() interface{}
	GetError(code string, parameters ...interface{}) error
	Validate(value interface{}) (interface{}, error)
}
