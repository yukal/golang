package test

import (
	"math"
	"reflect"
	"testing"
	"yu/golang/src"
)

func TestTypes(t *testing.T) {
	t.Run("Convert types within ranges", func(t *testing.T) {
		t.Run("Uint8ToInt8Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Uint8ToInt8Range(src.MIN_UINT8); res != src.MIN_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MIN_INT8, res)
			}
	
			if res := src.Uint8ToInt8Range(src.MAX_UINT8); res != src.MAX_INT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MAX_INT8, res)
			}
		})
	
		t.Run("Uint8ToInt16Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Uint8ToInt16Range(src.MIN_UINT8); res != int16(src.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(src.MIN_INT8), res)
			}
	
			if res := src.Uint8ToInt16Range(src.MAX_UINT8); res != int16(src.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int16(src.MAX_INT8), res)
			}
		})
	
		t.Run("Uint8ToInt32Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Uint8ToInt32Range(src.MIN_UINT8); res != int32(src.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(src.MIN_INT8), res)
			}
	
			if res := src.Uint8ToInt32Range(src.MAX_UINT8); res != int32(src.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int32(src.MAX_INT8), res)
			}
		})
	
		t.Run("Uint8ToInt64Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Uint8ToInt64Range(src.MIN_UINT8); res != int64(src.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(src.MIN_INT8), res)
			}
	
			if res := src.Uint8ToInt64Range(src.MAX_UINT8); res != int64(src.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int64(src.MAX_INT8), res)
			}
		})
	
		t.Run("Uint8ToIntRange", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Uint8ToIntRange(src.MIN_UINT8); res != int(src.MIN_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(src.MIN_INT8), res)
			}
	
			if res := src.Uint8ToIntRange(src.MAX_UINT8); res != int(src.MAX_INT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", int(src.MAX_INT8), res)
			}
		})
	
		t.Run("Int8ToUint8Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Int8ToUint8Range(src.MIN_INT8); res != src.MIN_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MIN_UINT8, res)
			}
	
			if res := src.Int8ToUint8Range(src.MAX_INT8); res != src.MAX_UINT8 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MAX_UINT8, res)
			}
		})
	
		t.Run("Int8ToUint16Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Int8ToUint16Range(src.MIN_INT8); res != uint16(src.MIN_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(src.MIN_UINT8), res)
			}
	
			if res := src.Int8ToUint16Range(src.MAX_INT8); res != uint16(src.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint16(src.MAX_UINT8), res)
			}
		})
	
		t.Run("Int8ToUint32Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Int8ToUint32Range(src.MIN_INT8); res != src.MIN_UINT32 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MIN_UINT32, res)
			}
	
			if res := src.Int8ToUint32Range(src.MAX_INT8); res != uint32(src.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint32(src.MAX_UINT8), res)
			}
		})
	
		t.Run("Int8ToUint64Range", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Int8ToUint64Range(src.MIN_INT8); res != src.MIN_UINT64 {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MIN_UINT64, res)
			}
	
			if res := src.Int8ToUint64Range(src.MAX_INT8); res != uint64(src.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint64(src.MAX_UINT8), res)
			}
		})
	
		t.Run("Int8ToUintRange", func(t *testing.T) {
			t.Parallel()
	
			if res := src.Int8ToUintRange(src.MIN_INT8); res != src.MIN_UINT {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", src.MIN_UINT, res)
			}
	
			if res := src.Int8ToUintRange(src.MAX_INT8); res != uint(src.MAX_UINT8) {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", uint(src.MAX_UINT8), res)
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
				if ref := reflect.ValueOf(src.MIN_UINT8).Kind(); ref != reflect.Uint8 {
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
				if ref := reflect.ValueOf(src.MIN_UINT16).Kind(); ref != reflect.Uint16 {
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
				if ref := reflect.ValueOf(src.MIN_UINT32).Kind(); ref != reflect.Uint32 {
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
				if ref := reflect.ValueOf(src.MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(src.MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(src.MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(src.MAX_UINT).Kind(); ref != reflect.Uint {
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
				if ref := reflect.ValueOf(src.MIN_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt8", func(t *testing.T) {
				t.Parallel()
				const expect = "int8"
				if ref := reflect.ValueOf(src.MAX_INT8).Kind(); ref != reflect.Int8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(src.MIN_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt16", func(t *testing.T) {
				t.Parallel()
				const expect = "int16"
				if ref := reflect.ValueOf(src.MAX_INT16).Kind(); ref != reflect.Int16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(src.MIN_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt32", func(t *testing.T) {
				t.Parallel()
				const expect = "int32"
				if ref := reflect.ValueOf(src.MAX_INT32).Kind(); ref != reflect.Int32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(src.MIN_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt64", func(t *testing.T) {
				t.Parallel()
				const expect = "int64"
				if ref := reflect.ValueOf(src.MAX_INT64).Kind(); ref != reflect.Int64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(src.MIN_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxInt", func(t *testing.T) {
				const expect = "int"
				if ref := reflect.ValueOf(src.MAX_INT).Kind(); ref != reflect.Int {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})

		t.Run("Uint", func(t *testing.T) {
			t.Parallel()

			t.Run("MinUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(src.MIN_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint8", func(t *testing.T) {
				t.Parallel()
				const expect = "uint8"
				if ref := reflect.ValueOf(src.MAX_UINT8).Kind(); ref != reflect.Uint8 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(src.MIN_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint16", func(t *testing.T) {
				t.Parallel()
				const expect = "uint16"
				if ref := reflect.ValueOf(src.MAX_UINT16).Kind(); ref != reflect.Uint16 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(src.MIN_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint32", func(t *testing.T) {
				t.Parallel()
				const expect = "uint32"
				if ref := reflect.ValueOf(src.MAX_UINT32).Kind(); ref != reflect.Uint32 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(src.MIN_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint64", func(t *testing.T) {
				t.Parallel()
				const expect = "uint64"
				if ref := reflect.ValueOf(src.MAX_UINT64).Kind(); ref != reflect.Uint64 {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MinUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(src.MIN_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})

			t.Run("MaxUint", func(t *testing.T) {
				const expect = "uint"
				if ref := reflect.ValueOf(src.MAX_UINT).Kind(); ref != reflect.Uint {
					t.Errorf("Expect(%s) => Got(%s)", expect, ref.String())
				}
			})
		})
	})
}
