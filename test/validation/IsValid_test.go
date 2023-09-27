package test

import (
	"testing"
	"time"
	"yu/golang/internal/app/validation"
)

func TestIsValid(t *testing.T) {
	type Article struct {
		Id       uint64
		RegionId uint8
		Hash     string
		Link     string
		Title    string
		Message  string
		Sex      uint8
		Age      uint8
		Height   uint16
		Weight   uint16

		Images []string
		Phones []string

		EmptyArr  [0]string
		FilledArr [3]string

		Map map[string]string

		Date time.Time
	}

	t.Run("min", func(t *testing.T) {
		t.Run("min(1):0", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Id: 0}
			filter := validation.FilterMap{
				"Id": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):1", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Id: 1}
			filter := validation.FilterMap{
				"Id": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):string_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: ""}
			filter := validation.FilterMap{
				"Title": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(4):string_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Title: ""}
			filter := validation.FilterMap{
				"Title": {"min": 4},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):string_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: "text"}
			filter := validation.FilterMap{
				"Title": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(4):string_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: "text"}
			filter := validation.FilterMap{
				"Title": {"min": 4},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):array_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.FilterMap{
				"EmptyArr": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):array_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.FilterMap{
				"EmptyArr": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):array_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{FilledArr: [3]string{"380001234567"}}
			filter := validation.FilterMap{
				"FilledArr": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):array_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{FilledArr: [3]string{"380001234567"}}
			filter := validation.FilterMap{
				"FilledArr": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):slice_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.FilterMap{
				"Phones": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):slice_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Phones: []string{}}
			filter := validation.FilterMap{
				"Phones": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):slice_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{"380001234567"}}
			filter := validation.FilterMap{
				"Phones": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):slice_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{"380001234567"}}
			filter := validation.FilterMap{
				"Phones": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("min(0):map_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.FilterMap{
				"Map": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):map_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Map: map[string]string{}}
			filter := validation.FilterMap{
				"Map": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(0):map_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{"key": "val"}}
			filter := validation.FilterMap{
				"Map": {"min": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("min(1):map_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{"key": "val"}}
			filter := validation.FilterMap{
				"Map": {"min": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

	})

	// ...

	t.Run("max", func(t *testing.T) {
		t.Run("max(100):10", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Id: 10}
			filter := validation.FilterMap{
				"Id": {"max": 100},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(10):100", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Id: 100}
			filter := validation.FilterMap{
				"Id": {"max": 10},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(20):string_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: ""}
			filter := validation.FilterMap{
				"Title": {"max": 20},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(8):string_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: ""}
			filter := validation.FilterMap{
				"Title": {"max": 8},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(20):string_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: "Hello Test!"}
			filter := validation.FilterMap{
				"Title": {"max": 20},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(8):string_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Title: "Hello Test!"}
			filter := validation.FilterMap{
				"Title": {"max": 8},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(3):array_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.FilterMap{
				"EmptyArr": {"max": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):array_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.FilterMap{
				"EmptyArr": {"max": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(3):array_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"FilledArr": {"max": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):array_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"FilledArr": {"max": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(3):slice_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.FilterMap{
				"Phones": {"max": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):slice_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.FilterMap{
				"Phones": {"max": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(3):slice_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"Phones": {"max": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):slice_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"Phones": {"max": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("max(3):map_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.FilterMap{
				"Map": {"max": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):map_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.FilterMap{
				"Map": {"max": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(2):map_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.FilterMap{
				"Map": {"max": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("max(1):map_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.FilterMap{
				"Map": {"max": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("eq", func(t *testing.T) {
		t.Run("eq(0):1", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Id: 1}
			filter := validation.FilterMap{
				"Id": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(1):1", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Id: 1}
			filter := validation.FilterMap{
				"Id": {"eq": 1},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):string_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: ""}
			filter := validation.FilterMap{
				"Title": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(11):string_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Title: ""}
			filter := validation.FilterMap{
				"Title": {"eq": 11},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):string_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Title: "Hello Test!"}
			filter := validation.FilterMap{
				"Title": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(11):string_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Title: "Hello Test!"}
			filter := validation.FilterMap{
				"Title": {"eq": 11},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):array_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.FilterMap{
				"EmptyArr": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):array_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{EmptyArr: [0]string{}}
			filter := validation.FilterMap{
				"EmptyArr": {"eq": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):array_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"FilledArr": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):array_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{FilledArr: [3]string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"FilledArr": {"eq": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):slice_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{}}
			filter := validation.FilterMap{
				"Phones": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):slice_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Phones: []string{}}
			filter := validation.FilterMap{
				"Phones": {"eq": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):slice_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"Phones": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):slice_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Phones: []string{"A", "B", "C"}}
			filter := validation.FilterMap{
				"Phones": {"eq": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		// ...

		t.Run("eq(0):map_empty", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{}}
			filter := validation.FilterMap{
				"Map": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(3):map_empty", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Map: map[string]string{}}
			filter := validation.FilterMap{
				"Map": {"eq": 3},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(0):map_filled", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.FilterMap{
				"Map": {"eq": 0},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("eq(2):map_filled", func(t *testing.T) {
			t.Parallel()
			const expect = true

			article := &Article{Map: map[string]string{
				"key1": "val1",
				"key2": "val2",
			}}

			filter := validation.FilterMap{
				"Map": {"eq": 2},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

	// ...

	t.Run("year", func(t *testing.T) {
		t.Run("same", func(t *testing.T) {
			t.Parallel()
			const expect = true

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			if err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.FilterMap{
				"Date": {"year": 2006},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("diff", func(t *testing.T) {
			t.Parallel()
			const expect = false

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			if err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.FilterMap{
				"Date": {"year": 2024},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("empty_data", func(t *testing.T) {
			t.Parallel()
			const expect = false

			article := &Article{}
			filter := validation.FilterMap{
				"Date": {"year": 2024},
			}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})

		t.Run("empty_filter", func(t *testing.T) {
			t.Parallel()
			const expect = true

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			if err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.FilterMap{}

			if res := filter.IsValid(article); res != expect {
				t.Errorf("Expect( %T(%[1]v) ) => Got( %T(%[2]v) )", expect, res)
			}
		})
	})

}
