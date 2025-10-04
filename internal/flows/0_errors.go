package flows

import "errors"

var (
	ErrResourceNotFound = errors.New("some resource could not be found")
	ErrServerFailure    = errors.New("server issue check logs")
)
