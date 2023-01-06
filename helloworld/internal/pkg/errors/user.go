package errors

import "errors"

// error define
var (
	ErrBadParam     = errors.New("bad param")
	ErrUserNotFound = errors.New("user not found")
)
