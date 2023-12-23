package test

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"yu/golang/internal/app"
	"yu/golang/pkg/validation"
)

func TestValidate(t *testing.T) {
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

	// TODO: Write tests using other validation rules

	t.Run("range", func(t *testing.T) {
		filter := validation.Filter{
			{
				Field: "RegionId",
				Check: validation.Range{1, 25},
			},
		}

		t.Run("string", func(t *testing.T) {
			t.Run("success within range", func(t *testing.T) {
				expect := []string{}

				filter := validation.Filter{
					{
						Field: "Title",
						Check: validation.Range{10, 20},
					},
				}

				hints := filter.Validate(Article{
					Title: "love is...",
				})

				if diffs := app.SliceDifference(hints, expect); len(diffs) > 0 {
					t.Errorf("Expect( %v ) => Got( %v )", expect, diffs)
				}
			})

			t.Run("fail with short text", func(t *testing.T) {
				expect := []string{"Title must contain 8..16 characters"}

				filter := validation.Filter{
					{
						Field: "Title",
						Check: validation.Range{8, 16},
					},
				}

				hints := filter.Validate(Article{
					Title: "smile",
				})

				if diffs := app.SliceDifference(hints, expect); len(diffs) > 0 {
					t.Errorf("Expect( %v ) => Got( %v )", expect, diffs)
				}
			})

			t.Run("fail with long text", func(t *testing.T) {
				expect := []string{"Title must contain 5..15 characters"}

				filter := validation.Filter{
					{
						Field: "Title",
						Check: validation.Range{5, 15},
					},
				}

				hints := filter.Validate(Article{
					Title: "somebody to love",
				})

				if diffs := app.SliceDifference(hints, expect); len(diffs) > 0 {
					t.Errorf("Expect( %v ) => Got( %v )", expect, diffs)
				}
			})
		})

		t.Run("int", func(t *testing.T) {
			t.Run("success on first pos", func(t *testing.T) {
				article := Article{RegionId: 1}

				if hints := filter.Validate(article); len(hints) > 0 {
					t.Error(strings.Join(hints, ",\n"))
				}
			})

			t.Run("success on last pos", func(t *testing.T) {
				article := Article{RegionId: 25}

				if hints := filter.Validate(article); len(hints) > 0 {
					t.Error(strings.Join(hints, ",\n"))
				}
			})

			t.Run("fail on below the permissible limit", func(t *testing.T) {
				expect := []string{"RegionId must be in the range 1..25"}

				hints := filter.Validate(Article{
					RegionId: 0,
				})

				if diffs := app.SliceDifference(hints, expect); len(diffs) > 0 {
					t.Errorf("Expect( %v ) => Got( %v )", expect, diffs)
				}
			})

			t.Run("fail on above the permissible limit", func(t *testing.T) {
				expect := []string{"RegionId must be in the range 1..25"}

				hints := filter.Validate(Article{
					RegionId: 26,
				})

				if diffs := app.SliceDifference(hints, expect); len(diffs) > 0 {
					t.Errorf("Expect( %v ) => Got( %v )", expect, diffs)
				}
			})
		})
	})

	t.Run("year", func(t *testing.T) {
		t.Run("same", func(t *testing.T) {
			var (
				tm  time.Time
				err error
			)

			if tm, err = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"); err != nil {
				t.Error(err)
			}

			article := Article{Date: tm}
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2006},
				},
			}

			if hints := filter.Validate(article); len(hints) > 0 {
				t.Error(strings.Join(hints, ",\n"))
			}
		})

		t.Run("diff", func(t *testing.T) {
			var (
				expectMsg = fmt.Sprintf("Date "+validation.MsgEq, 2024)
				tm        time.Time
				err       error
			)

			if tm, err = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"); err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			if hints := filter.Validate(article); len(hints) > 0 {
				if hints[0] != expectMsg {
					t.Errorf("Expect( %v ) => Got( %v )", expectMsg, hints[0])
				}
			}
		})

		t.Run("empty_data", func(t *testing.T) {
			expectMsg := fmt.Sprintf("Date "+validation.MsgEq, 2024)

			article := &Article{}
			filter := validation.Filter{
				{
					Field: "Date",
					Check: validation.Rule{"year", 2024},
				},
			}

			if hints := filter.Validate(article); len(hints) > 0 {
				if hints[0] != expectMsg {
					t.Errorf("Expect( %v ) => Got( %v )", expectMsg, hints[0])
				}
			}
		})

		t.Run("empty_filter", func(t *testing.T) {
			var (
				expectHints = 0
				tm          time.Time
				err         error
			)

			if tm, err = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"); err != nil {
				t.Error(err)
			}

			article := &Article{Date: tm}
			filter := validation.Filter{}

			if hints := filter.Validate(article); len(hints) > expectHints {
				t.Errorf("Expect( %v ) => Got( %v )", expectHints, len(hints))
			}
		})
	})

}
