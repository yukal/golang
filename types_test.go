package main

import (
	"math"
	"reflect"
	"testing"
)

func TestTypes(t *testing.T) {
	t.Run("Convert types within ranges", func(t *testing.T) {
		t.Run("Uint8ToInt8Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Uint8ToInt8Range(MIN_UINT8); res != MIN_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MIN_INT8, res)
			}
	
			if res := Uint8ToInt8Range(MAX_UINT8); res != MAX_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MAX_INT8, res)
			}
		})
	
		t.Run("Uint8ToInt16Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Uint8ToInt16Range(MIN_UINT8); res != int16(MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(MIN_INT8), res)
			}
	
			if res := Uint8ToInt16Range(MAX_UINT8); res != int16(MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(MAX_INT8), res)
			}
		})
	
		t.Run("Uint8ToInt32Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Uint8ToInt32Range(MIN_UINT8); res != int32(MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(MIN_INT8), res)
			}
	
			if res := Uint8ToInt32Range(MAX_UINT8); res != int32(MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(MAX_INT8), res)
			}
		})
	
		t.Run("Uint8ToInt64Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Uint8ToInt64Range(MIN_UINT8); res != int64(MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(MIN_INT8), res)
			}
	
			if res := Uint8ToInt64Range(MAX_UINT8); res != int64(MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(MAX_INT8), res)
			}
		})
	
		t.Run("Uint8ToIntRange", func(t *testing.T) {
			t.Parallel()
	
			if res := Uint8ToIntRange(MIN_UINT8); res != int(MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(MIN_INT8), res)
			}
	
			if res := Uint8ToIntRange(MAX_UINT8); res != int(MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(MAX_INT8), res)
			}
		})
	
		t.Run("Int8ToUint8Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Int8ToUint8Range(MIN_INT8); res != MIN_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MIN_UINT8, res)
			}
	
			if res := Int8ToUint8Range(MAX_INT8); res != MAX_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MAX_UINT8, res)
			}
		})
	
		t.Run("Int8ToUint16Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Int8ToUint16Range(MIN_INT8); res != uint16(MIN_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(MIN_UINT8), res)
			}
	
			if res := Int8ToUint16Range(MAX_INT8); res != uint16(MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(MAX_UINT8), res)
			}
		})
	
		t.Run("Int8ToUint32Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Int8ToUint32Range(MIN_INT8); res != MIN_UINT32 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MIN_UINT32, res)
			}
	
			if res := Int8ToUint32Range(MAX_INT8); res != uint32(MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint32(MAX_UINT8), res)
			}
		})
	
		t.Run("Int8ToUint64Range", func(t *testing.T) {
			t.Parallel()
	
			if res := Int8ToUint64Range(MIN_INT8); res != MIN_UINT64 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MIN_UINT64, res)
			}
	
			if res := Int8ToUint64Range(MAX_INT8); res != uint64(MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint64(MAX_UINT8), res)
			}
		})
	
		t.Run("Int8ToUintRange", func(t *testing.T) {
			t.Parallel()
	
			if res := Int8ToUintRange(MIN_INT8); res != MIN_UINT {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", MIN_UINT, res)
			}
	
			if res := Int8ToUintRange(MAX_INT8); res != uint(MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint(MAX_UINT8), res)
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
				if ref := reflect.ValueOf(MIN_UINT8).Kind(); ref != reflect.Uint8 {
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
				if ref := reflect.ValueOf(MIN_UINT16).Kind(); ref != reflect.Uint16 {
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
				if ref := reflect.ValueOf(MIN_UINT32).Kind(); ref != reflect.Uint32 {
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
				if ref := reflect.ValueOf(MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(MAX_UINT).Kind(); ref != reflect.Uint {
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
				if ref := reflect.ValueOf(MIN_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt8", func(t *testing.T) {
				t.Parallel()
				const expect = "int8"
				if ref := reflect.ValueOf(MAX_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(MIN_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(MAX_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(MIN_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(MAX_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(MIN_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(MAX_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(MIN_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(MAX_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})

		t.Run("Uint", func(t *testing.T) {
			t.Parallel()

			t.Run("MinUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(MIN_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(MAX_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(MIN_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(MAX_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(MIN_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(MAX_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(MAX_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})
	})
}
