package test

import (
	"testing"
	"yu/golang/src/validation"
)

func TestIsEqual(t *testing.T) {

	t.Run("numeric", func(t *testing.T) {
		t.Run("IsEqual(0,-1)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(0, -1); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(0,0)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(0, 0); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(0,1)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(0, 1); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ......................

		t.Run("IsEqual(1,-1)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, -1); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,0)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, 0); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,1)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(1, 1); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("array", func(t *testing.T) {
		var arrEmpty [0]string
		var arrFilled [1]string

		t.Run("IsEqual(0,arrEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(0, arrEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(0,arrFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(0, arrFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,arrEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, arrEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,arrFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(1, arrFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("slice", func(t *testing.T) {
		sliceEmpty := []string{}
		sliceFilled := []string{"1"}

		t.Run("IsEqual(0,sliceEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(0, sliceEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(0,sliceFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(0, sliceFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,sliceEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, sliceEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,sliceFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(1, sliceFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("map", func(t *testing.T) {
		mapEmpty := map[string]string{}
		mapFilled := map[string]string{"item": "val"}

		t.Run("IsEqual(0,mapEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(0, mapEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(0,mapFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(0, mapFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,mapEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, mapEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,mapFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(1, mapFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("string", func(t *testing.T) {
		strEmpty := ""
		strFilled := "a"

		t.Run("IsEqual(0,strEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(0, strEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(0,strFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(0, strFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,strEmpty)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, strEmpty); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,strFilled)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			if res := validation.IsEqual(1, strFilled); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("emptiness", func(t *testing.T) {
		t.Run("IsEqual(nil,nil)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(nil, nil); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(1,nil)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(1, nil); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsEqual(nil,1)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEqual(nil, 1); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

}
