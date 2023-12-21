package test

import (
	"reflect"
	"testing"
	"yu/golang/pkg/validation"
)

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
			"m": "val3",
		}
	)

	t.Run("eq", func(t *testing.T) {

		t.Run("numeric", func(t *testing.T) {
			t.Run("IsEqual(4,-4)", func(t *testing.T) {
				const expect = "must be exactly 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,4)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,8)", func(t *testing.T) {
				const expect = "must be exactly 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("array", func(t *testing.T) {
			t.Run("IsEqual(0,arrEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,arrEmpty)", func(t *testing.T) {
				const expect = "must contain exactly 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(0,arrFilled)", func(t *testing.T) {
				const expect = "must contain exactly 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,arrFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("slice", func(t *testing.T) {
			t.Run("IsEqual(0,sliceEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,sliceEmpty)", func(t *testing.T) {
				const expect = "must contain exactly 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(0,sliceFilled)", func(t *testing.T) {
				const expect = "must contain exactly 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,sliceFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("map", func(t *testing.T) {
			t.Run("IsEqual(0,mapEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,mapEmpty)", func(t *testing.T) {
				const expect = "must contain exactly 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(0,mapFilled)", func(t *testing.T) {
				const expect = "must contain exactly 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,mapFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("string", func(t *testing.T) {
			t.Run("IsEqual(0,strEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,strEmpty)", func(t *testing.T) {
				const expect = "must contain exactly 4 characters"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(0,strFilled)", func(t *testing.T) {
				const expect = "must contain exactly 0 characters"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(4,strFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("emptiness", func(t *testing.T) {
			t.Run("IsEqual(nil,nil)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(1,nil)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("IsEqual(nil,1)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				if res := validation.Compare("eq", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})
	})

	t.Run("max", func(t *testing.T) {

		t.Run("numeric", func(t *testing.T) {
			t.Run("max(4,-4)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,4)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,8)", func(t *testing.T) {
				const expect = "must be up to 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("array", func(t *testing.T) {
			t.Run("max(0,arrEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,arrEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(0,arrFilled)", func(t *testing.T) {
				const expect = "must contain up to 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,arrFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("slice", func(t *testing.T) {
			t.Run("max(0,sliceEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,sliceEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(0,sliceFilled)", func(t *testing.T) {
				const expect = "must contain up to 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,sliceFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("map", func(t *testing.T) {
			t.Run("max(0,mapEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,mapEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(0,mapFilled)", func(t *testing.T) {
				const expect = "must contain up to 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,mapFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("string", func(t *testing.T) {
			t.Run("max(0,strEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,strEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(0,strFilled)", func(t *testing.T) {
				const expect = "must contain up to 0 characters"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(4,strFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("emptiness", func(t *testing.T) {
			t.Run("max(nil,nil)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(1,nil)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("max(nil,1)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				if res := validation.Compare("max", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

	})

	t.Run("min", func(t *testing.T) {
		t.Run("numeric", func(t *testing.T) {
			t.Run("min(4,-4)", func(t *testing.T) {
				const expect = "must be at least 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,4)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,8)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("array", func(t *testing.T) {
			t.Run("min(0,arrEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,arrEmpty)", func(t *testing.T) {
				const expect = "must contain at least 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(0,arrFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,arrFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("slice", func(t *testing.T) {
			t.Run("min(0,sliceEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,sliceEmpty)", func(t *testing.T) {
				const expect = "must contain at least 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(0,sliceFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,sliceFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("map", func(t *testing.T) {
			t.Run("min(0,mapEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,mapEmpty)", func(t *testing.T) {
				const expect = "must contain at least 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(0,mapFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,mapFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("string", func(t *testing.T) {
			t.Run("min(0,strEmpty)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,strEmpty)", func(t *testing.T) {
				const expect = "must contain at least 4 characters"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(0,strFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(4,strFilled)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		t.Run("emptiness", func(t *testing.T) {
			t.Run("min(nil,nil)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(1,nil)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			t.Run("min(nil,1)", func(t *testing.T) {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				if res := validation.Compare("min", proto, value); res != expect {
					t.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})
	})

}
