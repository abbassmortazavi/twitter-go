package errors

import "errors"

var (
	ErrValidation     = errors.New("validation error")
	ErrNotFound       = errors.New("not found")
	ErrBadCredentials = errors.New("email/password bad credentials")
)
