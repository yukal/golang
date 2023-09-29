package test

import (
	"testing"
	"yu/golang/internal/app/validation"
)

func TestIsEachMatches(t *testing.T) {

	// TODO: check chan (channels)
	t.Run("FilledSlice", func(t *testing.T) {
		t.Parallel()
		const expect = true

		reg := `(?i)^[0-9a-f]{32}$`
		data := []string{"b0fb0c19711bcf3b73f41c909f66bded"}

		if res := validation.IsEachMatches(reg, data); res != expect {
			t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
		}
	})

	t.Run("FilledArray", func(t *testing.T) {
		t.Parallel()
		const expect = true

		reg := `^38[0-9]{10}$`
		data := [1]string{"380001234567"}

		if res := validation.IsEachMatches(reg, data); res != expect {
			t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
		}
	})

	t.Run("FilledMap", func(t *testing.T) {
		t.Parallel()
		const expect = true

		reg := `^https\://img\.domain\.com/[0-9A-Fa-f]{32}\.(?:pn|jpe?)g$`
		data := map[string]string{
			"img1": "https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
			"img2": "https://img.domain.com/4792592a98f8b9143de71d1db403d163.jpg",
			"img3": "https://img.domain.com/92f2b876b8ea94f711d2173539e73802.png",
		}

		if res := validation.IsEachMatches(reg, data); res != expect {
			t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
		}
	})

	// ...

	t.Run("emptiness", func(t *testing.T) {
		t.Run("nil_&_nil", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEachMatches(nil, nil); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("nil_&_EmptySlice", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEachMatches(nil, []string{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("nil_&_EmptyArray)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEachMatches(nil, [0]string{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("nil_&_EmptyMap)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEachMatches(nil, map[string]string{}); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("EmptyRegex_&_nil", func(t *testing.T) {
			t.Parallel()
			const expect = false

			if res := validation.IsEachMatches(``, nil); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("EmptyRegex_&_EmptySlice)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			reg := ``
			data := []string{}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("EmptyRegex_&_EmptyArray", func(t *testing.T) {
			t.Parallel()
			const expect = true

			reg := ``
			data := [0]string{}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("EmptyRegex_&_EmptyMap", func(t *testing.T) {
			t.Parallel()
			const expect = true

			reg := ``
			data := map[string]string{}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("EmptyRegex_&_FilledSlice", func(t *testing.T) {
			t.Parallel()
			const expect = true

			reg := ``
			data := []string{"str"}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("EmptyRegex_&_FilledArray)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			reg := ``
			data := [1]string{"str"}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("EmptyRegex_&_FilledMap)", func(t *testing.T) {
			t.Parallel()
			const expect = true

			reg := ``
			data := map[string]string{"key": "val"}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("FilledRegex_&_EmptyData", func(t *testing.T) {
			t.Parallel()
			const expect = false

			reg := `(?i)^[0-9a-f]{32}$`
			data := []string{""}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("FilledRegex_&_EmptyData)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			reg := `^38[0-9]{10}$`
			data := [1]string{""}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("FilledRegex_&_EmptyData)", func(t *testing.T) {
			t.Parallel()
			const expect = false

			reg := `^https\://img\.domain\.com/[0-9A-Fa-f]{32}\.(?:pn|jpe?)g$`
			data := map[string]string{
				"img1": "",
				"img2": "",
				"img3": "",
			}

			if res := validation.IsEachMatches(reg, data); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

}
