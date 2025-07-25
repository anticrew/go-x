package xio

import (
	"context"
	"io"
)

type reader struct {
	ctx    context.Context
	origin io.Reader
}

func NewContextReader(ctx context.Context, origin io.Reader) io.Reader {
	return &reader{
		ctx:    ctx,
		origin: origin,
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	if err = r.ctx.Err(); err != nil {
		return 0, err
	}

	return r.origin.Read(p)
}

type readCloser struct {
	ctx    context.Context
	origin io.ReadCloser
}

func NewContextReadCloser(ctx context.Context, origin io.ReadCloser) io.ReadCloser {
	return &readCloser{
		ctx:    ctx,
		origin: origin,
	}
}

func (r *readCloser) Read(p []byte) (n int, err error) {
	if err = r.ctx.Err(); err != nil {
		return 0, err
	}

	return r.origin.Read(p)
}

func (r *readCloser) Close() error {
	return r.origin.Close()
}
