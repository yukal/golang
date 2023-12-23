package test

import (
	"reflect"
	"testing"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
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

	g := Goblin(t)

	g.Describe("eq", func() {
		g.Describe("numeric", func() {
			g.It("IsEqual(4,-4)", func() {
				const expect = "must be exactly 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,4)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,8)", func() {
				const expect = "must be exactly 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("IsEqual(0,arrEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,arrEmpty)", func() {
				const expect = "must contain exactly 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(0,arrFilled)", func() {
				const expect = "must contain exactly 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,arrFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("IsEqual(0,sliceEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,sliceEmpty)", func() {
				const expect = "must contain exactly 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(0,sliceFilled)", func() {
				const expect = "must contain exactly 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,sliceFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("IsEqual(0,mapEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,mapEmpty)", func() {
				const expect = "must contain exactly 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(0,mapFilled)", func() {
				const expect = "must contain exactly 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,mapFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("IsEqual(0,strEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,strEmpty)", func() {
				const expect = "must contain exactly 4 characters"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(0,strFilled)", func() {
				const expect = "must contain exactly 0 characters"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(4,strFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("IsEqual(nil,nil)", func() {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(1,nil)", func() {
				const expect = ""

				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("IsEqual(nil,1)", func() {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				if res := validation.Compare("eq", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})
	})

	g.Describe("max", func() {
		g.Describe("numeric", func() {
			g.It("max(4,-4)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,4)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,8)", func() {
				const expect = "must be up to 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("max(0,arrEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,arrEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(0,arrFilled)", func() {
				const expect = "must contain up to 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,arrFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("max(0,sliceEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,sliceEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(0,sliceFilled)", func() {
				const expect = "must contain up to 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,sliceFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("max(0,mapEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,mapEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(0,mapFilled)", func() {
				const expect = "must contain up to 0 items"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,mapFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("max(0,strEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,strEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(0,strFilled)", func() {
				const expect = "must contain up to 0 characters"

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(4,strFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("max(nil,nil)", func() {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(1,nil)", func() {
				const expect = ""

				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("max(nil,1)", func() {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				if res := validation.Compare("max", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

	})

	g.Describe("min", func() {
		g.Describe("numeric", func() {
			g.It("min(4,-4)", func() {
				const expect = "must be at least 4"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,4)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,8)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("min(0,arrEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,arrEmpty)", func() {
				const expect = "must contain at least 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(0,arrFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,arrFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("min(0,sliceEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,sliceEmpty)", func() {
				const expect = "must contain at least 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(0,sliceFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,sliceFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("min(0,mapEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,mapEmpty)", func() {
				const expect = "must contain at least 4 items"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(0,mapFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,mapFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("min(0,strEmpty)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,strEmpty)", func() {
				const expect = "must contain at least 4 characters"

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strEmpty)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(0,strFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(0)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(4,strFilled)", func() {
				const expect = ""

				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(strFilled)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("min(nil,nil)", func() {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(1,nil)", func() {
				const expect = ""

				proto := reflect.ValueOf(1)
				value := reflect.ValueOf(nil)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})

			g.It("min(nil,1)", func() {
				const expect = ""

				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(1)

				if res := validation.Compare("min", proto, value); res != expect {
					g.Errorf("Expect( %v ) => Got( %v )", expect, res)
				}
			})
		})
	})

}
