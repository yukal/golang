package test

import (
	"testing"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
)

// go test ./test/unit/validation/...
// go test -v -run TestFilter ./test/unit/validation/...

func TestFilter(t *testing.T) {
	type Person struct {
		Id         uint16 `json:"id"`
		Status     uint8  `json:"status"`
		Role       uint8  `json:"role"`
		LastName   string `json:"lastName"`
		FirstName  string `json:"firstName"`
		Patronymic string `json:"patronymic"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
	}

	g := Goblin(t)

	g.Describe(`Rule "min-fields"`, func() {
		g.It("success", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
				{
					Field: "Status",
					Check: validation.Range{1, 2},
				},
			}

			result := filter.IsValid(Person{
				Id:     1,
				Status: 1,
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
				{
					Field:    "Status",
					Check:    validation.Range{1, 2},
					Optional: true,
				},
				{
					Check: validation.Rule{"min-fields", 2},
				},
			}

			result := filter.IsValid(Person{Id: 1})
			g.Assert(result).IsFalse()
		})
	})

	g.Describe("non-zero", func() {
		g.It("success", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
			}

			result := filter.IsValid(Person{Id: 1})
			g.Assert(result).IsTrue()
		})

		g.It("failure", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
			}

			result := filter.IsValid(Person{})
			g.Assert(result).IsFalse()
		})
	})

	g.Describe("optional", func() {
		g.It("failure", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
				{
					Field:    "Status",
					Check:    validation.Range{1, 2},
					Optional: true,
				},
				{
					Field:    "Role",
					Check:    validation.Range{1, 3},
					Optional: true,
				},
				{
					Field:    "LastName",
					Check:    validation.Range{3, 20},
					Optional: true,
				},
				{
					Field:    "FirstName",
					Check:    validation.Range{3, 20},
					Optional: true,
				},
				{
					Field:    "Patronymic",
					Check:    validation.Range{3, 20},
					Optional: true,
				},
				{
					Field:    "Phone",
					Check:    validation.Rule{"match", `^\d{10}$`},
					Optional: true,
				},
				{
					Field: "Email",
					Check: validation.Group{
						validation.Rule{"max", 40},
						validation.Rule{"match", `(?i)^[\w.+-]+@[\w-]+\.[\w-.]+$`},
					},
					Optional: true,
				},
				{
					Check: validation.Rule{"min-fields", 2},
				},
			}

			result := filter.IsValid(Person{Id: 1})
			g.Assert(result).IsFalse()
		})
	})

	// if result := filter.IsValid(person); res != expect {
	// 	t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
	// }
}
