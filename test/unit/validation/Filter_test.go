package test

import (
	"testing"
	"yu/golang/pkg/validation"
)

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

	// TODO: Write tests using other validation rules

	t.Run("min-fields", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			const expect = true

			person := Person{Id: 1, Status: 1}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
				{
					Field: "Status",
					Check: validation.Range{1, 2},
				},
				{
					Check: validation.Rule{"min-fields", 2},
				},
			}

			if res := filter.IsValid(person); res != expect {
				t.Errorf("Expect( %v ) => Got( %v )", expect, res)
			}
		})

		t.Run("fail", func(t *testing.T) {
			const expect = false

			person := Person{Id: 1}
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

			if res := filter.IsValid(person); res != expect {
				t.Errorf("Expect( %v ) => Got( %v )", expect, res)
			}
		})
	})

	t.Run("non-zero", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			const expect = true

			person := Person{Id: 1}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
			}

			// if hints := filter.Validate(person); len(hints) > 0 {
			// 	t.Error(hints)
			// }

			if res := filter.IsValid(person); res != expect {
				t.Errorf("Expect( %v ) => Got( %v )", expect, res)
			}
		})

		t.Run("fail", func(t *testing.T) {
			const expect = false

			person := Person{}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.NON_ZERO,
				},
			}

			// if hints := filter.Validate(person); len(hints) > 0 {
			// 	t.Error(hints)
			// }

			if res := filter.IsValid(person); res != expect {
				t.Errorf("Expect( %v ) => Got( %v )", expect, res)
			}
		})
	})

	t.Run("optional", func(t *testing.T) {
		const expect = false

		person := Person{Id: 1}
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

		if res := filter.IsValid(person); res != expect {
			t.Errorf("Expect( %v ) => Got( %v )", expect, res)
		}
	})

	// if res := filter.IsValid(person); res != expect {
	// 	t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
	// }

}
