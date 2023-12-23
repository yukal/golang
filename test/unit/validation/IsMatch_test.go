package test

import (
	"testing"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
)

func TestIsMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe("IsMatch", func() {
		g.It("success when value match the mask", func() {
			result := validation.IsMatch(
				`(?i)^[0-9a-f]{32}$`,
				"b0fb0c19711bcf3b73f41c909f66bded")

			g.Assert(result).IsTrue()
		})

		g.It("failure when a value does not match the mask", func() {
			result := validation.IsMatch(
				`(?i)^[0-9a-f]{32}$`,
				"Z0fz0c19711bcf3b73f41c909f66bded")

			g.Assert(result).IsFalse()
		})

		g.It("success when given an empty mask", func() {
			result := validation.IsMatch("", "abra")
			g.Assert(result).IsTrue()
		})

		g.It("failure when given nil instead of mask", func() {
			result := validation.IsEachMatches(nil, "cadabra")
			g.Assert(result).IsFalse()
		})

		g.It(`failure when given args: ("", nil)`, func() {
			g.Assert(validation.IsMatch("", nil)).IsFalse()
		})

		g.It(`failure when given args: (nil, nil)`, func() {
			g.Assert(validation.IsMatch(nil, nil)).IsFalse()
		})
	})

}
