package main

import (
	"testing"
)

func TestValidation(t *testing.T) {
	t.Run("Compare", func(t *testing.T) {
		t.Run("min", func(t *testing.T) {

			// int8

			t.Run("i8_i8", func(t *testing.T) {
				t.Run("minI8_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_i16", func(t *testing.T) {
				t.Run("minI8_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT8, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_i32", func(t *testing.T) {
				t.Run("minI8_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT8, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_i64", func(t *testing.T) {
				t.Run("minI8_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT8, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_i", func(t *testing.T) {
				t.Run("minI8_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT8, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int16

			t.Run("i16_i8", func(t *testing.T) {
				t.Run("minI16_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_i16", func(t *testing.T) {
				t.Run("minI16_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_i32", func(t *testing.T) {
				t.Run("minI16_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT16, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_i64", func(t *testing.T) {
				t.Run("minI16_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT16, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_i", func(t *testing.T) {
				t.Run("minI16_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT16, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int32

			t.Run("i32_i8", func(t *testing.T) {
				t.Run("minI32_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_i16", func(t *testing.T) {
				t.Run("minI32_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_i32", func(t *testing.T) {
				t.Run("minI32_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT32, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_i64", func(t *testing.T) {
				t.Run("minI32_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT32, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT32, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_i", func(t *testing.T) {
				t.Run("minI32_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_INT32, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT32, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int64

			t.Run("i64_i8", func(t *testing.T) {
				t.Run("minI64_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_i16", func(t *testing.T) {
				t.Run("minI64_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_i32", func(t *testing.T) {
				t.Run("minI64_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_i64", func(t *testing.T) {
				t.Run("minI64_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT64, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_i", func(t *testing.T) {
				t.Run("minI64_minI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT64, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int

			t.Run("i_i8", func(t *testing.T) {
				t.Run("minI_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_i16", func(t *testing.T) {
				t.Run("minI_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_i32", func(t *testing.T) {
				t.Run("minI_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_i64", func(t *testing.T) {
				t.Run("minI_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_i", func(t *testing.T) {
				t.Run("minI_minI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// ................................
			// int8 & uint

			t.Run("i8_u8", func(t *testing.T) {
				t.Run("minI8_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_u16", func(t *testing.T) {
				t.Run("minI8_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_u32", func(t *testing.T) {
				t.Run("minI8_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_u64", func(t *testing.T) {
				t.Run("minI8_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i8_u", func(t *testing.T) {
				t.Run("minI8_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT8, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI8_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT8, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI8_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT8, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int16 & uint

			t.Run("i16_u8", func(t *testing.T) {
				t.Run("minI16_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_u16", func(t *testing.T) {
				t.Run("minI16_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_u32", func(t *testing.T) {
				t.Run("minI16_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_u64", func(t *testing.T) {
				t.Run("minI16_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i16_u", func(t *testing.T) {
				t.Run("minI16_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT16, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI16_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT16, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI16_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT16, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int32 & uint

			t.Run("i32_u8", func(t *testing.T) {
				t.Run("minI32_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_u16", func(t *testing.T) {
				t.Run("minI32_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_u32", func(t *testing.T) {
				t.Run("minI32_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT32, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_u64", func(t *testing.T) {
				t.Run("minI32_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT32, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i32_u", func(t *testing.T) {
				t.Run("minI32_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT32, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI32_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT32, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI32_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT32, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int64 & uint

			t.Run("i64_u8", func(t *testing.T) {
				t.Run("minI64_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_u16", func(t *testing.T) {
				t.Run("minI64_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_u32", func(t *testing.T) {
				t.Run("minI64_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_u64", func(t *testing.T) {
				t.Run("minI64_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT64, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i64_u", func(t *testing.T) {
				t.Run("minI64_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT64, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI64_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT64, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI64_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT64, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// int & uint

			t.Run("i_u8", func(t *testing.T) {
				t.Run("minI_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_u16", func(t *testing.T) {
				t.Run("minI_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_u32", func(t *testing.T) {
				t.Run("minI_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_u64", func(t *testing.T) {
				t.Run("minI_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("i_u", func(t *testing.T) {
				t.Run("minI_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_INT, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minI_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_INT, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxI_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_INT, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// ................................
			// uint8

			t.Run("u8_u8", func(t *testing.T) {
				t.Run("minU8_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_u16", func(t *testing.T) {
				t.Run("minU8_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_u32", func(t *testing.T) {
				t.Run("minU8_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_u64", func(t *testing.T) {
				t.Run("minU8_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_u", func(t *testing.T) {
				t.Run("minU8_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint16

			t.Run("u16_u8", func(t *testing.T) {
				t.Run("minU16_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_u16", func(t *testing.T) {
				t.Run("minU16_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_u32", func(t *testing.T) {
				t.Run("minU16_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_u64", func(t *testing.T) {
				t.Run("minU16_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_u", func(t *testing.T) {
				t.Run("minU16_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint32

			t.Run("u32_u8", func(t *testing.T) {
				t.Run("minU32_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_u16", func(t *testing.T) {
				t.Run("minU32_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_u32", func(t *testing.T) {
				t.Run("minU32_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT32, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_u64", func(t *testing.T) {
				t.Run("minU32_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT32, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_u", func(t *testing.T) {
				t.Run("minU32_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT32, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint64

			t.Run("u64_u8", func(t *testing.T) {
				t.Run("minU64_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_u16", func(t *testing.T) {
				t.Run("minU64_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_u32", func(t *testing.T) {
				t.Run("minU64_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_u64", func(t *testing.T) {
				t.Run("minU64_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT64, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_u", func(t *testing.T) {
				t.Run("minU64_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT64, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint

			t.Run("u_u8", func(t *testing.T) {
				t.Run("minU_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxU8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minU8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_UINT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_u16", func(t *testing.T) {
				t.Run("minU_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxU16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minU16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_UINT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_u32", func(t *testing.T) {
				t.Run("minU_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxU32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minU32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_UINT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_u64", func(t *testing.T) {
				t.Run("minU_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxU64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minU64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_UINT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_u", func(t *testing.T) {
				t.Run("minU_minU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxU", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minU", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_UINT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint8 & int

			t.Run("u8_i8", func(t *testing.T) {
				t.Run("minU8_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT8, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_i16", func(t *testing.T) {
				t.Run("minU8_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT8, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_i32", func(t *testing.T) {
				t.Run("minU8_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT8, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_i64", func(t *testing.T) {
				t.Run("minU8_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT8, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u8_i", func(t *testing.T) {
				t.Run("minU8_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT8, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT8, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU8_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT8, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU8_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT8, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint16 & int

			t.Run("u16_i8", func(t *testing.T) {
				t.Run("minU16_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT16, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_i16", func(t *testing.T) {
				t.Run("minU16_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT16, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_i32", func(t *testing.T) {
				t.Run("minU16_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT16, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_i64", func(t *testing.T) {
				t.Run("minU16_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT16, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u16_i", func(t *testing.T) {
				t.Run("minU16_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT16, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT16, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU16_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT16, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU16_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT16, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint32 & int

			t.Run("u32_i8", func(t *testing.T) {
				t.Run("minU32_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT32, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_i16", func(t *testing.T) {
				t.Run("minU32_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT32, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_i32", func(t *testing.T) {
				t.Run("minU32_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT32, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_i64", func(t *testing.T) {
				t.Run("minU32_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT32, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT32, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u32_i", func(t *testing.T) {
				t.Run("minU32_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT32, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MAX_UINT32, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU32_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT32, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU32_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT32, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint64 & int

			t.Run("u64_i8", func(t *testing.T) {
				t.Run("minU64_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT64, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_i16", func(t *testing.T) {
				t.Run("minU64_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT64, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_i32", func(t *testing.T) {
				t.Run("minU64_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT64, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_i64", func(t *testing.T) {
				t.Run("minU64_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT64, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u64_i", func(t *testing.T) {
				t.Run("minU64_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT64, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU64_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT64, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU64_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT64, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			// uint & int

			t.Run("u_i8", func(t *testing.T) {
				t.Run("minU_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxI8", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minI8", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_INT8); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_i16", func(t *testing.T) {
				t.Run("minU_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxI16", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minI16", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_INT16); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_i32", func(t *testing.T) {
				t.Run("minU_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxI32", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minI32", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_INT32); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_i64", func(t *testing.T) {
				t.Run("minU_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxI64", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minI64", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_INT64); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})

			t.Run("u_i", func(t *testing.T) {
				t.Run("minU_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MIN_UINT, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("minU_maxI", func(t *testing.T) {
					t.Parallel()
					const expect = true

					if res := Compare("min", MIN_UINT, MAX_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})

				t.Run("maxU_minI", func(t *testing.T) {
					t.Parallel()
					const expect = false

					if res := Compare("min", MAX_UINT, MIN_INT); res != expect {
						t.Errorf("Expect(%#v) Got(%#v)", expect, res)
					}
				})
			})
		})

		t.Run("max", func(t *testing.T) {
			t.Skip()
		})

		t.Run("eq", func(t *testing.T) {
			t.Skip()
		})
	})
}
