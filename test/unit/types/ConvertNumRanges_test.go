package test

import (
	"testing"
	"yu/golang/internal/app"

	. "github.com/franela/goblin"
)

// go test ./test/unit/types...
// go test -v -run TestConvertNumRanges ./test/unit/types

func TestConvertNumRanges(t *testing.T) {
	g := Goblin(t)

	// Converting number within its range
	g.Describe("Uint to Int", func() {
		g.It("uint8  to int8:  [0..255] => [-128..127]", func() {
			g.Assert(app.Uint8ToInt8Range(app.MIN_UINT8)).Equal(app.MIN_INT8)
			g.Assert(app.Uint8ToInt8Range(app.MAX_UINT8)).Equal(app.MAX_INT8)
		})

		g.It("uint8  to int16: [0..255] => [-128..127]", func() {
			g.Assert(app.Uint8ToInt16Range(app.MIN_UINT8)).Equal(int16(app.MIN_INT8))
			g.Assert(app.Uint8ToInt16Range(app.MAX_UINT8)).Equal(int16(app.MAX_INT8))
		})

		g.It("uint8  to int32: [0..255] => [-128..127]", func() {
			g.Assert(app.Uint8ToInt32Range(app.MIN_UINT8)).Equal(int32(app.MIN_INT8))
			g.Assert(app.Uint8ToInt32Range(app.MAX_UINT8)).Equal(int32(app.MAX_INT8))
		})

		g.It("uint8  to int64: [0..255] => [-128..127]", func() {
			g.Assert(app.Uint8ToInt64Range(app.MIN_UINT8)).Equal(int64(app.MIN_INT8))
			g.Assert(app.Uint8ToInt64Range(app.MAX_UINT8)).Equal(int64(app.MAX_INT8))
		})

		g.It("uint8  to int:   [0..255] => [-128..127]", func() {
			g.Assert(app.Uint8ToIntRange(app.MIN_UINT8)).Equal(int(app.MIN_INT8))
			g.Assert(app.Uint8ToIntRange(app.MAX_UINT8)).Equal(int(app.MAX_INT8))
		})

		// uint16 to intX

		g.Xit("uint16 to int8:  [0..65535] => [-128..127]", func() {})
		g.Xit("uint16 to int16: [0..65535] => [-32768..32767]", func() {})
		g.Xit("uint16 to int32: [0..65535] => [-32768..32767]", func() {})
		g.Xit("uint16 to int64: [0..65535] => [-32768..32767]", func() {})
		g.Xit("uint16 to int:   [0..65535] => [-32768..32767]", func() {})

		// uint32 to intX

		g.Xit("uint32 to int8:  [0..4294967295] => [-128..127]", func() {})
		g.Xit("uint32 to int16: [0..4294967295] => [-32768..32767]", func() {})
		g.Xit("uint32 to int32: [0..4294967295] => [-2147483648..2147483647]", func() {})
		g.Xit("uint32 to int64: [0..4294967295] => [-2147483648..2147483647]", func() {})
		g.Xit("uint32 to int:   [0..4294967295] => [-2147483648..2147483647]", func() {})

		// uint64 to intX

		g.Xit("uint64 to int8:  [0..18446744073709551615] => [-128..127]", func() {})
		g.Xit("uint64 to int16: [0..18446744073709551615] => [-32768..32767]", func() {})
		g.Xit("uint64 to int32: [0..18446744073709551615] => [-2147483648..2147483647]", func() {})
		g.Xit("uint64 to int64: [0..18446744073709551615] => [-9223372036854775808..9223372036854775807]", func() {})
		g.Xit("uint64 to int:   [0..18446744073709551615] => [-9223372036854775808..9223372036854775807]", func() {})
	})

	g.Describe("Int To Uint", func() {
		g.It("int8  to uint8:  [-128..127] => [0..255]", func() {
			g.Assert(app.Int8ToUint8Range(app.MIN_INT8)).Equal(app.MIN_UINT8)
			g.Assert(app.Int8ToUint8Range(app.MAX_INT8)).Equal(app.MAX_UINT8)
		})

		g.It("int8  to uint16: [-128..127] => [0..255]", func() {
			g.Assert(app.Int8ToUint16Range(app.MIN_INT8)).Equal(uint16(app.MIN_UINT8))
			g.Assert(app.Int8ToUint16Range(app.MAX_INT8)).Equal(uint16(app.MAX_UINT8))
		})

		g.It("int8  to uint32: [-128..127] => [0..255]", func() {
			g.Assert(app.Int8ToUint32Range(app.MIN_INT8)).Equal(app.MIN_UINT32)
			g.Assert(app.Int8ToUint32Range(app.MAX_INT8)).Equal(uint32(app.MAX_UINT8))
		})

		g.It("int8  to uint64: [-128..127] => [0..255]", func() {
			g.Assert(app.Int8ToUint64Range(app.MIN_INT8)).Equal(app.MIN_UINT64)
			g.Assert(app.Int8ToUint64Range(app.MAX_INT8)).Equal(uint64(app.MAX_UINT8))
		})

		g.It("int8  to uint:   [-128..127] => [0..255]", func() {
			g.Assert(app.Int8ToUintRange(app.MIN_INT8)).Equal(app.MIN_UINT)
			g.Assert(app.Int8ToUintRange(app.MAX_INT8)).Equal(uint(app.MAX_UINT8))
		})

		// int16 to uintX

		g.Xit("int16 to uint8:  [-32768..32767] => [0..255]", func() {})
		g.Xit("int16 to uint16: [-32768..32767] => [0..65535]", func() {})
		g.Xit("int16 to uint32: [-32768..32767] => [0..65535]", func() {})
		g.Xit("int16 to uint64: [-32768..32767] => [0..65535]", func() {})
		g.Xit("int16 to uint:   [-32768..32767] => [0..65535]", func() {})

		// uint32 to intX

		g.Xit("int32 to uint8:  [-2147483648..2147483647] => [0..255]", func() {})
		g.Xit("int32 to uint16: [-2147483648..2147483647] => [0..65535]", func() {})
		g.Xit("int32 to uint32: [-2147483648..2147483647] => [0..4294967295]", func() {})
		g.Xit("int32 to uint64: [-2147483648..2147483647] => [0..4294967295]", func() {})
		g.Xit("int32 to uint:   [-2147483648..2147483647] => [0..4294967295]", func() {})

		// uint64 to intX

		g.Xit("int64 to uint8:  [-9223372036854775808..9223372036854775807] => [0..255]", func() {})
		g.Xit("int64 to uint16: [-9223372036854775808..9223372036854775807] => [0..65535]", func() {})
		g.Xit("int64 to uint32: [-9223372036854775808..9223372036854775807] => [0..4294967295]", func() {})
		g.Xit("int64 to uint64: [-9223372036854775808..9223372036854775807] => [0..18446744073709551615]", func() {})
		g.Xit("int64 to uint:   [-9223372036854775808..9223372036854775807] => [0..18446744073709551615]", func() {})
	})
}
