package xio

import (
	"bytes"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuffer_WriteByte(t *testing.T) {
	t.Parallel()

	type testCase struct {
		b        byte
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-a": {
			b:        'a',
			buf:      NewBuffer(),
			expected: "a",
		},
		"empty-1": {
			b:        '1',
			buf:      NewBuffer(),
			expected: "1",
		},
		"empty-#": {
			b:        '#',
			buf:      NewBuffer(),
			expected: "#",
		},
		"content-a": {
			b:        'a',
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-a",
		},
		"content-1": {
			b:        '1',
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-1",
		},
		"content-#": {
			b:        '#',
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-#",
		},
		"quoted-a": {
			b:        'a',
			buf:      NewBuffer().WithQuotes(),
			expected: `"a"`,
		},
		"quoted-1": {
			b:        '1',
			buf:      NewBuffer().WithQuotes(),
			expected: `"1"`,
		},
		"quoted-#": {
			b:        '#',
			buf:      NewBuffer().WithQuotes(),
			expected: `"#"`,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteByte(test.b)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteBytes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		b        []byte
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-abc": {
			b:        []byte{'a', 'b', 'c'},
			buf:      NewBuffer(),
			expected: "abc",
		},
		"empty-123": {
			b:        []byte{'1', '2', '3'},
			buf:      NewBuffer(),
			expected: "123",
		},
		"empty-#@$": {
			b:        []byte{'#', '@', '$'},
			buf:      NewBuffer(),
			expected: "#@$",
		},
		"content-a": {
			b:        []byte{'a', 'b', 'c'},
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-abc",
		},
		"content-1": {
			b:        []byte{'1', '2', '3'},
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-123",
		},
		"content-#@$": {
			b:        []byte{'#', '@', '$'},
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-#@$",
		},
		"quoted-abc": {
			b:        []byte{'a', 'b', 'c'},
			buf:      NewBuffer().WithQuotes(),
			expected: `"abc"`,
		},
		"quoted-123": {
			b:        []byte{'1', '2', '3'},
			buf:      NewBuffer().WithQuotes(),
			expected: `"123"`,
		},
		"quoted-#@$": {
			b:        []byte{'#', '@', '$'},
			buf:      NewBuffer().WithQuotes(),
			expected: `"#@$"`,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteBytes(test.b)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		s        string
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-abc": {
			s:        "abc",
			buf:      NewBuffer(),
			expected: "abc",
		},
		"empty-123": {
			s:        "123",
			buf:      NewBuffer(),
			expected: "123",
		},
		"empty-#@$": {
			s:        "#@$",
			buf:      NewBuffer(),
			expected: "#@$",
		},
		"content-a": {
			s:        "abc",
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-abc",
		},
		"content-1": {
			s:        "123",
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-123",
		},
		"content-#@$": {
			s:        "#@$",
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-#@$",
		},
		"quoted-abc": {
			s:        "abc",
			buf:      NewBuffer().WithQuotes(),
			expected: `"abc"`,
		},
		"quoted-123": {
			s:        "123",
			buf:      NewBuffer().WithQuotes(),
			expected: `"123"`,
		},
		"quoted-#@$": {
			s:        "#@$",
			buf:      NewBuffer().WithQuotes(),
			expected: `"#@$"`,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteString(test.s)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteInt64(t *testing.T) {
	t.Parallel()

	type testCase struct {
		i        int64
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-10": {
			i:        10,
			buf:      NewBuffer(),
			expected: "10",
		},
		"empty-0": {
			i:        0,
			buf:      NewBuffer(),
			expected: "0",
		},
		"empty-(-10)": {
			i:        -10,
			buf:      NewBuffer(),
			expected: "-10",
		},
		"content-10": {
			i:        10,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-10",
		},
		"content-0": {
			i:        0,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-0",
		},
		"content-(-10)": {
			i:        -10,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content--10",
		},
		"quoted-10": {
			i:        10,
			buf:      NewBuffer().WithQuotes(),
			expected: `"10"`,
		},
		"quoted-0": {
			i:        0,
			buf:      NewBuffer().WithQuotes(),
			expected: `"0"`,
		},
		"quoted-(-10)": {
			i:        -10,
			buf:      NewBuffer().WithQuotes(),
			expected: `"-10"`,
		},
		"edge-max": {
			i:        math.MaxInt64,
			buf:      NewBuffer(),
			expected: "9223372036854775807",
		},
		"edge-min": {
			i:        -math.MaxInt64,
			buf:      NewBuffer(),
			expected: "-9223372036854775807",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteInt64(test.i)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteUint64(t *testing.T) {
	t.Parallel()

	type testCase struct {
		i        uint64
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-10": {
			i:        10,
			buf:      NewBuffer(),
			expected: "10",
		},
		"empty-0": {
			i:        0,
			buf:      NewBuffer(),
			expected: "0",
		},
		"content-10": {
			i:        10,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-10",
		},
		"content-0": {
			i:        0,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-0",
		},
		"quoted-10": {
			i:        10,
			buf:      NewBuffer().WithQuotes(),
			expected: `"10"`,
		},
		"quoted-0": {
			i:        0,
			buf:      NewBuffer().WithQuotes(),
			expected: `"0"`,
		},
		"edge-max": {
			i:        math.MaxUint64,
			buf:      NewBuffer(),
			expected: "18446744073709551615",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteUint64(test.i)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteFloat64(t *testing.T) {
	t.Parallel()

	type testCase struct {
		i        float64
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-3.14159": {
			i:        3.14159,
			buf:      NewBuffer(),
			expected: "3.14159",
		},
		"empty-10": {
			i:        10,
			buf:      NewBuffer(),
			expected: "10",
		},
		"empty-0": {
			i:        0,
			buf:      NewBuffer(),
			expected: "0",
		},
		"empty-(-10)": {
			i:        -10,
			buf:      NewBuffer(),
			expected: "-10",
		},
		"empty-(-3.14159)": {
			i:        -3.14159,
			buf:      NewBuffer(),
			expected: "-3.14159",
		},
		"content-3.14159": {
			i:        3.14159,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-3.14159",
		},
		"content-10": {
			i:        10,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-10",
		},
		"content-0": {
			i:        0,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-0",
		},
		"content-(-10)": {
			i:        -10,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content--10",
		},
		"content-(-3.14159)": {
			i:        -3.14159,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content--3.14159",
		},
		"quoted-3.14159": {
			i:        3.14159,
			buf:      NewBuffer().WithQuotes(),
			expected: `"3.14159"`,
		},
		"quoted-10": {
			i:        10,
			buf:      NewBuffer().WithQuotes(),
			expected: `"10"`,
		},
		"quoted-0": {
			i:        0,
			buf:      NewBuffer().WithQuotes(),
			expected: `"0"`,
		},
		"quoted-(-10)": {
			i:        -10,
			buf:      NewBuffer().WithQuotes(),
			expected: `"-10"`,
		},
		"quoted-(-3.14159)": {
			i:        -3.14159,
			buf:      NewBuffer().WithQuotes(),
			expected: `"-3.14159"`,
		},
		"edge-max": {
			i:   math.MaxFloat64,
			buf: NewBuffer(),
			expected: "179769313486231570000000000000000000000000000000000000000000000000000000000000000000" +
				"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
				"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
				"000000000000000000000000000000000000000000000",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteFloat64(test.i, 64)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteBool(t *testing.T) {
	t.Parallel()

	type testCase struct {
		b        bool
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-true": {
			b:        true,
			buf:      NewBuffer(),
			expected: "true",
		},
		"empty-false": {
			b:        false,
			buf:      NewBuffer(),
			expected: "false",
		},
		"content-true": {
			b:        true,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-true",
		},
		"content-false": {
			b:        false,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-false",
		},
		"quoted-true": {
			b:        true,
			buf:      NewBuffer().WithQuotes(),
			expected: `"true"`,
		},
		"quoted-false": {
			b:        false,
			buf:      NewBuffer().WithQuotes(),
			expected: `"false"`,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteBool(test.b)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteTime(t *testing.T) {
	t.Parallel()

	type testCase struct {
		t        time.Time
		layout   string
		buf      *Buffer
		expected string
	}

	loc, err := time.LoadLocation("Europe/Moscow")
	require.NoError(t, err)

	var (
		dateUTC = time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC)
		dateMOW = time.Date(2025, 5, 1, 10, 0, 0, 0, loc)
	)

	testCases := map[string]testCase{
		"empty-utc-rfc3339": {
			t:        dateUTC,
			layout:   time.RFC3339,
			buf:      NewBuffer(),
			expected: dateUTC.Format(time.RFC3339),
		},
		"empty-mow-rfc3339": {
			t:        dateMOW,
			layout:   time.RFC3339,
			buf:      NewBuffer(),
			expected: dateMOW.Format(time.RFC3339),
		},
		"empty-utc-empty": {
			t:        dateUTC,
			layout:   "",
			buf:      NewBuffer(),
			expected: "",
		},
		"content-utc-rfc3339": {
			t:        dateUTC,
			layout:   time.RFC3339,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-" + dateUTC.Format(time.RFC3339),
		},
		"content-mow-rfc3339": {
			t:        dateMOW,
			layout:   time.RFC3339,
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-" + dateMOW.Format(time.RFC3339),
		},
		"content-utc-empty": {
			t:        dateUTC,
			layout:   "",
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-",
		},
		"quoted-utc-rfc3339": {
			t:        dateUTC,
			layout:   time.RFC3339,
			buf:      NewBuffer().WithQuotes(),
			expected: fmt.Sprintf(`"%s"`, dateUTC.Format(time.RFC3339)),
		},
		"quoted-mow-rfc3339": {
			t:        dateMOW,
			layout:   time.RFC3339,
			buf:      NewBuffer().WithQuotes(),
			expected: fmt.Sprintf(`"%s"`, dateMOW.Format(time.RFC3339)),
		},
		"quoted-utc-empty": {
			t:        dateUTC,
			layout:   "",
			buf:      NewBuffer().WithQuotes(),
			expected: `""`,
		},
		"empty-utc-custom": {
			t:        dateUTC,
			layout:   "2006-01-02",
			buf:      NewBuffer(),
			expected: dateUTC.Format("2006-01-02"),
		},
		"empty-mow-custom": {
			t:        dateMOW,
			layout:   "2006-01-02",
			buf:      NewBuffer(),
			expected: dateMOW.Format("2006-01-02"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.buf.WriteTime(test.t, test.layout)
			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_Write(t *testing.T) {
	t.Parallel()

	type testCase struct {
		b        []byte
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty-abc": {
			b:        []byte{'a', 'b', 'c'},
			buf:      NewBuffer(),
			expected: "abc",
		},
		"empty-123": {
			b:        []byte{'1', '2', '3'},
			buf:      NewBuffer(),
			expected: "123",
		},
		"empty-#@$": {
			b:        []byte{'#', '@', '$'},
			buf:      NewBuffer(),
			expected: "#@$",
		},
		"content-a": {
			b:        []byte{'a', 'b', 'c'},
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-abc",
		},
		"content-1": {
			b:        []byte{'1', '2', '3'},
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-123",
		},
		"content-#@$": {
			b:        []byte{'#', '@', '$'},
			buf:      NewBuffer().WriteString("content-"),
			expected: "content-#@$",
		},
		"quoted-abc": {
			b:        []byte{'a', 'b', 'c'},
			buf:      NewBuffer().WithQuotes(),
			expected: `"abc"`,
		},
		"quoted-123": {
			b:        []byte{'1', '2', '3'},
			buf:      NewBuffer().WithQuotes(),
			expected: `"123"`,
		},
		"quoted-#@$": {
			b:        []byte{'#', '@', '$'},
			buf:      NewBuffer().WithQuotes(),
			expected: `"#@$"`,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			count, err := test.buf.Write(test.b)
			assert.Equal(t, len(test.b), count)
			require.NoError(t, err)

			assert.Equal(t, test.expected, test.buf.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_WriteTo(t *testing.T) {
	t.Parallel()

	type testCase struct {
		buf      *Buffer
		expected string
	}

	testCases := map[string]testCase{
		"empty": {
			buf:      NewBuffer(),
			expected: "",
		},
		"content": {
			buf:      NewBuffer().WriteString("content"),
			expected: "content",
		},
		"quoted-content": {
			buf:      NewBuffer().WithQuotes().WriteString("content"),
			expected: `"content"`,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			bytesBuffer := bytes.NewBuffer(nil)
			count, err := test.buf.WriteTo(bytesBuffer)
			assert.EqualValues(t, test.buf.Len(), count)
			require.NoError(t, err)

			assert.Equal(t, test.expected, bytesBuffer.String())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_Bytes(t *testing.T) {
	t.Parallel()

	buf := NewBuffer()
	buf.WriteString("content")
	buf.WriteInt64(25)
	buf.WriteBool(true)

	assert.Equal(t, []byte("content25true"), buf.Bytes())
}

func TestBuffer_Len(t *testing.T) {
	t.Parallel()

	type testCase struct {
		buf *Buffer
		len int
	}

	testCases := map[string]testCase{
		"empty": {
			buf: NewBuffer(),
			len: 0,
		},
		"content": {
			buf: NewBuffer().WriteString("content"),
			len: 7,
		},
		"cap-not-len": {
			buf: &Buffer{
				buf: make([]byte, 0, 1024),
			},
			len: 0,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.len, test.buf.Len())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_Cap(t *testing.T) {
	t.Parallel()

	type testCase struct {
		buf *Buffer
		cap int
	}

	testCases := map[string]testCase{
		"empty": {
			buf: NewBuffer(),
			cap: defaultSize,
		},
		"content": {
			buf: NewBuffer().WriteString("content"),
			cap: defaultSize,
		},
		"cap-not-len": {
			buf: &Buffer{
				buf: make([]byte, 512, 1024),
			},
			cap: 1024,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.cap, test.buf.Cap())

			assert.NotPanics(t, func() {
				test.buf.Dispose()
			})
		})
	}
}

func TestBuffer_Reset(t *testing.T) {
	t.Parallel()

	buf := NewBuffer()
	buf.WriteString("content")
	assert.Equal(t, "content", buf.String())

	buf.Reset()
	assert.Empty(t, buf.String())

	buf.WriteString("content2")
	assert.Equal(t, "content2", buf.String())

	assert.NotPanics(t, func() {
		buf.Dispose()
	})
}

func TestBuffer_CutSuffix(t *testing.T) {
	t.Parallel()

	type testCase struct {
		buf      *Buffer
		suffix   []byte
		expected string
	}

	testCases := map[string]testCase{
		"empty": {
			buf:      NewBuffer(),
			suffix:   []byte{},
			expected: "",
		},
		"no-suffix": {
			buf:      NewBuffer(),
			suffix:   []byte("suffix"),
			expected: "",
		},
		"large-suffix": {
			buf:      NewBuffer().WriteString("suf"),
			suffix:   []byte("suffix"),
			expected: "suf",
		},
		"partial-suffix": {
			buf:      NewBuffer().WriteString("suf_ix"),
			suffix:   []byte("suffix"),
			expected: "suf_ix",
		},
		"cut-suffix": {
			buf:      NewBuffer().WriteString("content-suffix"),
			suffix:   []byte("suffix"),
			expected: "content-",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			buf := test.buf
			defer buf.Dispose()

			buf = buf.CutSuffix(test.suffix)

			assert.Equal(t, test.expected, buf.String())
		})
	}
}

func BenchmarkBuffer(b *testing.B) {
	prepareBuf := func(buf *Buffer, quoted bool) *Buffer {
		if quoted {
			return buf.WithQuotes()
		}

		return buf
	}

	writeAll := func(buf *Buffer, quoted bool, count int) {
		for range count {
			prepareBuf(buf, quoted).WriteByte('b')
			prepareBuf(buf, quoted).WriteString("string")
			prepareBuf(buf, quoted).WriteInt64(-10)
			prepareBuf(buf, quoted).WriteUint64(10)
			prepareBuf(buf, quoted).WriteFloat64(7.5, 64)
			prepareBuf(buf, quoted).WriteBool(true)
			prepareBuf(buf, quoted).WriteTime(time.Now(), time.RFC3339)
		}
	}

	const (
		multiCount = 1_000
	)

	type benchData struct {
		quoted bool
		count  int
	}

	benchmarks := map[string]benchData{
		"single-default": {
			count: 1,
		},
		"single-quoted": {
			quoted: true,
			count:  1,
		},
		"multi-default": {
			count: multiCount,
		},
		"multi-quoted": {
			quoted: true,
			count:  multiCount,
		},
	}

	for name, bench := range benchmarks {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()

			for b.Loop() {
				buf := NewBuffer()
				writeAll(buf, bench.quoted, bench.count)

				buf.Dispose()
			}
		})
	}
}
