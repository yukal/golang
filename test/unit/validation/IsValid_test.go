package test

import (
	"testing"
	"time"
	"yu/golang/pkg/validation"

	. "github.com/franela/goblin"
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

	g := Goblin(t)

	g.Describe(`Rule "range" (text)`, func() {
		g.It("success when given value within the range", func() {
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Range{10, 20},
				},
			}

			result := filter.IsValid(Article{
				Title: "love is...",
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when given too-short text", func() {
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Range{8, 16},
				},
			}

			result := filter.IsValid(Article{
				Title: "smile",
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when given too-long text", func() {
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Range{5, 15},
				},
			}

			result := filter.IsValid(Article{
				Title: "somebody to love",
			})

			g.Assert(result).IsFalse()
		})
	})

	g.Describe(`Rule "range" (numeric)`, func() {
		g.It("success when given value within the range", func() {
			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			result := filter.IsValid(Article{
				RegionId: 15,
			})

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty filter", func() {
			filter := validation.Filter{}
			result := filter.IsValid(Article{
				RegionId: 15,
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when given an empty data", func() {
			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			result := filter.IsValid(Article{})
			g.Assert(result).IsFalse()
		})

		g.It("failure when sending below-range value", func() {
			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			result := filter.IsValid(Article{
				RegionId: 0,
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when sending above-range value", func() {
			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			result := filter.IsValid(Article{
				RegionId: 26,
			})

			g.Assert(result).IsFalse()
		})
	})

	g.Describe(`Rule "year"`, func() {
		g.It("success when the value match", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2006},
				},
			}

			result := filter.IsValid(Article{
				Date: tm,
			})

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty filter", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := validation.Filter{}
			result := filter.IsValid(Article{
				Date: tm,
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when given an empty data", func() {
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			result := filter.IsValid(Article{})

			g.Assert(result).IsFalse()
		})

		g.It("failure when the value is not match", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			result := filter.IsValid(Article{
				Date: tm,
			})

			g.Assert(result).IsFalse()
		})
	})

	// ...

	g.Describe(`Rule "min"`, func() {
		g.It("failure on min(1):0", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"min", 1},
				},
			}

			result := filter.IsValid(Article{Id: 0})
			g.Assert(result).IsFalse()
		})

		g.It("success on min(1):1", func() {
			filter := validation.Filter{
				{
					Field: "Id",
					Check: validation.Rule{"min", 1},
				},
			}

			result := filter.IsValid(Article{Id: 1})
			g.Assert(result).IsTrue()
		})

		// ...

		g.It("min(0):string_empty", func() {
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

		g.It("min(4):string_empty", func() {
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

		g.It("min(0):string_filled", func() {
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

		g.It("min(4):string_filled", func() {
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

		g.It("min(0):array_empty", func() {
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

		g.It("min(1):array_empty", func() {
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

		g.It("min(0):array_filled", func() {
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

		g.It("min(1):array_filled", func() {
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

		g.It("min(0):slice_empty", func() {
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

		g.It("min(1):slice_empty", func() {
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

		g.It("min(0):slice_filled", func() {
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

		g.It("min(1):slice_filled", func() {
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

		g.It("min(0):map_empty", func() {
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

		g.It("min(1):map_empty", func() {
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

		g.It("min(0):map_filled", func() {
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

		g.It("min(1):map_filled", func() {
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

	g.Describe(`Rule "max"`, func() {
		g.It("max(100):10", func() {
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

		g.It("max(10):100", func() {
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

		g.It("max(20):string_empty", func() {
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

		g.It("max(8):string_empty", func() {
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

		g.It("max(20):string_filled", func() {
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

		g.It("max(8):string_filled", func() {
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

		g.It("max(3):array_empty", func() {
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

		g.It("max(2):array_empty", func() {
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

		g.It("max(3):array_filled", func() {
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

		g.It("max(2):array_filled", func() {
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

		g.It("max(3):slice_empty", func() {
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

		g.It("max(2):slice_empty", func() {
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

		g.It("max(3):slice_filled", func() {
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

		g.It("max(2):slice_filled", func() {
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

		g.It("max(3):map_empty", func() {
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

		g.It("max(2):map_empty", func() {
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

		g.It("max(2):map_filled", func() {
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

		g.It("max(1):map_filled", func() {
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

	g.Describe(`Rule "eq"`, func() {
		g.It("eq(0):1", func() {
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

		g.It("eq(1):1", func() {
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

		g.It("eq(0):string_empty", func() {
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

		g.It("eq(11):string_empty", func() {
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

		g.It("eq(0):string_filled", func() {
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

		g.It("eq(11):string_filled", func() {
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

		g.It("eq(0):array_empty", func() {
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

		g.It("eq(3):array_empty", func() {
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

		g.It("eq(0):array_filled", func() {
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

		g.It("eq(3):array_filled", func() {
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

		g.It("eq(0):slice_empty", func() {
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

		g.It("eq(3):slice_empty", func() {
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

		g.It("eq(0):slice_filled", func() {
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

		g.It("eq(3):slice_filled", func() {
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

		g.It("eq(0):map_empty", func() {
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

		g.It("eq(3):map_empty", func() {
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

		g.It("eq(0):map_filled", func() {
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

		g.It("eq(2):map_filled", func() {
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

}
