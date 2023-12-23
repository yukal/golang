package test

import (
	"testing"
	"yu/golang/internal/app"
	"yu/golang/pkg/validation"
)

func TestIsMax(t *testing.T) {
	// int & int

	t.Run("i8_i8", func(t *testing.T) {
		t.Run("minI8_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_i16", func(t *testing.T) {
		t.Run("minI8_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_i32", func(t *testing.T) {
		t.Run("minI8_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_i64", func(t *testing.T) {
		t.Run("minI8_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_i", func(t *testing.T) {
		t.Run("minI8_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_rune", func(t *testing.T) {
		t.Run("minI8_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int16

	t.Run("i16_i8", func(t *testing.T) {
		t.Run("minI16_minI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_i16", func(t *testing.T) {
		t.Run("minI16_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_i32", func(t *testing.T) {
		t.Run("minI16_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_i64", func(t *testing.T) {
		t.Run("minI16_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_i", func(t *testing.T) {
		t.Run("minI16_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_rune", func(t *testing.T) {
		t.Run("minI16_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int32

	t.Run("i32_i8", func(t *testing.T) {
		t.Run("minI32_minI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_i16", func(t *testing.T) {
		t.Run("minI32_minI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_i32", func(t *testing.T) {
		t.Run("minI32_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_i64", func(t *testing.T) {
		t.Run("minI32_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_i", func(t *testing.T) {
		t.Run("minI32_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_rune", func(t *testing.T) {
		t.Run("minI32_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int64

	t.Run("i64_i8", func(t *testing.T) {
		t.Run("minI64_minI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_i16", func(t *testing.T) {
		t.Run("minI64_minI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_i32", func(t *testing.T) {
		t.Run("minI64_minI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_i64", func(t *testing.T) {
		t.Run("minI64_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_i", func(t *testing.T) {
		t.Run("minI64_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_rune", func(t *testing.T) {
		t.Run("minI64_minRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// intX

	t.Run("i_i8", func(t *testing.T) {
		t.Run("minI_minI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_i16", func(t *testing.T) {
		t.Run("minI_minI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_i32", func(t *testing.T) {
		t.Run("minI_minI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_i64", func(t *testing.T) {
		t.Run("minI_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_i", func(t *testing.T) {
		t.Run("minI_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_INT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_rune", func(t *testing.T) {
		t.Run("minI_minRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// rune

	t.Run("rune_i8", func(t *testing.T) {
		t.Run("minRune_minI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_i16", func(t *testing.T) {
		t.Run("minRune_minI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_i32", func(t *testing.T) {
		t.Run("minRune_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_i64", func(t *testing.T) {
		t.Run("minRune_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_i", func(t *testing.T) {
		t.Run("minRune_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_rune", func(t *testing.T) {
		t.Run("minRune_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// ................................
	// int & uint

	t.Run("i8_byte", func(t *testing.T) {
		t.Run("minI8_minByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_u8", func(t *testing.T) {
		t.Run("minI8_minU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_u16", func(t *testing.T) {
		t.Run("minI8_minU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_u32", func(t *testing.T) {
		t.Run("minI8_minU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_u64", func(t *testing.T) {
		t.Run("minI8_minU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i8_u", func(t *testing.T) {
		t.Run("minI8_minU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI8_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI8_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int16

	t.Run("i16_byte", func(t *testing.T) {
		t.Run("minI16_minByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_u8", func(t *testing.T) {
		t.Run("minI16_minU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_u16", func(t *testing.T) {
		t.Run("minI16_minU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_u32", func(t *testing.T) {
		t.Run("minI16_minU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_u64", func(t *testing.T) {
		t.Run("minI16_minU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i16_u", func(t *testing.T) {
		t.Run("minI16_minU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI16_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI16_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int32

	t.Run("i32_byte", func(t *testing.T) {
		t.Run("minI32_minByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_u8", func(t *testing.T) {
		t.Run("minI32_minU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_u16", func(t *testing.T) {
		t.Run("minI32_minU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_u32", func(t *testing.T) {
		t.Run("minI32_minU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_u64", func(t *testing.T) {
		t.Run("minI32_minU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i32_u", func(t *testing.T) {
		t.Run("minI32_minU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI32_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI32_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// int64

	t.Run("i64_byte", func(t *testing.T) {
		t.Run("minI64_minByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_u8", func(t *testing.T) {
		t.Run("minI64_minU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_u16", func(t *testing.T) {
		t.Run("minI64_minU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_u32", func(t *testing.T) {
		t.Run("minI64_minU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_u64", func(t *testing.T) {
		t.Run("minI64_minU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i64_u", func(t *testing.T) {
		t.Run("minI64_minU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI64_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI64_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// intX

	t.Run("i_byte", func(t *testing.T) {
		t.Run("minI_minByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_u8", func(t *testing.T) {
		t.Run("minI_minU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_u16", func(t *testing.T) {
		t.Run("minI_minU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_u32", func(t *testing.T) {
		t.Run("minI_minU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_u64", func(t *testing.T) {
		t.Run("minI_minU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("i_u", func(t *testing.T) {
		t.Run("minI_minU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_INT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minI_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_INT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxI_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_INT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// rune

	t.Run("rune_u8", func(t *testing.T) {
		t.Run("minRune_minU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_u16", func(t *testing.T) {
		t.Run("minRune_minU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_u32", func(t *testing.T) {
		t.Run("minRune_minU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_u64", func(t *testing.T) {
		t.Run("minRune_minU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_u", func(t *testing.T) {
		t.Run("minRune_minU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("rune_byte", func(t *testing.T) {
		t.Run("minRune_minByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minRune_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_RUNE, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxRune_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_RUNE, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// ................................
	// uint & uint

	t.Run("byte_u8", func(t *testing.T) {
		t.Run("minByte_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_u16", func(t *testing.T) {
		t.Run("minByte_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_u32", func(t *testing.T) {
		t.Run("minByte_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_u64", func(t *testing.T) {
		t.Run("minByte_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_u", func(t *testing.T) {
		t.Run("minByte_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint8

	t.Run("u8_byte", func(t *testing.T) {
		t.Run("minU8_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_u8", func(t *testing.T) {
		t.Run("minU8_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_u16", func(t *testing.T) {
		t.Run("minU8_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_u32", func(t *testing.T) {
		t.Run("minU8_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_u64", func(t *testing.T) {
		t.Run("minU8_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_u", func(t *testing.T) {
		t.Run("minU8_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint16

	t.Run("u16_byte", func(t *testing.T) {
		t.Run("minU16_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_u8", func(t *testing.T) {
		t.Run("minU16_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_u16", func(t *testing.T) {
		t.Run("minU16_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_u32", func(t *testing.T) {
		t.Run("minU16_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_u64", func(t *testing.T) {
		t.Run("minU16_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_u", func(t *testing.T) {
		t.Run("minU16_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint32

	t.Run("u32_byte", func(t *testing.T) {
		t.Run("minU32_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_u8", func(t *testing.T) {
		t.Run("minU32_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_u16", func(t *testing.T) {
		t.Run("minU32_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_u32", func(t *testing.T) {
		t.Run("minU32_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_u64", func(t *testing.T) {
		t.Run("minU32_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_u", func(t *testing.T) {
		t.Run("minU32_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint64

	t.Run("u64_byte", func(t *testing.T) {
		t.Run("minU64_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_u8", func(t *testing.T) {
		t.Run("minU64_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_u16", func(t *testing.T) {
		t.Run("minU64_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_u32", func(t *testing.T) {
		t.Run("minU64_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_u64", func(t *testing.T) {
		t.Run("minU64_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_u", func(t *testing.T) {
		t.Run("minU64_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uintX

	t.Run("u_byte", func(t *testing.T) {
		t.Run("minU_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxByte", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minByte", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_BYTE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_u8", func(t *testing.T) {
		t.Run("minU_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxU8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minU8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_UINT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_u16", func(t *testing.T) {
		t.Run("minU_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxU16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minU16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_UINT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_u32", func(t *testing.T) {
		t.Run("minU_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxU32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minU32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_UINT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_u64", func(t *testing.T) {
		t.Run("minU_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxU64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minU64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_UINT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_u", func(t *testing.T) {
		t.Run("minU_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxU", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minU", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_UINT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// ................................
	// uint & int

	t.Run("byte_i8", func(t *testing.T) {
		t.Run("minByte_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_i16", func(t *testing.T) {
		t.Run("minByte_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_i32", func(t *testing.T) {
		t.Run("minByte_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_i64", func(t *testing.T) {
		t.Run("minByte_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_i", func(t *testing.T) {
		t.Run("minByte_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("byte_rune", func(t *testing.T) {
		t.Run("minByte_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_BYTE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_BYTE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minByte_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_BYTE, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxByte_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_BYTE, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint8

	t.Run("u8_i8", func(t *testing.T) {
		t.Run("minU8_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_i16", func(t *testing.T) {
		t.Run("minU8_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_i32", func(t *testing.T) {
		t.Run("minU8_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_i64", func(t *testing.T) {
		t.Run("minU8_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_i", func(t *testing.T) {
		t.Run("minU8_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u8_rune", func(t *testing.T) {
		t.Run("minU8_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU8_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT8, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU8_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT8, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint16

	t.Run("u16_i8", func(t *testing.T) {
		t.Run("minU16_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_i16", func(t *testing.T) {
		t.Run("minU16_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_i32", func(t *testing.T) {
		t.Run("minU16_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_i64", func(t *testing.T) {
		t.Run("minU16_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_i", func(t *testing.T) {
		t.Run("minU16_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u16_rune", func(t *testing.T) {
		t.Run("minU16_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU16_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT16, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU16_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT16, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint32

	t.Run("u32_i8", func(t *testing.T) {
		t.Run("minU32_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_i16", func(t *testing.T) {
		t.Run("minU32_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_i32", func(t *testing.T) {
		t.Run("minU32_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_i64", func(t *testing.T) {
		t.Run("minU32_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_i", func(t *testing.T) {
		t.Run("minU32_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u32_rune", func(t *testing.T) {
		t.Run("minU32_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU32_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT32, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU32_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT32, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uint64

	t.Run("u64_i8", func(t *testing.T) {
		t.Run("minU64_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_i16", func(t *testing.T) {
		t.Run("minU64_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_i32", func(t *testing.T) {
		t.Run("minU64_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_i64", func(t *testing.T) {
		t.Run("minU64_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_i", func(t *testing.T) {
		t.Run("minU64_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u64_rune", func(t *testing.T) {
		t.Run("minU64_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU64_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT64, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU64_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT64, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	// uintX

	t.Run("u_i8", func(t *testing.T) {
		t.Run("minU_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxI8", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minI8", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_INT8); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_i16", func(t *testing.T) {
		t.Run("minU_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxI16", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minI16", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_INT16); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_i32", func(t *testing.T) {
		t.Run("minU_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxI32", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minI32", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_INT32); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_i64", func(t *testing.T) {
		t.Run("minU_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxI64", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minI64", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_INT64); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_i", func(t *testing.T) {
		t.Run("minU_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxI", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minI", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_INT); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

	t.Run("u_rune", func(t *testing.T) {
		t.Run("minU_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MIN_UINT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_maxRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("minU_maxRune", func(t *testing.T) {
			const expect = false

			if res := validation.IsMax(app.MIN_UINT, app.MAX_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})

		t.Run("maxU_minRune", func(t *testing.T) {
			const expect = true

			if res := validation.IsMax(app.MAX_UINT, app.MIN_RUNE); res != expect {
				t.Errorf("Expect(%#v) Got(%#v)", expect, res)
			}
		})
	})

}
