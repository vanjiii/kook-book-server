package errors

import "errors"

var (
	ErrBadPass = errors.New("insecure password")
)
