package xcontext

import "context"

func Nop(ctx context.Context, fn func()) error {
	ch := make(chan struct{}, 1)

	go func(ch chan struct{}) {
		defer close(ch)
		fn()
		ch <- struct{}{}
	}(ch)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
		return nil
	}
}

func Error(ctx context.Context, fn func() error) error {
	ch := make(chan error, 1)

	go func(ch chan error) {
		defer close(ch)
		ch <- fn()
	}(ch)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
	}
}

func NopWaiter(fn func()) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return Nop(ctx, fn)
	}
}

func ErrorWaiter(fn func() error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return Error(ctx, fn)
	}
}
