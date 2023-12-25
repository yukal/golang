package test

import (
	"bytes"
	"fmt"
	"testing"
	"yu/golang/internal/app"

	. "github.com/franela/goblin"
)

// go test ./test/unit/types...
// go test -v -run TestIsPrimitive ./test/unit/types

func TestIsPrimitive(t *testing.T) {
	g := Goblin(t)

	type TestItem struct {
		In  any
		Out any
	}

	items := []TestItem{
		// int
		{In: app.MIN_INT8, Out: true},
		{In: app.MIN_INT16, Out: true},
		{In: app.MIN_INT32, Out: true},
		{In: app.MIN_INT64, Out: true},
		{In: app.MIN_INT, Out: true},
		{In: app.MIN_RUNE, Out: true},

		//  uint
		{In: app.MIN_UINT8, Out: true},
		{In: app.MIN_UINT16, Out: true},
		{In: app.MIN_UINT32, Out: true},
		{In: app.MIN_UINT64, Out: true},
		{In: app.MIN_UINT, Out: true},
		{In: app.MIN_BYTE, Out: true},

		// float
		{In: app.MIN_FLOAT32, Out: true},
		{In: app.MIN_FLOAT64, Out: true},

		// complex
		{In: app.MIN_COMPLEX64, Out: true},
		{In: app.MIN_COMPLEX128, Out: true},
	}

	g.Describe("IsPrimitive", func() {
		for _, item := range items {
			g.It(fmt.Sprintf("success when given %T", item.In), func() {
				g.Assert(app.IsPrimitive(item.In)).Equal(item.Out)
			})
		}

		g.It("success when given bool", func() {
			g.Assert(app.IsPrimitive(false)).IsTrue()
			g.Assert(app.IsPrimitive(true)).IsTrue()
		})

		g.It("success when given string", func() {
			g.Assert(app.IsPrimitive("")).IsTrue()
		})

		g.It("failure when given struct", func() {
			g.Assert(app.IsPrimitive(struct{}{})).IsFalse()
		})

		g.It("failure when given slice", func() {
			g.Assert(app.IsPrimitive([]any{})).IsFalse()
		})

		g.It("failure when given map", func() {
			g.Assert(app.IsPrimitive(map[any]any{})).IsFalse()
		})

		g.It("failure when given buffer", func() {
			g.Assert(app.IsPrimitive(new(bytes.Buffer))).IsFalse()
		})

		g.It("failure when given function", func() {
			g.Assert(app.IsPrimitive(func() {})).IsFalse()
		})
	})
}
