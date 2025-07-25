package xcontext

import (
	"context"
)

var _background = withCancelCause(context.Background(), func(cause error) {
	_ = cause
})

// Background returns a non-nil, empty [Context]. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming requests.
// Always returns the same instance.
func Background() Context {
	return _background
}

func wrap(ctx context.Context, cancel context.CancelCauseFunc) Context {
	return &cancelContext{
		ctx:    ctx,
		cancel: cancel,
	}
}
