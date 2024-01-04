package validation

import (
	"testing"

	. "github.com/franela/goblin"
)

// go test ./pkg/validation/...
// go test -v -run TestFilter ./pkg/validation/...

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

	g.Describe(`Rule "minFields"`, func() {
		g.It("success", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
				{
					Field: "Status",
					Check: Range{1, 2},
				},
			}

			result := filter.IsValid(Person{
				Id:     1,
				Status: 1,
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
				{
					Field:    "Status",
					Check:    Range{1, 2},
					Optional: true,
				},
				{
					Check: Rule{"minFields", 2},
				},
			}

			result := filter.IsValid(Person{Id: 1})
			g.Assert(result).IsFalse()
		})
	})

	g.Describe(`Rule "NON_ZERO"`, func() {
		g.It("success", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Person{Id: 1})
			g.Assert(result).IsTrue()
		})

		g.It("failure", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Person{})
			g.Assert(result).IsFalse()
		})
	})

	g.Describe("optional", func() {
		g.It("failure", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
				{
					Field:    "Status",
					Check:    Range{1, 2},
					Optional: true,
				},
				{
					Field:    "Role",
					Check:    Range{1, 3},
					Optional: true,
				},
				{
					Field:    "LastName",
					Check:    Range{3, 20},
					Optional: true,
				},
				{
					Field:    "FirstName",
					Check:    Range{3, 20},
					Optional: true,
				},
				{
					Field:    "Patronymic",
					Check:    Range{3, 20},
					Optional: true,
				},
				{
					Field:    "Phone",
					Check:    Rule{"match", `^\d{10}$`},
					Optional: true,
				},
				{
					Field: "Email",
					Check: Group{
						Rule{"max", 40},
						Rule{"match", `(?i)^[\w.+-]+@[\w-]+\.[\w-.]+$`},
					},
					Optional: true,
				},
				{
					Check: Rule{"minFields", 2},
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
