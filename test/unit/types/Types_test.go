package test

import (
	"bytes"
	"math"
	"reflect"
	"testing"
	"yu/golang/internal/app"
)

// go test ./test/unit/types...
// go test -run TestTypes ./test/unit/types

func TestTypes(t *testing.T) {
	t.Run("Convert types within ranges", func(t *testing.T) {
		t.Run("Uint8ToInt8Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Uint8ToInt8Range(app.MIN_UINT8); res != app.MIN_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_INT8, res)
			}

			if res := app.Uint8ToInt8Range(app.MAX_UINT8); res != app.MAX_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MAX_INT8, res)
			}
		})

		t.Run("Uint8ToInt16Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Uint8ToInt16Range(app.MIN_UINT8); res != int16(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(app.MIN_INT8), res)
			}

			if res := app.Uint8ToInt16Range(app.MAX_UINT8); res != int16(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(app.MAX_INT8), res)
			}
		})

		t.Run("Uint8ToInt32Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Uint8ToInt32Range(app.MIN_UINT8); res != int32(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(app.MIN_INT8), res)
			}

			if res := app.Uint8ToInt32Range(app.MAX_UINT8); res != int32(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(app.MAX_INT8), res)
			}
		})

		t.Run("Uint8ToInt64Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Uint8ToInt64Range(app.MIN_UINT8); res != int64(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(app.MIN_INT8), res)
			}

			if res := app.Uint8ToInt64Range(app.MAX_UINT8); res != int64(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(app.MAX_INT8), res)
			}
		})

		t.Run("Uint8ToIntRange", func(t *testing.T) {
			t.Parallel()

			if res := app.Uint8ToIntRange(app.MIN_UINT8); res != int(app.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(app.MIN_INT8), res)
			}

			if res := app.Uint8ToIntRange(app.MAX_UINT8); res != int(app.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(app.MAX_INT8), res)
			}
		})

		t.Run("Int8ToUint8Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Int8ToUint8Range(app.MIN_INT8); res != app.MIN_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT8, res)
			}

			if res := app.Int8ToUint8Range(app.MAX_INT8); res != app.MAX_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MAX_UINT8, res)
			}
		})

		t.Run("Int8ToUint16Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Int8ToUint16Range(app.MIN_INT8); res != uint16(app.MIN_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(app.MIN_UINT8), res)
			}

			if res := app.Int8ToUint16Range(app.MAX_INT8); res != uint16(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(app.MAX_UINT8), res)
			}
		})

		t.Run("Int8ToUint32Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Int8ToUint32Range(app.MIN_INT8); res != app.MIN_UINT32 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT32, res)
			}

			if res := app.Int8ToUint32Range(app.MAX_INT8); res != uint32(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint32(app.MAX_UINT8), res)
			}
		})

		t.Run("Int8ToUint64Range", func(t *testing.T) {
			t.Parallel()

			if res := app.Int8ToUint64Range(app.MIN_INT8); res != app.MIN_UINT64 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT64, res)
			}

			if res := app.Int8ToUint64Range(app.MAX_INT8); res != uint64(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint64(app.MAX_UINT8), res)
			}
		})

		t.Run("Int8ToUintRange", func(t *testing.T) {
			t.Parallel()

			if res := app.Int8ToUintRange(app.MIN_INT8); res != app.MIN_UINT {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", app.MIN_UINT, res)
			}

			if res := app.Int8ToUintRange(app.MAX_INT8); res != uint(app.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint(app.MAX_UINT8), res)
			}
		})
	})

	t.Run("Wrong Types", func(t *testing.T) {
		t.Skip()

		t.Run("Int", func(t *testing.T) {
			t.Run("MinInt8", func(t *testing.T) {
				t.Parallel()
				const expect = "int8"
				if ref := reflect.ValueOf(math.MinInt8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt8", func(t *testing.T) {
				t.Parallel()
				const expect = "int8"
				if ref := reflect.ValueOf(math.MaxInt8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(math.MinInt16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(math.MaxInt16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(math.MinInt32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(math.MaxInt32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(math.MinInt64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(math.MaxInt64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(math.MinInt).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(math.MaxInt).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})

		t.Run("Uint", func(t *testing.T) {
			t.Parallel()

			t.Run("MinUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(app.MIN_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(math.MaxUint8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(app.MIN_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(math.MaxUint16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(app.MIN_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(math.MaxUint32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MAX_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})
	})

	t.Run("Right Types", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			t.Run("MinInt8", func(t *testing.T) {
				t.Parallel()
				const expect = "int8"
				if ref := reflect.ValueOf(app.MIN_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt8", func(t *testing.T) {
				t.Parallel()
				const expect = "int8"
				if ref := reflect.ValueOf(app.MAX_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(app.MIN_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(app.MAX_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(app.MIN_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(app.MAX_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(app.MIN_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(app.MAX_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(app.MIN_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(app.MAX_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})

		t.Run("Uint", func(t *testing.T) {
			t.Parallel()

			t.Run("MinUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(app.MIN_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(app.MAX_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(app.MIN_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(app.MAX_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(app.MIN_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(app.MAX_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(app.MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(app.MAX_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})
	})

	t.Run("isNumeric", func(t *testing.T) {

		t.Run("int8", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_INT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int16", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_INT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int32", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_INT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_INT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_INT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint8", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint16", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint32", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_UINT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("byte", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_BYTE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("rune", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_RUNE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("float32", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_FLOAT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("float64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_FLOAT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("complex64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_COMPLEX64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("complex128", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsNumeric(app.MIN_COMPLEX128); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("bool", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric(false); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}

			if res := app.IsNumeric(true); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("string", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric(""); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("struct", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric(struct{}{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("slice", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric([]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("map", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric(map[any]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("buffer", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric(new(bytes.Buffer)); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("function", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsNumeric(func() {}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	t.Run("isPrimitive", func(t *testing.T) {

		t.Run("int8", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int16", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int32", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("int", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_INT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint8", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT8); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint16", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT16); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint32", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("uint", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_UINT); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("byte", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_BYTE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("rune", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_RUNE); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("float32", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_FLOAT32); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("float64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_FLOAT64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("complex64", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_COMPLEX64); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("complex128", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(app.MIN_COMPLEX128); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("bool", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(false); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}

			if res := app.IsPrimitive(true); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("string", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := app.IsPrimitive(""); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("struct", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsPrimitive(struct{}{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("slice", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsPrimitive([]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("map", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsPrimitive(map[any]any{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("buffer", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsPrimitive(new(bytes.Buffer)); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("function", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := app.IsPrimitive(func() {}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})
}
