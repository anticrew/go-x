package xio

import (
	"io"
	"time"
)

type StringReader interface {
	ReadString(delim byte) (string, error)
}

type IntReader interface {
	ReadInt(base, bitSize int) (int64, error)
}

type UintReader interface {
	ReadUint(base, bitSize int) (uint64, error)
}

type ComplexReader interface {
	ReadComplex(bitSize int) (complex128, error)
}

type FloatReader interface {
	ReadFloat(base, bitSize int) (float64, error)
}

type BoolReader interface {
	ReadBool(base int) (bool, error)
}

type TimeReader interface {
	ReadTime(layouts ...string) (time.Time, error)
}

type Reader interface {
	io.ByteReader
	StringReader
	io.RuneReader
	IntReader
	UintReader
	ComplexReader
	FloatReader
	BoolReader
	TimeReader
}
