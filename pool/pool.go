package pool

import "sync"

type Pool[T any] interface {
	Get() T
	Put(T)
}

type Options[T any] struct {
	Allow func(T) bool
	Reset func(T) T
}

type Option[T any] func(o Options[T]) Options[T]

type optionChain[T any] []Option[T]

func (o optionChain[T]) apply() Options[T] {
	var opt Options[T]

	for _, f := range o {
		opt = f(opt)
	}

	return opt
}

func WithAllow[T any](fn func(T) bool) Option[T] {
	return func(o Options[T]) Options[T] {
		o.Allow = fn
		return o
	}
}

func WithReset[T any](fn func(T) T) Option[T] {
	return func(o Options[T]) Options[T] {
		o.Reset = fn
		return o
	}
}

type typedSyncPool[T any] struct {
	p       *sync.Pool
	options Options[T]
}

func NewPool[T any](new func() T, options ...Option[T]) Pool[T] {
	opt := optionChain[T](options).apply()

	return &typedSyncPool[T]{
		p: &sync.Pool{
			New: func() any {
				return new()
			},
		},
		options: opt,
	}
}

func (p *typedSyncPool[T]) Get() T {
	return p.p.Get().(T)
}

func (p *typedSyncPool[T]) Put(t T) {
	if allow := p.options.Allow; allow != nil {
		if !allow(t) {
			return
		}
	}

	if reset := p.options.Reset; reset != nil {
		t = reset(t)
	}

	p.p.Put(t)
}
