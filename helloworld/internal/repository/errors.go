package repository

import "errors"

// err .
var (
	ErrBadParam     = errors.New("bad param")
	ErrUserNotFound = errors.New("user not found")
)
