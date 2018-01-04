package fields

// Field describes a field interface.
type Field interface {
	GetName() string
	IsRequired() bool
	GetDefault() interface{}
	GetError(code string, parameters ...interface{}) error
	Validate(value interface{}) (interface{}, error)
}
