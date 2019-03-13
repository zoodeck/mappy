package mappy

import "errors"

var (
	DoesNotContainError = errors.New("JsonObject does not contain given key")
	InvalidTypeError    = errors.New("Value for given key is invalid type")
	MissingValueError   = errors.New("Missing value error")
	ValuePresentError   = errors.New("Value is present, not map")
)
