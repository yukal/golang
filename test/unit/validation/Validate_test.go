package test

import (
	"testing"
	"time"
	"yu/golang/pkg/validation"
	"yu/golang/test"

	. "github.com/franela/goblin"
)

func TestValidate(t *testing.T) {
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

	g := Goblin(t)

	g.Describe("Range (string)", func() {
		g.It("success when given value within the range", func() {
			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Range{10, 20},
				},
			}

			hints := filter.Validate(Article{
				Title: "love is...",
			})

			g.Assert(len(hints)).Equal(0, test.WRONG_LENGTH)
		})

		g.It("failure when given too-short text", func() {
			const expect = "title must contain 8..16 characters"

			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Range{8, 16},
				},
			}

			hints := filter.Validate(Article{
				Title: "smile",
			})

			g.Assert(len(hints)).Equal(1, test.WRONG_LENGTH)
			g.Assert(hints[0]).Equal(expect)
		})

		g.It("failure when given too-long text", func() {
			const expect = "title must contain 5..15 characters"

			filter := validation.Filter{
				{
					Field: "Title",
					Check: validation.Range{5, 15},
				},
			}

			hints := filter.Validate(Article{
				Title: "somebody to love",
			})

			g.Assert(len(hints)).Equal(1, test.WRONG_LENGTH)
			g.Assert(hints[0]).Equal(expect)
		})
	})

	g.Describe("Range (integer)", func() {
		g.It("success when given value within the range", func() {
			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			hints := filter.Validate(Article{
				RegionId: 15,
			})

			g.Assert(len(hints)).Equal(0, test.WRONG_LENGTH)
		})

		g.It("success when given an empty filter", func() {
			filter := validation.Filter{}
			hints := filter.Validate(Article{
				RegionId: 15,
			})

			g.Assert(len(hints)).Equal(0, test.WRONG_LENGTH)
		})

		g.It("failure when given an empty data", func() {
			const expect = "regionId must be in the range 1..25"

			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			hints := filter.Validate(Article{})

			g.Assert(len(hints)).Equal(1, test.WRONG_LENGTH)
			g.Assert(hints[0]).Equal(expect)
		})

		g.It("failure when given below-range value", func() {
			const expect = "regionId must be in the range 1..25"

			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			hints := filter.Validate(Article{
				RegionId: 0,
			})

			g.Assert(len(hints)).Equal(1, test.WRONG_LENGTH)
			g.Assert(hints[0]).Equal(expect)
		})

		g.It("failure when given above-range value", func() {
			const expect = "regionId must be in the range 1..25"

			filter := validation.Filter{
				{
					Field: "RegionId",
					Check: validation.Range{1, 25},
				},
			}

			hints := filter.Validate(Article{
				RegionId: 26,
			})

			g.Assert(len(hints)).Equal(1, test.WRONG_LENGTH)
			g.Assert(hints[0]).Equal(expect)
		})
	})

	g.Describe("Rule (year)", func() {
		g.It("success when the value match", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2006},
				},
			}

			hints := filter.Validate(Article{
				Date: tm,
			})

			g.Assert(len(hints)).Equal(0, test.WRONG_LENGTH)
		})

		g.It("success when given an empty filter", func() {
			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := validation.Filter{}
			hints := filter.Validate(Article{
				Date: tm,
			})

			g.Assert(len(hints)).Equal(0, test.WRONG_LENGTH)
		})

		g.It("failure when given an empty data", func() {
			const expect = "date must be exactly 2024"

			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			hints := filter.Validate(Article{})

			g.Assert(len(hints)).Equal(1)
			g.Assert(hints[0]).Equal(expect)
		})

		g.It("failure when the value is not match", func() {
			const expect = "date must be exactly 2024"

			tm, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
			g.Assert(err).IsNil(err)

			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			hints := filter.Validate(Article{
				Date: tm,
			})

			g.Assert(len(hints)).Equal(1, test.WRONG_LENGTH)
			g.Assert(hints[0]).Equal(expect)
		})
	})

}
