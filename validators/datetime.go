package validators

import "time"

type DatetimeValidator interface {
	Validate(value time.Time) (time.Time, *Error)
}
