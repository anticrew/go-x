package xio

import (
	"context"
	"io"
)

type writer struct {
	ctx    context.Context
	origin io.Writer
}

func NewContextWriter(ctx context.Context, origin io.Writer) io.Writer {
	return &writer{
		ctx:    ctx,
		origin: origin,
	}
}

func (w *writer) Write(p []byte) (n int, err error) {
	if err = w.ctx.Err(); err != nil {
		return 0, err
	}

	return w.origin.Write(p)
}
