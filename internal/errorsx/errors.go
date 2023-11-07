package errorsx

import "errors"

var (
	ErrBadPass = errors.New("insecure password")
)
