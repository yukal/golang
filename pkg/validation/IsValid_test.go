package validation

import (
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test ./pkg/validation/...
// go test -v -run TestIsValid ./pkg/validation/...

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
		FilledArr [4]string
		Options   map[string]string

		Date time.Time
	}

	var (
		strEmpty    = ""
		strFilled   = "love"
		arrEmpty    = [0]string{}
		arrFilled   = [4]string{"c", "o", "d", "e"}
		sliceEmpty  = []string{}
		sliceFilled = []string{"t", "e", "s", "t"}
		mapEmpty    = map[string]string{}
		mapFilled   = map[string]string{
			"i": "val1",
			"t": "val2",
			"e": "val3",
			"m": "val4",
		}
	)

	g := Goblin(t)

	g.Describe(`Rule "range" (text)`, func() {
		g.It("success when given value within the range", func() {
			filter := Filter{
				{
					Field: "Title",
					Check: Range{10, 20},
				},
			}

			result := filter.IsValid(Article{
				Title: "love is...",
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when given too-short text", func() {
			filter := Filter{
				{
					Field: "Title",
					Check: Range{8, 16},
				},
			}

			result := filter.IsValid(Article{
				Title: "smile",
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when given too-long text", func() {
			filter := Filter{
				{
					Field: "Title",
					Check: Range{5, 15},
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
			filter := Filter{
				{
					Field: "RegionId",
					Check: Range{1, 25},
				},
			}

			result := filter.IsValid(Article{
				RegionId: 15,
			})

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty filter", func() {
			filter := Filter{}
			result := filter.IsValid(Article{
				RegionId: 15,
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when given an empty data", func() {
			filter := Filter{
				{
					Field: "RegionId",
					Check: Range{1, 25},
				},
			}

			result := filter.IsValid(Article{})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given below-range value", func() {
			filter := Filter{
				{
					Field: "RegionId",
					Check: Range{1, 25},
				},
			}

			result := filter.IsValid(Article{
				RegionId: 0,
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when given above-range value", func() {
			filter := Filter{
				{
					Field: "RegionId",
					Check: Range{1, 25},
				},
			}

			result := filter.IsValid(Article{
				RegionId: 26,
			})

			g.Assert(result).IsFalse()
		})
	})

	g.Describe(`Rule "year"`, func() {
		g.It("success when the value matches a specific year", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2006},
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

			filter := Filter{}
			result := filter.IsValid(Article{
				Date: tm,
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when given an empty data", func() {
			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2024},
				},
			}

			result := filter.IsValid(Article{})

			g.Assert(result).IsFalse()
		})

		g.It("failure when the value is not match", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2024},
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
		g.Describe("numeric", func() {
			g.It("success when the value exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 18},
					},
				}

				result := filter.IsValid(Article{Age: 21})
				g.Assert(result).IsTrue()
			})

			g.It("success when the value reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 18},
					},
				}

				result := filter.IsValid(Article{Age: 18})
				g.Assert(result).IsTrue()
			})

			g.It("success when the value contains an initial zero as min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 0},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 18},
					},
				}

				result := filter.IsValid(Article{Age: 16})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 21},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is greater than the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 2},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length is greater than the min threshold", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"min", 2},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length is greater than the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 2},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length is greater than the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 2},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})
	})

	// ...

	g.Describe(`Rule "max"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Weight",
						Check: Rule{"max", 60},
					},
				}

				result := filter.IsValid(Article{Weight: 50})
				g.Assert(result).IsTrue()
			})

			g.It("success when the value reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Weight",
						Check: Rule{"max", 60},
					},
				}

				result := filter.IsValid(Article{Weight: 60})
				g.Assert(result).IsTrue()
			})

			g.It("success when the value contains an initial zero as max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 0},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("success when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Weight",
						Check: Rule{"max", 60},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Weight",
						Check: Rule{"max", 60},
					},
				}

				result := filter.IsValid(Article{Weight: 65})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 4},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 2},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"max", 4},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"max", 2},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 4},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 2},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 4},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 2},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsFalse()
			})
		})
	})

	// ...

	g.Describe(`Rule "eq"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value equals the expected number", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 21},
					},
				}

				result := filter.IsValid(Article{Age: 21})
				g.Assert(result).IsTrue()
			})

			g.It("success when the value equals the initial zero", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 0},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than the expected number", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 21},
					},
				}

				result := filter.IsValid(Article{Age: 18})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the value is greater than the expected number", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 21},
					},
				}

				result := filter.IsValid(Article{Age: 28})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 21},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length matches a filled string", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length matches an empty string", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 0},
					},
				}

				result := filter.IsValid(Article{Title: strEmpty})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than expected", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Title: strEmpty})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is greater than expected", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("array", func() {
			g.It("success when the length matches a filled array", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length matches an empty array", func() {
				filter := Filter{
					{
						Field: "EmptyArr",
						Check: Rule{"eq", 0},
					},
				}

				result := filter.IsValid(Article{EmptyArr: arrEmpty})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than expected", func() {
				filter := Filter{
					{
						Field: "EmptyArr",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{EmptyArr: arrEmpty})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is greater than expected", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{FilledArr: arrFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("slice", func() {
			g.It("success when the length matches a filled slice", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length matches an empty slice", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 0},
					},
				}

				result := filter.IsValid(Article{Images: sliceEmpty})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than expected", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Images: sliceEmpty})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is greater than expected", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("map", func() {
			g.It("success when the length matches a filled map", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length matches an empty map", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 0},
					},
				}

				result := filter.IsValid(Article{Options: mapEmpty})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than expected", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Options: mapEmpty})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is greater than expected", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when an empty value was passed", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})
	})

}
