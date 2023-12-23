package test

import (
	"bytes"
	"math"
	"reflect"
	"testing"
	"yu/golang/internal/app"

	. "github.com/franela/goblin"
)

// go test ./test/unit/types...
// go test -run TestTypes ./test/unit/types

func TestTypes(t *testing.T) {
	g := Goblin(t)

	g.Describe("Convert types within ranges", func() {
		g.It("Uint8ToInt8Range", func() {

			if res := app.Uint8ToInt8Range(app.MIN_UINT8); res != app.MIN_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_INT8, res)
			}

			if res := app.Uint8ToInt8Range(app.MAX_UINT8); res != app.MAX_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MAX_INT8, res)
			}
		})

		g.It("Uint8ToInt16Range", func() {

			if res := app.Uint8ToInt16Range(app.MIN_UINT8); res != int16(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(app.MIN_INT8), res)
			}

			if res := app.Uint8ToInt16Range(app.MAX_UINT8); res != int16(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(app.MAX_INT8), res)
			}
		})

		g.It("Uint8ToInt32Range", func() {

			if res := app.Uint8ToInt32Range(app.MIN_UINT8); res != int32(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(app.MIN_INT8), res)
			}

			if res := app.Uint8ToInt32Range(app.MAX_UINT8); res != int32(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(app.MAX_INT8), res)
			}
		})

		g.It("Uint8ToInt64Range", func() {

			if res := app.Uint8ToInt64Range(app.MIN_UINT8); res != int64(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(app.MIN_INT8), res)
			}

			if res := app.Uint8ToInt64Range(app.MAX_UINT8); res != int64(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(app.MAX_INT8), res)
			}
		})

		g.It("Uint8ToIntRange", func() {

			if res := app.Uint8ToIntRange(app.MIN_UINT8); res != int(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(app.MIN_INT8), res)
			}

			if res := app.Uint8ToIntRange(app.MAX_UINT8); res != int(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(app.MAX_INT8), res)
			}
		})

		g.It("Int8ToUint8Range", func() {

			if res := app.Int8ToUint8Range(app.MIN_INT8); res != app.MIN_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT8, res)
			}

			if res := app.Int8ToUint8Range(app.MAX_INT8); res != app.MAX_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MAX_UINT8, res)
			}
		})

		g.It("Int8ToUint16Range", func() {

			if res := app.Int8ToUint16Range(app.MIN_INT8); res != uint16(app.MIN_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(app.MIN_UINT8), res)
			}

			if res := app.Int8ToUint16Range(app.MAX_INT8); res != uint16(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(app.MAX_UINT8), res)
			}
		})

		g.It("Int8ToUint32Range", func() {

			if res := app.Int8ToUint32Range(app.MIN_INT8); res != app.MIN_UINT32 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT32, res)
			}

			if res := app.Int8ToUint32Range(app.MAX_INT8); res != uint32(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint32(app.MAX_UINT8), res)
			}
		})

		g.It("Int8ToUint64Range", func() {

			if res := app.Int8ToUint64Range(app.MIN_INT8); res != app.MIN_UINT64 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT64, res)
			}

			if res := app.Int8ToUint64Range(app.MAX_INT8); res != uint64(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint64(app.MAX_UINT8), res)
			}
		})

		g.It("Int8ToUintRange", func() {

			if res := app.Int8ToUintRange(app.MIN_INT8); res != app.MIN_UINT {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT, res)
			}

			if res := app.Int8ToUintRange(app.MAX_INT8); res != uint(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint(app.MAX_UINT8), res)
			}
		})
	})

	g.Describe("Wrong Types", func() {
		g.Describe("Int", func() {
			g.Xit("MinInt8", func() {
				const expect = "int8"
				if ref := reflect.ValueOf(math.MinInt8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxInt8", func() {
				const expect = "int8"
				if ref := reflect.ValueOf(math.MaxInt8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinInt16", func() {
				const expect = "int16"
				if ref := reflect.ValueOf(math.MinInt16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxInt16", func() {
				const expect = "int16"
				if ref := reflect.ValueOf(math.MaxInt16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinInt32", func() {
				const expect = "int32"
				if ref := reflect.ValueOf(math.MinInt32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxInt32", func() {
				const expect = "int32"
				if ref := reflect.ValueOf(math.MaxInt32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinInt64", func() {
				const expect = "int64"
				if ref := reflect.ValueOf(math.MinInt64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxInt64", func() {
				const expect = "int64"
				if ref := reflect.ValueOf(math.MaxInt64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinInt", func() {
				const expect = "int"
				if ref := reflect.ValueOf(math.MinInt).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxInt", func() {
				const expect = "int"
				if ref := reflect.ValueOf(math.MaxInt).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})

		g.Describe("Uint", func() {

			g.Xit("MinUint8", func() {
				const expect = "uint8"
				if ref := reflect.ValueOf(app.MIN_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxUint8", func() {
				const expect = "uint8"
				if ref := reflect.ValueOf(math.MaxUint8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinUint16", func() {
				const expect = "uint16"
				if ref := reflect.ValueOf(app.MIN_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxUint16", func() {
				const expect = "uint16"
				if ref := reflect.ValueOf(math.MaxUint16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinUint32", func() {
				const expect = "uint32"
				if ref := reflect.ValueOf(app.MIN_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxUint32", func() {
				const expect = "uint32"
				if ref := reflect.ValueOf(math.MaxUint32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinUint64", func() {
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxUint64", func() {
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MinUint", func() {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.Xit("MaxUint", func() {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MAX_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})
	})

	g.Describe("Right Types", func() {
		g.Describe("Int", func() {
			g.It("MinInt8", func() {
				const expect = "int8"
				if ref := reflect.ValueOf(app.MIN_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxInt8", func() {
				const expect = "int8"
				if ref := reflect.ValueOf(app.MAX_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinInt16", func() {
				const expect = "int16"
				if ref := reflect.ValueOf(app.MIN_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxInt16", func() {
				const expect = "int16"
				if ref := reflect.ValueOf(app.MAX_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinInt32", func() {
				const expect = "int32"
				if ref := reflect.ValueOf(app.MIN_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxInt32", func() {
				const expect = "int32"
				if ref := reflect.ValueOf(app.MAX_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinInt64", func() {
				const expect = "int64"
				if ref := reflect.ValueOf(app.MIN_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxInt64", func() {
				const expect = "int64"
				if ref := reflect.ValueOf(app.MAX_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinInt", func() {
				const expect = "int"
				if ref := reflect.ValueOf(app.MIN_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxInt", func() {
				const expect = "int"
				if ref := reflect.ValueOf(app.MAX_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})

		g.Describe("Uint", func() {

			g.It("MinUint8", func() {
				const expect = "uint8"
				if ref := reflect.ValueOf(app.MIN_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxUint8", func() {
				const expect = "uint8"
				if ref := reflect.ValueOf(app.MAX_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinUint16", func() {
				const expect = "uint16"
				if ref := reflect.ValueOf(app.MIN_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxUint16", func() {
				const expect = "uint16"
				if ref := reflect.ValueOf(app.MAX_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinUint32", func() {
				const expect = "uint32"
				if ref := reflect.ValueOf(app.MIN_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxUint32", func() {
				const expect = "uint32"
				if ref := reflect.ValueOf(app.MAX_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinUint64", func() {
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxUint64", func() {
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MinUint", func() {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			g.It("MaxUint", func() {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MAX_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})
	})

	g.Describe("isNumeric", func() {

		g.It("int8", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_INT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int16", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_INT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int32", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_INT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int64", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_INT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_INT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint8", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint16", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint32", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint64", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("byte", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_BYTE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("rune", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_RUNE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("float32", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_FLOAT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("float64", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_FLOAT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("complex64", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_COMPLEX64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("complex128", func() {
			const expect = true

			if res := app.IsNumeric(app.MIN_COMPLEX128); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("bool", func() {
			const expect = false

			if res := app.IsNumeric(false); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}

			if res := app.IsNumeric(true); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("string", func() {
			const expect = false

			if res := app.IsNumeric(""); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("struct", func() {
			const expect = false

			if res := app.IsNumeric(struct{}{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("slice", func() {
			const expect = false

			if res := app.IsNumeric([]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("map", func() {
			const expect = false

			if res := app.IsNumeric(map[any]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("buffer", func() {
			const expect = false

			if res := app.IsNumeric(new(bytes.Buffer)); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("function", func() {
			const expect = false

			if res := app.IsNumeric(func() {}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	g.Describe("isPrimitive", func() {

		g.It("int8", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int16", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int32", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int64", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("int", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint8", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint16", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint32", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint64", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("uint", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("byte", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_BYTE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("rune", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_RUNE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("float32", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_FLOAT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("float64", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_FLOAT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("complex64", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_COMPLEX64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("complex128", func() {
			const expect = true

			if res := app.IsPrimitive(app.MIN_COMPLEX128); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("bool", func() {
			const expect = true

			if res := app.IsPrimitive(false); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}

			if res := app.IsPrimitive(true); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("string", func() {
			const expect = true

			if res := app.IsPrimitive(""); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("struct", func() {
			const expect = false

			if res := app.IsPrimitive(struct{}{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("slice", func() {
			const expect = false

			if res := app.IsPrimitive([]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("map", func() {
			const expect = false

			if res := app.IsPrimitive(map[any]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("buffer", func() {
			const expect = false

			if res := app.IsPrimitive(new(bytes.Buffer)); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		g.It("function", func() {
			const expect = false

			if res := app.IsPrimitive(func() {}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})
}
