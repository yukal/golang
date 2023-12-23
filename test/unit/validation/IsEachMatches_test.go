package test

import (
	"testing"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
)

func TestIsEachMatches(t *testing.T) {
	g := Goblin(t)

	// TODO: check channels

	g.Describe("slice", func() {
		reg := `(?i)^[0-9a-f]{32}$`

		g.It("success when values match the mask", func() {
			result := validation.IsEachMatches(reg, []string{
				"b0fb0c19711bcf3b73f41c909f66bded",
				"f41c909f66bdedb0fb0c19711bcf3b73",
			})
			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty slice", func() {
			result := validation.IsEachMatches(reg, []string{
				// empty
			})
			g.Assert(result).IsTrue()
		})

		g.It(`success when given an empty mask`, func() {
			result := validation.IsEachMatches(``, []string{"str"})
			g.Assert(result).IsTrue()
		})

		g.It("failure when given nil instead of mask", func() {
			result := validation.IsEachMatches(nil, []string{})
			g.Assert(result).IsFalse()
		})

		g.It("failure when at least 1 value does not match", func() {
			result := validation.IsEachMatches(reg, []string{
				"b0fb0c19711bcf3b73f41c909f66bded",
				"zzz",
			})
			g.Assert(result).IsFalse()
		})

	})

	g.Describe("array", func() {
		reg := `^38[0-9]{10}$`

		g.It("success when values match the mask", func() {
			result := validation.IsEachMatches(reg, [2]string{
				"380001234567",
				"380007654321",
			})

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty array", func() {
			result := validation.IsEachMatches(reg, [0]string{
				// empty
			})
			g.Assert(result).IsTrue()
		})

		g.It(`success when given an empty mask`, func() {
			result := validation.IsEachMatches(``, [1]string{"str"})
			g.Assert(result).IsTrue()
		})

		g.It("failure when given nil instead of mask", func() {
			result := validation.IsEachMatches(nil, [0]string{})
			g.Assert(result).IsFalse()
		})

		g.It("failure when at least 1 value does not match", func() {
			result := validation.IsEachMatches(reg, []string{
				"380001234567",
				"0001234567",
			})
			g.Assert(result).IsFalse()
		})

	})

	g.Describe("map", func() {
		reg := `^https\://img\.domain\.com/[0-9A-Fa-f]{32}\.(?:pn|jpe?)g$`

		g.It("success when values match the mask", func() {
			result := validation.IsEachMatches(reg, map[string]string{
				"img1": "https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
				"img2": "https://img.domain.com/4792592a98f8b9143de71d1db403d163.jpg",
				"img3": "https://img.domain.com/92f2b876b8ea94f711d2173539e73802.png",
			})

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty map", func() {
			result := validation.IsEachMatches(reg, map[string]string{
				// empty
			})
			g.Assert(result).IsTrue()
		})

		g.It(`success when given an empty mask`, func() {
			result := validation.IsEachMatches(``, map[string]string{"k": "v"})
			g.Assert(result).IsTrue()
		})

		g.It("failure when given nil instead of mask", func() {
			result := validation.IsEachMatches(nil, map[string]string{})
			g.Assert(result).IsFalse()
		})

		g.It("failure when at least 1 value does not match", func() {
			result := validation.IsEachMatches(reg, map[string]string{
				"img1": "https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
				"img2": "https://hack.it/animation.gif",
			})
			g.Assert(result).IsFalse()
		})

	})

	g.Describe("emptiness", func() {
		g.It(`failure when given args: ("", nil)`, func() {
			result := validation.IsEachMatches(``, nil)
			g.Assert(result).IsFalse()
		})

		g.It("failure when given args: (nil, nil)", func() {
			result := validation.IsEachMatches(nil, nil)
			g.Assert(result).IsFalse()
		})
	})
}
