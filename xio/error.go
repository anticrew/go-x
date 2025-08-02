package xio

import "errors"

var (
	ErrInvalidWhence = errors.New("seek: invalid whence")
	ErrInvalidOffset = errors.New("seek: invalid offset")
)
