package recipe

import "errors"

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("found duplicate")
)
