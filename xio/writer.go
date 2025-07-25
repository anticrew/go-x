package xio

import (
	"io"
	"time"
)

type RuneWriter interface {
	WriteRune(r rune) (size int, err error)
}

type IntWriter interface {
	WriteInt(i int64, base int) (size int, err error)
}

type UintWriter interface {
	WriteUint(u uint64, base int) (size int, err error)
}

type ComplexWriter interface {
	WriteComplex(c complex128, fmt byte, prec, bitSize int) (size int, err error)
}

type FloatWriter interface {
	WriteFloat(f float64, fmt byte, prec, bitSize int) (size int, err error)
}

type BoolWriter interface {
	WriteBool(b bool) (size int, err error)
}

type TimeWriter interface {
	WriteTime(t time.Time, layout string) (size int, err error)
}

type Writer interface {
	io.ByteWriter
	io.StringWriter
	RuneWriter
	IntWriter
	UintWriter
	ComplexWriter
	FloatWriter
	BoolWriter
	TimeWriter
}
