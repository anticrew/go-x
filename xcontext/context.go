package xcontext

import (
	"context"
	"time"
)

type CancelFunc func(err error)

type Context interface {
	context.Context

	WithTimeout(t time.Duration) Context
	WithDeadline(d time.Time) Context
	WithCancel() Context

	Cause() error
	Cancel(err error)
}

type cancelContext struct {
	ctx    context.Context
	cancel context.CancelCauseFunc
}

func withCancel(ctx context.Context, cancel context.CancelFunc) *cancelContext {
	ctxCause, cancelCause := context.WithCancelCause(ctx)

	return withCancelCause(ctxCause, func(cause error) {
		defer cancel()

		cancelCause(cause)
	})
}

func withCancelCause(ctx context.Context, cancel context.CancelCauseFunc) *cancelContext {
	return &cancelContext{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *cancelContext) Deadline() (time.Time, bool) {
	return s.ctx.Deadline()
}

func (s *cancelContext) Done() <-chan struct{} {
	return s.ctx.Done()
}

func (s *cancelContext) Err() error {
	return s.ctx.Err()
}

func (s *cancelContext) Value(key any) any {
	return s.ctx.Value(key)
}

func (s *cancelContext) WithTimeout(t time.Duration) Context {
	return withCancel(context.WithTimeout(s.ctx, t))
}

func (s *cancelContext) WithDeadline(d time.Time) Context {
	return withCancel(context.WithDeadline(s.ctx, d))
}

func (s *cancelContext) WithCancel() Context {
	return withCancelCause(context.WithCancelCause(s.ctx))
}

func (s *cancelContext) Cause() error {
	return context.Cause(s.ctx)
}

func (s *cancelContext) Cancel(err error) {
	s.cancel(err)
}
