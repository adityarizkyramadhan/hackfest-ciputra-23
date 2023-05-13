package customserror

import "errors"

var (
	ErrDatabaseNull       = errors.New("database nil pointer, init connection to db first")
	ErrIncorectPassword   = errors.New("incorrect password")
	ErrUsernameExist      = errors.New("username already exist!, try another one")
	ErrTimeout            = errors.New("timeout")
	ErrMissingBearerToken = errors.New("bearer token is not provided")
	ErrNullPointer        = errors.New("the pointer value is nil")
	ErrNoQuery            = errors.New("no query provided")
)
