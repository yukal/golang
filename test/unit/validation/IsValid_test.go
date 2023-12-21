package test

import (
	"testing"
	"time"
	"yu/golang/pkg/validation"
)

func TestIsValid(t *testing.T) {
	type Article struct {
		Id       uint64 `json:"id"`
		RegionId uint8  `json:"regionId"`
		Hash     string `json:"hash"`
		Link     string `json:"link"`
		Title    string `json:"title"`
		Message  string `json:"message"`
		Sex      uint8  `json:"sex"`
		Age      uint8  `json:"age"`
		Height   uint16 `json:"height"`
		Weight   uint16 `json:"weight"`

		Images []string `json:"images"`
		Phones []string `json:"phones"`

		EmptyArr  [0]string `json:"emptyArr"`
		FilledArr [3]string `json:"filledArr"`

		Map map[string]string `json:"map"`

		Date time.Time `json:"date"`
	}

	t.Run("min", func(t *testing.T) {
		t.Run("min(1):0", func(t *testing.T) {
			const expect = false

			article := &Article{Id: 0}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):1", func(t *testing.T) {
			const expect = true

			article := &Article{Id: 1}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):string_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Title: ""}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(4):string_empty", func(t *testing.T) {
			const expect = false

			article := &Article{Title: ""}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"min", 4},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):string_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Title: "text"}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(4):string_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Title: "text"}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"min", 4},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):array_empty", func(t *testing.T) {
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.Filter{
				{
					Field: "EmptyArr",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):array_empty", func(t *testing.T) {
			const expect = false

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.Filter{
				{
					Field: "EmptyArr",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):array_filled", func(t *testing.T) {
			const expect = true

			article := &Article{FilledArr: [3]string{"380001234567"}}
			filter := validation.Filter{
				{
					Field: "FilledArr",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):array_filled", func(t *testing.T) {
			const expect = true

			article := &Article{FilledArr: [3]string{"380001234567"}}
			filter := validation.Filter{
				{
					Field: "FilledArr",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):slice_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):slice_empty", func(t *testing.T) {
			const expect = false

			article := &Article{Phones: []string{}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):slice_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{"380001234567"}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):slice_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{"380001234567"}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):map_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):map_empty", func(t *testing.T) {
			const expect = false

			article := &Article{Map: map[string]string{}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):map_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{"key": "val"}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"min", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):map_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{"key": "val"}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"min", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

	})

	// ...

	t.Run("max", func(t *testing.T) {
		t.Run("max(100):10", func(t *testing.T) {
			const expect = true

			article := &Article{Id: 10}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"max", 100},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(10):100", func(t *testing.T) {
			const expect = false

			article := &Article{Id: 100}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"max", 10},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(20):string_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Title: ""}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"max", 20},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(8):string_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Title: ""}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"max", 8},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(20):string_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Title: "Hello Test!"}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"max", 20},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(8):string_filled", func(t *testing.T) {
			const expect = false

			article := &Article{Title: "Hello Test!"}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"max", 8},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(3):array_empty", func(t *testing.T) {
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.Filter{
				{
					Field: "EmptyArr",
					Check: validation.Rule{"max", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):array_empty", func(t *testing.T) {
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.Filter{
				{
					Field: "EmptyArr",
					Check: validation.Rule{"max", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(3):array_filled", func(t *testing.T) {
			const expect = true

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "FilledArr",
					Check: validation.Rule{"max", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):array_filled", func(t *testing.T) {
			const expect = false

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "FilledArr",
					Check: validation.Rule{"max", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(3):slice_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"max", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):slice_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"max", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(3):slice_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"max", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):slice_filled", func(t *testing.T) {
			const expect = false

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"max", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(3):map_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"max", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):map_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"max", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):map_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"max", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(1):map_filled", func(t *testing.T) {
			const expect = false

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"max", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("eq", func(t *testing.T) {
		t.Run("eq(0):1", func(t *testing.T) {
			const expect = false

			article := &Article{Id: 1}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(1):1", func(t *testing.T) {
			const expect = true

			article := &Article{Id: 1}
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"eq", 1},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):string_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Title: ""}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(11):string_empty", func(t *testing.T) {
			const expect = false

			article := &Article{Title: ""}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"eq", 11},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):string_filled", func(t *testing.T) {
			const expect = false

			article := &Article{Title: "Hello Test!"}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(11):string_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Title: "Hello Test!"}
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Rule{"eq", 11},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):array_empty", func(t *testing.T) {
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.Filter{
				{
					Field: "EmptyArr",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):array_empty", func(t *testing.T) {
			const expect = false

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.Filter{
				{
					Field: "EmptyArr",
					Check: validation.Rule{"eq", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):array_filled", func(t *testing.T) {
			const expect = false

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "FilledArr",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):array_filled", func(t *testing.T) {
			const expect = true

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "FilledArr",
					Check: validation.Rule{"eq", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):slice_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):slice_empty", func(t *testing.T) {
			const expect = false

			article := &Article{Phones: []string{}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"eq", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):slice_filled", func(t *testing.T) {
			const expect = false

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):slice_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.Filter{
				{
					Field: "Phones",
					Check: validation.Rule{"eq", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):map_empty", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):map_empty", func(t *testing.T) {
			const expect = false

			article := &Article{Map: map[string]string{}}
			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"eq", 3},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):map_filled", func(t *testing.T) {
			const expect = false

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"eq", 0},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(2):map_filled", func(t *testing.T) {
			const expect = true

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.Filter{
				{
					Field: "Map",
					Check: validation.Rule{"eq", 2},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("year", func(t *testing.T) {
		t.Run("same", func(t *testing.T) {
			const expect = true

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			if err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2006},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("diff", func(t *testing.T) {
			const expect = false

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			if err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("empty_data", func(t *testing.T) {
			const expect = false

			article := &Article{}
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("empty_filter", func(t *testing.T) {
			const expect = true

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			if err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.Filter{}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

}
