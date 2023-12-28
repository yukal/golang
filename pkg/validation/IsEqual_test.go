package validation

import (
	"fmt"
	"testing"
	"yu/golang/internal/app"

	. "github.com/franela/goblin"
)

// go test ./pkg/validation/...
// go test -v -run TestIsEqual ./pkg/validation/...

func TestIsEqual(t *testing.T) {
	g := Goblin(t)

	// ................................
	// int & int

	g.Describe("int & int", func() {

		g.Describe("int8", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_INT8, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_INT8, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_INT32, false},
				// {"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_RUNE, false},
				// {"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_RUNE, false},
				// {"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_RUNE, false},
				// {"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_RUNE, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int16", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_INT16, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_INT16, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int32", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_INT32, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_INT32, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int64", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_INT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_INT64, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_INT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_INT, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_INT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_INT64, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_INT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_INT, true},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})

	// ................................
	// int & uint

	g.Describe("int & uint", func() {

		g.Describe("int8", func() {
			items := [][4]any{
				// {"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_BYTE, false},
				// {"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_BYTE, false},
				// {"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_BYTE, false},
				// {"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_BYTE, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_UINT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_UINT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_UINT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_UINT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT8, app.MIN_UINT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT8, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT8, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT8, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int16", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_UINT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_UINT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_UINT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_UINT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT16, app.MIN_UINT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT16, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT16, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT16, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int32", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_UINT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_UINT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_UINT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_UINT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT32, app.MIN_UINT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT32, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT32, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT32, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int64", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_UINT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_UINT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_UINT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_UINT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT64, app.MIN_UINT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT64, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT64, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT64, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_UINT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_UINT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_UINT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_UINT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_INT, app.MIN_UINT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_INT, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_INT, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_INT, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})

	// ................................
	// uint & uint

	g.Describe("uint & uint", func() {

		g.Describe("uint8", func() {
			items := [][4]any{
				// {"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_BYTE, true},
				// {"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_BYTE, true},
				// {"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_BYTE, false},
				// {"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_BYTE, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_UINT8, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_UINT8, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_UINT16, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_UINT32, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_UINT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_UINT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint16", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_UINT8, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_UINT16, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_UINT16, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_UINT32, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_UINT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_UINT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint32", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_UINT8, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_UINT16, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_UINT32, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_UINT32, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_UINT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_UINT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_UINT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_UINT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint64", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_UINT8, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_UINT16, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_UINT32, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_UINT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_UINT64, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_UINT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_UINT, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_UINT8, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_UINT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_UINT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_UINT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_UINT16, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_UINT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_UINT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_UINT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_UINT32, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_UINT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_UINT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_UINT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_UINT64, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_UINT64, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_UINT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_UINT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_UINT, true},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_UINT, true},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_UINT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})

	// ................................
	// uint & int

	g.Describe("uint & int", func() {

		g.Describe("uint8", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT8, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT8, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT8, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT8, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint16", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT16, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT16, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT16, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT16, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint32", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT32, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT32, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT32, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT32, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint64", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT64, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT64, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT64, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT64, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint", func() {
			items := [][4]any{
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_INT8, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_INT8, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_INT8, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_INT8, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_INT16, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_INT16, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_INT16, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_INT16, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_INT32, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_INT32, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_INT32, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_INT32, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_INT64, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_INT64, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_INT64, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_INT64, false},
				{"min-%[1]T == min-%[2]T", app.MIN_UINT, app.MIN_INT, false},
				{"max-%[1]T == max-%[2]T", app.MAX_UINT, app.MAX_INT, false},
				{"max-%[1]T == min-%[2]T", app.MIN_UINT, app.MAX_INT, false},
				{"min-%[1]T == max-%[2]T", app.MAX_UINT, app.MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsEqual(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})
}
