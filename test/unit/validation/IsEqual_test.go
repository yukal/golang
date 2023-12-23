package test

import (
	"testing"
	"yu/golang/internal/app"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
)

func TestIsEqual(t *testing.T) {
	g := Goblin(t)

	// int & int

	g.Describe("i8_i8", func() {
		g.It("minI8_minI8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxI8", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_i16", func() {
		g.It("minI8_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_i32", func() {
		g.It("minI8_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_i64", func() {
		g.It("minI8_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_i", func() {
		g.It("minI8_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_rune", func() {
		g.It("minI8_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int16

	g.Describe("i16_i8", func() {
		g.It("minI16_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_i16", func() {
		g.It("minI16_minI16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxI16", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_i32", func() {
		g.It("minI16_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_i64", func() {
		g.It("minI16_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_i", func() {
		g.It("minI16_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_rune", func() {
		g.It("minI16_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int32

	g.Describe("i32_i8", func() {
		g.It("minI32_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_i16", func() {
		g.It("minI32_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_i32", func() {
		g.It("minI32_minI32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxI32", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_i64", func() {
		g.It("minI32_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_i", func() {
		g.It("minI32_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_rune", func() {
		g.It("minI32_minRune", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxRune", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int64

	g.Describe("i64_i8", func() {
		g.It("minI64_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_i16", func() {

		g.It("minI64_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_i32", func() {

		g.It("minI64_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_i64", func() {

		g.It("minI64_minI64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxI64", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_i", func() {

		g.It("minI64_minI", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxI", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_rune", func() {

		g.It("minI64_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// intX

	g.Describe("i_i8", func() {

		g.It("minI_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_i16", func() {

		g.It("minI_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_i32", func() {

		g.It("minI_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_i64", func() {

		g.It("minI_minI64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxI64", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_i", func() {

		g.It("minI_minI", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_INT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxI", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_INT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_rune", func() {

		g.It("minI_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// rune

	g.Describe("rune_i8", func() {

		g.It("minRune_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_i16", func() {

		g.It("minRune_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_i32", func() {

		g.It("minRune_minI32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxI32", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_i64", func() {

		g.It("minRune_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_i", func() {

		g.It("minRune_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_rune", func() {

		g.It("minRune_minRune", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxRune", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// ................................
	// int & uint

	g.Describe("i8_byte", func() {

		g.It("minI8_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_u8", func() {

		g.It("minI8_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_u16", func() {

		g.It("minI8_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_u32", func() {

		g.It("minI8_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_u64", func() {
		g.It("minI8_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i8_u", func() {
		g.It("minI8_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI8_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI8_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int16

	g.Describe("i16_byte", func() {
		g.It("minI16_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_u8", func() {
		g.It("minI16_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_u16", func() {
		g.It("minI16_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_u32", func() {
		g.It("minI16_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_u64", func() {
		g.It("minI16_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i16_u", func() {
		g.It("minI16_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI16_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI16_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int32

	g.Describe("i32_byte", func() {
		g.It("minI32_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_u8", func() {
		g.It("minI32_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_u16", func() {
		g.It("minI32_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_u32", func() {
		g.It("minI32_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_u64", func() {
		g.It("minI32_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i32_u", func() {
		g.It("minI32_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI32_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI32_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int64

	g.Describe("i64_byte", func() {
		g.It("minI64_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_u8", func() {
		g.It("minI64_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_u16", func() {
		g.It("minI64_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_u32", func() {
		g.It("minI64_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_u64", func() {
		g.It("minI64_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i64_u", func() {
		g.It("minI64_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI64_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI64_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// intX

	g.Describe("i_byte", func() {
		g.It("minI_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_u8", func() {
		g.It("minI_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_u16", func() {
		g.It("minI_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_u32", func() {
		g.It("minI_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_u64", func() {
		g.It("minI_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("i_u", func() {
		g.It("minI_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minI_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_INT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxI_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_INT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// rune

	g.Describe("rune_u8", func() {
		g.It("minRune_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_u16", func() {
		g.It("minRune_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_u32", func() {
		g.It("minRune_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_u64", func() {
		g.It("minRune_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_u", func() {
		g.It("minRune_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("rune_byte", func() {
		g.It("minRune_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minRune_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_RUNE, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxRune_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_RUNE, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// ................................
	// uint & uint

	g.Describe("byte_u8", func() {
		g.It("minByte_minU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_u16", func() {
		g.It("minByte_minU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_u32", func() {
		g.It("minByte_minU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_u64", func() {
		g.It("minByte_minU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_u", func() {
		g.It("minByte_minU", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint8

	g.Describe("u8_byte", func() {
		g.It("minU8_minByte", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxByte", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_u8", func() {
		g.It("minU8_minU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_u16", func() {
		g.It("minU8_minU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_u32", func() {
		g.It("minU8_minU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_u64", func() {
		g.It("minU8_minU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_u", func() {
		g.It("minU8_minU", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint16

	g.Describe("u16_byte", func() {
		g.It("minU16_minByte", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_u8", func() {
		g.It("minU16_minU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_u16", func() {
		g.It("minU16_minU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_u32", func() {
		g.It("minU16_minU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_u64", func() {
		g.It("minU16_minU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_u", func() {
		g.It("minU16_minU", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint32

	g.Describe("u32_byte", func() {
		g.It("minU32_minByte", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_u8", func() {
		g.It("minU32_minU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_u16", func() {
		g.It("minU32_minU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_u32", func() {
		g.It("minU32_minU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_u64", func() {
		g.It("minU32_minU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_u", func() {
		g.It("minU32_minU", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint64

	g.Describe("u64_byte", func() {
		g.It("minU64_minByte", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_u8", func() {
		g.It("minU64_minU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_u16", func() {
		g.It("minU64_minU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_u32", func() {
		g.It("minU64_minU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_u64", func() {
		g.It("minU64_minU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_u", func() {
		g.It("minU64_minU", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxU", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uintX

	g.Describe("u_byte", func() {
		g.It("minU_minByte", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minByte", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_u8", func() {
		g.It("minU_minU8", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minU8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_u16", func() {
		g.It("minU_minU16", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minU16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_u32", func() {
		g.It("minU_minU32", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minU32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_u64", func() {
		g.It("minU_minU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxU64", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minU64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_u", func() {
		g.It("minU_minU", func() {
			const expect = true

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxU", func() {
			const expect = true

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxU", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minU", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// ................................
	// uint & int

	g.Describe("byte_i8", func() {
		g.It("minByte_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_i16", func() {
		g.It("minByte_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_i32", func() {
		g.It("minByte_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_i64", func() {
		g.It("minByte_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_i", func() {
		g.It("minByte_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("byte_rune", func() {
		g.It("minByte_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minByte_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_BYTE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxByte_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_BYTE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint8

	g.Describe("u8_i8", func() {
		g.It("minU8_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_i16", func() {
		g.It("minU8_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_i32", func() {
		g.It("minU8_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_i64", func() {
		g.It("minU8_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_i", func() {
		g.It("minU8_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u8_rune", func() {
		g.It("minU8_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU8_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU8_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint16

	g.Describe("u16_i8", func() {
		g.It("minU16_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_i16", func() {
		g.It("minU16_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_i32", func() {
		g.It("minU16_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_i64", func() {
		g.It("minU16_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_i", func() {
		g.It("minU16_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u16_rune", func() {
		g.It("minU16_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU16_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU16_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint32

	g.Describe("u32_i8", func() {
		g.It("minU32_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_i16", func() {
		g.It("minU32_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_i32", func() {
		g.It("minU32_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_i64", func() {
		g.It("minU32_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_i", func() {
		g.It("minU32_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u32_rune", func() {
		g.It("minU32_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU32_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU32_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint64

	g.Describe("u64_i8", func() {
		g.It("minU64_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_i16", func() {
		g.It("minU64_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_i32", func() {
		g.It("minU64_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_i64", func() {
		g.It("minU64_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_i", func() {
		g.It("minU64_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u64_rune", func() {
		g.It("minU64_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU64_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU64_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uintX

	g.Describe("u_i8", func() {
		g.It("minU_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minI8", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_i16", func() {
		g.It("minU_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minI16", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_i32", func() {
		g.It("minU_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minI32", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_i64", func() {
		g.It("minU_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minI64", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_i", func() {
		g.It("minU_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxI", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minI", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	g.Describe("u_rune", func() {
		g.It("minU_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("minU_maxRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MIN_UINT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		g.It("maxU_minRune", func() {
			const expect = false

			if res := validation.IsEqual(app.MAX_UINT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

}
