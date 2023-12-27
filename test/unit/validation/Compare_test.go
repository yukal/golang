package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
)

// go test ./test/unit/validation/...
// go test -v -run TestCompare ./test/unit/validation/...

func TestCompare(t *testing.T) {
	var (
		strEmpty    = ""
		strFilled   = "love"
		arrEmpty    = [0]string{}
		arrFilled   = [4]string{"c", "o", "d", "e"}
		sliceEmpty  = []string{}
		sliceFilled = []string{"t", "e", "s", "t"}
		mapEmpty    = map[string]string{}
		mapFilled   = map[string]string{
			"i": "val1",
			"t": "val2",
			"e": "val3",
			"m": "val4",
		}
	)

	g := Goblin(t)

	g.Describe(`Rule NON_ZERO`, func() {
		nonEmptyValues := []any{
			int8(-1),
			int16(-1),
			int32(-1),
			int64(-1),
			int(-1),
			uint8(1),
			uint16(1),
			uint32(1),
			uint64(1),
			uint(1),
			float32(1),
			float64(-1),
			complex64(1),
			complex128(-1),
			true,
			"ok",
			time.Now(),
		}

		for _, val := range nonEmptyValues {
			val := val
			title := fmt.Sprintf("success if given a non-empty %T", val)

			g.It(title, func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(val)

				result := validation.Compare(validation.NON_ZERO, proto, value)
				g.Assert(result).Equal("")
			})
		}

		emptyValues := []any{
			*new(int8),       // 0
			*new(int16),      // 0
			*new(int32),      // 0
			*new(int64),      // 0
			*new(int),        // 0
			*new(uint8),      // 0
			*new(uint16),     // 0
			*new(uint32),     // 0
			*new(uint64),     // 0
			*new(uint),       // 0
			*new(float32),    // 0
			*new(float64),    // 0
			*new(complex64),  // 0
			*new(complex128), // 0
			*new(bool),       // false
			*new(string),     // ""
		}

		for _, val := range emptyValues {
			val := val
			title := fmt.Sprintf("failure if given an empty %T", val)

			g.It(title, func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(val)

				result := validation.Compare(validation.NON_ZERO, proto, value)
				g.Assert(result).Equal(validation.MsgEmpty)
			})
		}

		g.It("failure if given an empty array", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new([0]string))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgEmpty)
		})

		g.It("failure if given an empty slice", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new([]string))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgEmpty)
		})

		g.It("failure if given an empty map", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(map[string]string))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgEmpty)
		})

		g.It("failure if given an empty chan", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(chan string))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgEmpty)
		})

		g.It("failure if given an empty struct", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(struct{}))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgEmpty)
		})

		g.It("failure if given an empty func", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(func()))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgEmpty)
		})

		g.It("failure if given a nil value", func() {
			proto := reflect.ValueOf(*new(interface{}))
			value := reflect.ValueOf(*new(interface{}))

			result := validation.Compare(validation.NON_ZERO, proto, value)
			g.Assert(result).Equal(validation.MsgInvalidValue)
		})
	})

	g.Describe(`Rule "eq" (equal)`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value equals the expected number", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the expected number", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must be exactly 4")
			})

			g.It("failure when the value is greater than the expected number", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must be exactly 4")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length matches a filled array", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty array", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 items")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(arrFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length matches a filled slice", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty slice", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 items")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(sliceFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length matches a filled map", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty map", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 items")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(mapFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length matches a filled string", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty string", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 characters")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 characters")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(strFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})

			g.It("failure when given empty args: (nil, 1)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				result := validation.Compare("eq", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})
	})

	g.Describe(`Rule "min"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value exceeds the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the value reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("must be at least 4")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 characters")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(strFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(arrFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(sliceFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(mapFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})

			g.It("failure when given empty args: (nil, 1)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				result := validation.Compare("min", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})
	})

	g.Describe(`Rule "max"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value is less than the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the value reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value exceeds the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("must be up to 4")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(arrFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(sliceFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(mapFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(strFilled)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 characters")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(strFilled)
				value := reflect.ValueOf(nil)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})

			g.It("failure when given empty args: (nil, 1)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				result := validation.Compare("max", proto, value)
				g.Assert(result).Equal(validation.MsgInvalidValue)
			})
		})
	})

}
