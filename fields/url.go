package fields

import (
	"regexp"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// https://github.com/asaskevich/govalidator
var urlRE = regexp.MustCompile(
	`^((ftp|tcp|udp|wss?|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(\[(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`,
)

// URLErrors is a code-error mapping for URL field.
var URLErrors = map[uint]string{
	codes.Required: "This field is required.",
	codes.Invalid:  "Ensure this value is valid URL string.",
}

// URL is integer field.
type URL struct {
	Name       string
	Validators []validators.StringValidator
	Required   bool
	Default    float64
	Errors     map[uint]string
	ErrorFunc  ErrorFunc
}

// IsRequired returns true if field is required.
func (field *URL) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *URL) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *URL) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *URL) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(field, code, value, field.Errors, URLErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *URL) Validate(v interface{}) (interface{}, error) {
	value, ok := v.(string)

	if !ok {
		return nil, field.GetError(codes.Invalid, v)
	}

	if !urlRE.MatchString(value) {
		return nil, field.GetError(codes.Invalid, v)
	}

	for _, validator := range field.Validators {
		var err *validators.Error

		value, err = validator.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, v, err.Parameters...)
		}
	}

	return value, nil
}
