package validators

type BooleanValidator interface {
	Validate(value bool) (bool, *Error)
}
