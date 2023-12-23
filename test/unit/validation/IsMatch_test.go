package test

import (
	"testing"
	"yu/golang/pkg/validation"
)

func TestIsMatch(t *testing.T) {

	t.Run("IsMatch(hexadecimal,hex)", func(t *testing.T) {
		const expect = true

		reg := `(?i)^[0-9a-f]{32}$`
		val := "b0fb0c19711bcf3b73f41c909f66bded"

		if res := validation.IsMatch(reg, val); res != expect {
			t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
		}
	})

	t.Run("IsMatch(hexadecimal,non-hex)", func(t *testing.T) {
		const expect = false

		reg := `(?i)^[0-9a-f]{32}$`
		val := "Z0fz0c19711bcf3b73f41c909f66bded"

		if res := validation.IsMatch(reg, val); res != expect {
			t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
		}
	})

	// ...

	t.Run("emptiness", func(t *testing.T) {
		t.Run("IsMatch(nil,nil)", func(t *testing.T) {
			const expect = false

			if res := validation.IsMatch(nil, nil); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsMatch(nil,emptyStr)", func(t *testing.T) {
			const expect = false

			if res := validation.IsMatch(nil, ""); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsMatch(emptyStr,nil)", func(t *testing.T) {
			const expect = false

			if res := validation.IsMatch("", nil); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("IsMatch(emptyStr,emptyStr)", func(t *testing.T) {
			const expect = true

			if res := validation.IsMatch("", ""); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

}
