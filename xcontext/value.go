package xcontext

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientType = errors.New("insufficient type")
	ErrNoValue          = errors.New("no value")
)

func Value[O any](ctx Context, key any) (O, error) {
	var zero O

	val := ctx.Value(key)
	if val == nil {
		return zero, ErrNoValue
	}

	if out, ok := val.(O); ok {
		return out, nil
	}

	return zero, fmt.Errorf("%w: %T, expected %T", ErrInsufficientType, val, zero)
}
