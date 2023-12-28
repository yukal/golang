package validation

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test ./pkg/validation/...
// go test -v -run TestCompare ./pkg/validation/...

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

				result := compare(NON_ZERO, proto, value)
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

				result := compare(NON_ZERO, proto, value)
				g.Assert(result).Equal(MsgEmpty)
			})
		}

		g.It("failure if given an empty array", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new([0]string))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgEmpty)
		})

		g.It("failure if given an empty slice", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new([]string))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgEmpty)
		})

		g.It("failure if given an empty map", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(map[string]string))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgEmpty)
		})

		g.It("failure if given an empty chan", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(chan string))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgEmpty)
		})

		g.It("failure if given an empty struct", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(struct{}))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgEmpty)
		})

		g.It("failure if given an empty func", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf(*new(func()))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgEmpty)
		})

		g.It("failure if given a nil value", func() {
			proto := reflect.ValueOf(*new(interface{}))
			value := reflect.ValueOf(*new(interface{}))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgInvalidValue)
		})
	})

	g.Describe(`Rule "eq" (equal)`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value equals the expected number", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the expected number", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must be exactly 4")
			})

			g.It("failure when the value is greater than the expected number", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must be exactly 4")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length matches a filled array", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty array", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 items")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(arrFilled)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length matches a filled slice", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty slice", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 items")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(sliceFilled)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length matches a filled map", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty map", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 items")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(mapFilled)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length matches a filled string", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length matches an empty string", func() {
				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than expected", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 characters")
			})

			g.It("failure when the length is greater than expected", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(strFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 characters")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(strFilled)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})

			g.It("failure when given empty args: (nil, 1)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})

	g.Describe(`Rule "min"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value exceeds the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the value reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must be at least 4")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(strFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(strFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 characters")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(strFilled)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(arrFilled)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(sliceFilled)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length is greater than the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(mapFilled)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})

			g.It("failure when given empty args: (nil, 1)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})

	g.Describe(`Rule "max"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value is less than the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the value reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value exceeds the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must be up to 4")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(arrFilled)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(sliceFilled)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(mapFilled)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(strFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(strFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 characters")
			})

			g.It("failure when a nil value was passed", func() {
				proto := reflect.ValueOf(strFilled)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})

			g.It("failure when given empty args: (nil, 1)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})

	g.Describe(`Rule "range"`, func() {
		g.Describe(`numeric`, func() {
			g.It("success when the value matches the range", func() {
				proto := reflect.ValueOf(Range{1, 25})
				value := reflect.ValueOf(15)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when given below-range value", func() {
				proto := reflect.ValueOf(Range{15, 25})
				value := reflect.ValueOf(5)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must be in the range 15..25")
			})

			g.It("failure when given above-range value", func() {
				proto := reflect.ValueOf(Range{15, 25})
				value := reflect.ValueOf(55)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must be in the range 15..25")
			})

			g.It("failure if at least 1 element of the range is empty", func() {
				value := reflect.ValueOf(5)
				protos := []Range{
					{},
					{nil, 15},
					{15, nil},
				}

				for _, item := range protos {
					proto := reflect.ValueOf(item)
					result := compare("range", proto, value)

					g.Assert(result).Equal(MsgInvalidRangeVal)
				}
			})

			g.It("failure when given a nil value", func() {
				proto := reflect.ValueOf(Range{1, 25})
				value := reflect.ValueOf(nil)

				result := compare("range", proto, value)
				g.Assert(result).Equal("invalid value")
			})
		})

		// ...

		g.Describe(`array`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(arrFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(arrFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(arrFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 items")
			})

			g.It("failure if at least 1 element of the range is empty", func() {
				value := reflect.ValueOf(arrFilled)
				protos := []Range{
					{},
					{nil, 4},
					{4, nil},
				}

				for _, item := range protos {
					proto := reflect.ValueOf(item)
					result := compare("range", proto, value)

					g.Assert(result).Equal(MsgInvalidRangeVal)
				}
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(sliceFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(sliceFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(sliceFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 items")
			})

			g.It("failure if at least 1 element of the range is empty", func() {
				value := reflect.ValueOf(sliceFilled)
				protos := []Range{
					{},
					{nil, 4},
					{4, nil},
				}

				for _, item := range protos {
					proto := reflect.ValueOf(item)
					result := compare("range", proto, value)

					g.Assert(result).Equal(MsgInvalidRangeVal)
				}
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(mapFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(mapFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(mapFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 items")
			})

			g.It("failure if at least 1 element of the range is empty", func() {
				value := reflect.ValueOf(mapFilled)
				protos := []Range{
					{},
					{nil, 4},
					{4, nil},
				}

				for _, item := range protos {
					proto := reflect.ValueOf(item)
					result := compare("range", proto, value)

					g.Assert(result).Equal(MsgInvalidRangeVal)
				}
			})
		})

		// ...

		g.Describe(`string`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 characters")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 characters")
			})

			g.It("failure if at least 1 element of the range is empty", func() {
				value := reflect.ValueOf(strFilled)
				protos := []Range{
					{},
					{nil, 4},
					{4, nil},
				}

				for _, item := range protos {
					proto := reflect.ValueOf(item)
					result := compare("range", proto, value)

					g.Assert(result).Equal(MsgInvalidRangeVal)
				}
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure when given empty args: (nil, nil)", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				result := compare("range", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})

	g.Describe(`Rule "year"`, func() {
		g.It("success when the value matches a specific year", func() {
			tm, err := time.Parse(time.RFC3339, "2023-12-25T16:04:05Z")
			g.Assert(err).IsNil(err)

			proto := reflect.ValueOf(2023)
			value := reflect.ValueOf(tm)

			result := compare("year", proto, value)
			g.Assert(result).Equal("")
		})

		g.It("failure when the value is not match", func() {
			proto := reflect.ValueOf(2024)
			value := reflect.ValueOf(*new(time.Time))

			result := compare("year", proto, value)
			g.Assert(result).Equal("must be exactly 2024")
		})
	})

}