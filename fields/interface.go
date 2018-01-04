package fields

// Field describes a field interface.
type Field interface {
	GetName() string
	IsRequired() bool
	GetDefault() interface{}
	GetError(code uint, value interface{}, parameters ...interface{}) error
	Validate(value interface{}) (interface{}, error)
}

type ErrorFunc func(code uint, value interface{}, parameters ...interface{}) error
