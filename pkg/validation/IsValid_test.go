package validation

import (
	"math"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test ./pkg/validation/...
// go test -v -run TestIsValid ./pkg/validation/...
//
// go test -v -run TestIsValidEq ./pkg/validation/...
// go test -v -run TestIsValidMin ./pkg/validation/...
// go test -v -run TestIsValidMax ./pkg/validation/...
// go test -v -run TestIsValidRange ./pkg/validation/...
// go test -v -run TestIsValidYear ./pkg/validation/...
// go test -v -run TestIsValidMatch ./pkg/validation/...
// go test -v -run TestIsValidEachEq ./pkg/validation/...
// go test -v -run TestIsValidEachMin ./pkg/validation/...
// go test -v -run TestIsValidEachMax ./pkg/validation/...
// go test -v -run TestIsValidEachMatch ./pkg/validation/...
// go test -v -run TestIsValidMinFields ./pkg/validation/...
// go test -v -run TestIsValidNonZero ./pkg/validation/...

// go test ./pkg/validation/...
// go test -v -run TestIsValidEq ./pkg/validation/...

func TestIsValidEq(t *testing.T) {
	type Article struct {
		Title   string
		Age     uint8
		Images  []string
		Phones  [4]string
		Options map[int]string
	}

	g := Goblin(t)

	g.Describe(`Rule "eq" (equal)`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value equals a threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 18},
					},
				}

				result := filter.IsValid(Article{Age: 18})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 21},
					},
				}

				result := filter.IsValid(Article{Age: 18})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 18},
					},
				}

				result := filter.IsValid(Article{Age: 21})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("string", func() {
			strFilled := "all you need is love"

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 20},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 30},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the length exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 10},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"eq", 8},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 8},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "We all live in a yellow submarine",
				2: "While My Guitar Gently Weeps",
				3: "All you need is love",
				4: "Let it be",
			}

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 4},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 8},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 2},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`emptiness`, func() {
			fieldsToCheck := []string{"Age", "Title", "Images", "Options"}
			article := Article{
				Age:    21,
				Title:  "All you need is love",
				Phones: [4]string{"0001234567"},
				Images: []string{"img1", "img2"},
				Options: map[int]string{
					1: "one",
					2: "two",
				},
			}

			g.It("success when given a zero proto and empty value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"eq", 0},
						},
					}

					result := filter.IsValid(Article{})
					g.Assert(result).IsTrue(fieldName)
				}
			})

			g.It("failure when missing rule value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"eq"},
						},
					}

					result := filter.IsValid(article)
					g.Assert(result).IsFalse()
				}
			})

			g.It("failure when given an empty rule", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{},
						},
					}

					result := filter.IsValid(article)
					g.Assert(result).IsFalse()
				}
			})
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidMin ./pkg/validation/...

func TestIsValidMin(t *testing.T) {
	type Article struct {
		Title   string
		Age     uint8
		Images  []string
		Phones  [4]string
		Options map[int]string
	}

	g := Goblin(t)

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
						Check: Rule{"min", 21},
					},
				}

				result := filter.IsValid(Article{Age: 21})
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
		})

		// ...

		g.Describe("string", func() {
			strFilled := "all you need is love"

			g.It("success when the length exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 10},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 20},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 30},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"min", 2},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"min", 4},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length exceeds the min threshold", func() {
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

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{Images: sliceFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "We all live in a yellow submarine",
				2: "While My Guitar Gently Weeps",
				3: "All you need is love",
				4: "Let it be",
			}

			g.It("success when the length exceeds the min threshold", func() {
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

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 8},
					},
				}

				result := filter.IsValid(Article{Options: mapFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`emptiness`, func() {
			fieldsToCheck := []string{"Age", "Title", "Images", "Options"}
			article := Article{
				Age:    21,
				Title:  "All you need is love",
				Phones: [4]string{"0001234567"},
				Images: []string{"img1", "img2"},
				Options: map[int]string{
					1: "one",
					2: "two",
				},
			}

			g.It("success when given a zero proto and empty value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"min", 0},
						},
					}

					result := filter.IsValid(Article{})
					g.Assert(result).IsTrue(fieldName)
				}
			})

			g.It("failure when missing rule value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"min"},
						},
					}

					result := filter.IsValid(article)
					g.Assert(result).IsFalse()
				}
			})

			g.It("failure when given an empty rule", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{},
						},
					}

					result := filter.IsValid(article)
					g.Assert(result).IsFalse()
				}
			})
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidMax ./pkg/validation/...

func TestIsValidMax(t *testing.T) {
	type Article struct {
		Title   string
		Age     uint8
		Images  []string
		Phones  [4]string
		Options map[int]string
	}

	g := Goblin(t)

	g.Describe(`Rule "max"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 21},
					},
				}

				result := filter.IsValid(Article{Age: 18})
				g.Assert(result).IsTrue()
			})

			g.It("success when the value reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 21},
					},
				}

				result := filter.IsValid(Article{Age: 21})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the value exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 12},
					},
				}

				result := filter.IsValid(Article{Age: 21})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("string", func() {
			strFilled := "all you need is love"

			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 30},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 20},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 10},
					},
				}

				result := filter.IsValid(Article{Title: strFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"max", 8},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"max", 4},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"max", 2},
					},
				}

				result := filter.IsValid(Article{Phones: arrFilled})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

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
			mapFilled := map[int]string{
				1: "We all live in a yellow submarine",
				2: "While My Guitar Gently Weeps",
				3: "All you need is love",
				4: "Let it be",
			}

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

		// ...

		g.Describe(`emptiness`, func() {
			fieldsToCheck := []string{"Age", "Title", "Images", "Options"}
			article := Article{
				Age:    21,
				Title:  "All you need is love",
				Phones: [4]string{"0001234567"},
				Images: []string{"img1", "img2"},
				Options: map[int]string{
					1: "one",
					2: "two",
				},
			}

			g.It("success when given a zero proto and empty value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"max", 0},
						},
					}

					result := filter.IsValid(Article{})
					g.Assert(result).IsTrue(fieldName)
				}
			})

			g.It("failure when missing rule proto", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"max"},
						},
					}

					result := filter.IsValid(article)
					g.Assert(result).IsFalse()
				}
			})

			g.It("failure when given an empty rule", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{},
						},
					}

					result := filter.IsValid(article)
					g.Assert(result).IsFalse()
				}
			})
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidRange ./pkg/validation/...

func TestIsValidRange(t *testing.T) {
	type Article struct {
		// Id       uint64
		// RegionId uint8
		// Hash     string
		// Link     string
		Title string
		// Message  string
		// Sex      uint8
		Age uint8
		// Height   uint16
		// Weight   uint16

		Images []string
		// Phones []string

		// EmptyArr  [0]string
		FilledArr [4]string
		Options   map[string]string

		// Date time.Time
	}

	g := Goblin(t)

	g.Describe(`Rule "range"`, func() {
		g.Describe(`numeric`, func() {
			g.It("success when the value matches the range", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				result := filter.IsValid(Article{Age: 18})
				g.Assert(result).IsTrue()
			})

			g.It("failure when given below-range value", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				result := filter.IsValid(Article{Age: 16})
				g.Assert(result).IsFalse()
			})

			g.It("failure when given above-range value", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				result := filter.IsValid(Article{Age: 31})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`string`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Range{4, 20},
					},
				}

				result := filter.IsValid(Article{Title: "all you need is love"})
				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Range{25, 45},
					},
				}

				result := filter.IsValid(Article{Title: "all you need is love"})
				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Range{3, 18},
					},
				}

				result := filter.IsValid(Article{Title: "all you need is love"})
				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`array`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Range{1, 4},
					},
				}

				result := filter.IsValid(Article{
					FilledArr: [4]string{"t", "e", "s", "t"},
				})

				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Range{10, 80},
					},
				}

				result := filter.IsValid(Article{
					FilledArr: [4]string{"t", "e", "s", "t"},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "FilledArr",
						Check: Range{1, 3},
					},
				}

				result := filter.IsValid(Article{
					FilledArr: [4]string{"t", "e", "s", "t"},
				})

				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Range{1, 4},
					},
				}

				result := filter.IsValid(Article{
					Images: []string{"jpeg", "jpg", "png", "gif"},
				})

				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Range{10, 80},
					},
				}

				result := filter.IsValid(Article{
					Images: []string{"jpeg", "jpg", "png", "gif"},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Range{1, 3},
					},
				}

				result := filter.IsValid(Article{
					Images: []string{"jpeg", "jpg", "png", "gif"},
				})

				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Range{1, 4},
					},
				}

				result := filter.IsValid(Article{
					Options: map[string]string{
						"jpeg": "image/jpeg",
						"jpg":  "image/jpeg",
						"png":  "image/png",
						"gif":  "image/gif",
					},
				})

				g.Assert(result).IsTrue()
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Range{10, 80},
					},
				}

				result := filter.IsValid(Article{
					Options: map[string]string{
						"jpeg": "image/jpeg",
						"jpg":  "image/jpeg",
						"png":  "image/png",
						"gif":  "image/gif",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Range{1, 3},
					},
				}

				result := filter.IsValid(Article{
					Options: map[string]string{
						"jpeg": "image/jpeg",
						"jpg":  "image/jpeg",
						"png":  "image/png",
						"gif":  "image/gif",
					},
				})

				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure if at least 1 element of the range is empty", func() {
				protos := []Range{
					{},
					{nil, 15},
					{15, nil},
				}

				for _, item := range protos {
					filter := Filter{
						{
							Field: "Age",
							Check: Range{item[0], item[1]},
						},
					}

					result := filter.IsValid(Article{Age: 31})
					g.Assert(result).IsFalse()
				}
			})

			g.It("failure when given an empty value", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				result := filter.IsValid(Article{})
				g.Assert(result).IsFalse()
			})
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidYear ./pkg/validation/...

func TestIsValidYear(t *testing.T) {
	type Article struct {
		Date time.Time
	}

	g := Goblin(t)

	g.Describe(`Rule "year"`, func() {
		g.It("success when the value matches a specific year", func() {
			tm, err := time.Parse(time.RFC3339, "2024-12-25T16:04:05Z")
			g.Assert(err).IsNil(err)

			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2024},
				},
			}

			result := filter.IsValid(Article{Date: tm})
			g.Assert(result).IsTrue()
		})

		g.It("failure when the value is not match", func() {
			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2024},
				},
			}

			result := filter.IsValid(Article{
				Date: *new(time.Time),
			})

			g.Assert(result).IsFalse()
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidMatch ./pkg/validation/...

func TestIsValidMatch(t *testing.T) {
	type Article struct {
		Hash string `json:"hash"`
	}

	g := Goblin(t)

	g.Describe(`Rule "match"`, func() {
		g.It("success when the value matches the mask", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", `(?i)^[0-9a-f]{32}$`},
				},
			}

			result := filter.IsValid(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty mask", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", ``},
				},
			}

			result := filter.IsValid(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when the value does not match the mask", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", `(?i)^[0-9a-f]{32}$`},
				},
			}

			result := filter.IsValid(Article{
				Hash: "Z0zZ0z19711zZz3z73z41z909z66zZzZ",
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when missing rule value", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match"},
				},
			}

			result := filter.IsValid(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when given an empty rule", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{},
				},
			}

			result := filter.IsValid(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(result).IsFalse()
		})

		g.It("failure when given an empty value", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", `(?i)^[0-9a-f]{32}$`},
				},
			}

			result := filter.IsValid(Article{})
			g.Assert(result).IsFalse(result)
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidEachEq ./pkg/validation/...

func TestIsValidEachEq(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eachEq" (equal)`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int
				Bands   [2]string
				Artists [2][2]string
				Songs   [2][]string
				Albums  [2]map[int]string
			}

			type Emptyness struct {
				Pages   [0]int
				Bands   [0]string
				Artists [0][0]string
				Songs   [0][]string
				Albums  [0]map[int]string
			}

			g.Describe("array:numeric", func() {
				g.It("success when the element value equals a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{15, 15},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 10},
						},
					}

					result := filter.IsValid(Emptyness{
						Pages: [0]int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the value is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{5, 10},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the value exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{25, 30},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:string", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Emptyness{
						Bands: [0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 10},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:array", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Emptyness{
						Artists: [0][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:slice", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Emptyness{
						Songs: [0][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})
					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:map", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
								// 1975: "Toys in the Attic",
								// 1976: "Rocks",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Emptyness{
						Albums: [0]map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int
				Bands   []string
				Artists [][2]string
				Artist  [][0]string
				Songs   [][]string
				Albums  []map[int]string
			}

			g.Describe("slice:numeric", func() {
				g.It("success when the element value equals a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{15, 15},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 10},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the value is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{5, 10},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the value exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{25, 30},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:string", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 10},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:array", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Slice{
						Artist: [][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:slice", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})
					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:map", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
								// 1975: "Toys in the Attic",
								// 1976: "Rocks",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int
				Bands   map[int]string
				Artists map[string][2]string
				Artist  map[string][0]string
				Songs   map[string][]string
				Albums  map[string]map[int]string
			}

			g.Describe("map:numeric", func() {
				g.It("success when the element value equals a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 15,
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 10},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the value is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    5,
							"Prologue": 15,
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the value exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachEq", 15},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 30,
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:string", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Aerosmith",
							2: "Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 9},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Metallica",
							2: "Nirvana",
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachEq", 10},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Led Zeppelin",
							2: "Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:array", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Aerosmith": {
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Map{
						Artist: map[string][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Nirvana": {
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							"Metallica": {
								"Lars Ulrich",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:slice", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Aerosmith": {
								"Walk This Way",
								"Dream On",
							},
							"Scorpions": {
								"Rock You Like a Hurricane",
								"Still Loving You",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Led Zeppelin": {
								"Kashmir",
								"Stairway To Heaven",
							},
							"Pink Floyd": {
								"Wish You Were Here",
							},
						},
					})
					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Nirvana": {
								"Lithium",
								"Smells Like Teen Spirit",
							},
							"Metallica": {
								"One",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:map", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 4},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 2},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachEq", 1},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidEachMin ./pkg/validation/...

func TestIsValidEachMin(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eachMin"`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int
				Bands   [2]string
				Artists [2][2]string
				Songs   [2][]string
				Albums  [2]map[int]string
			}

			type Emptyness struct {
				Pages   [0]int
				Bands   [0]string
				Artists [0][0]string
				Songs   [0][]string
				Albums  [0]map[int]string
			}

			g.Describe("array:numeric", func() {
				g.It("success when the element value exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element value reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 15},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Emptyness{
						Pages: [0]int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{5, 15},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:string", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 7},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 9},
						},
					}

					result := filter.IsValid(Emptyness{
						Bands: [0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 12},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:array", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Emptyness{
						Artists: [0][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:slice", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Emptyness{
						Songs: [0][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:map", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Emptyness{
						Albums: [0]map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int
				Bands   []string
				Artists [][2]string
				Artist  [][0]string
				Songs   [][]string
				Albums  []map[int]string
			}

			g.Describe("slice:numeric", func() {
				g.It("success when the element value exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element value reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 15},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{5, 15},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:string", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 7},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 9},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 12},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:array", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Slice{
						Artist: [][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:slice", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:map", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int
				Bands   map[int]string
				Artists map[string][2]string
				Artist  map[string][0]string
				Songs   map[string][]string
				Albums  map[string]map[int]string
			}

			g.Describe("map:numeric", func() {
				g.It("success when the element value exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element value reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 15},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMin", 10},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    5,
							"Prologue": 15,
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:string", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Aerosmith",
							2: "Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 7},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Metallica",
							2: "Nirvana",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 9},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMin", 12},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Led Zeppelin",
							2: "Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:array", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Aerosmith": {
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Map{
						Artist: map[string][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Nirvana": {
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							"Metallica": {
								"Lars Ulrich",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:slice", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Aerosmith": {
								"Walk This Way",
								"Dream On",
							},
							"Scorpions": {
								"Rock You Like a Hurricane",
								"Still Loving You",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Led Zeppelin": {
								"Kashmir",
								"Stairway To Heaven",
							},
							"Pink Floyd": {
								"Wish You Were Here",
								"High Hopes",
							},
						},
					})
					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Nirvana": {
								"Lithium",
								"Smells Like Teen Spirit",
							},
							"Metallica": {
								"One",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:map", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 1},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 2},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 3},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMin", 4},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidEachMax ./pkg/validation/...

func TestIsValidEachMax(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eachMax"`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int
				Bands   [2]string
				Artists [2][2]string
				Songs   [2][]string
				Albums  [2]map[int]string
			}

			type Emptyness struct {
				Pages   [0]int
				Bands   [0]string
				Artists [0][0]string
				Songs   [0][]string
				Albums  [0]map[int]string
			}

			g.Describe("array:numeric", func() {
				g.It("success when the element value is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 30},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element value reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 25},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 15},
						},
					}

					result := filter.IsValid(Emptyness{
						Pages: [0]int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 10},
						},
					}

					result := filter.IsValid(Array{
						Pages: [2]int{5, 15},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:string", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 18},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 9},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 9},
						},
					}

					result := filter.IsValid(Emptyness{
						Bands: [0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 10},
						},
					}

					result := filter.IsValid(Array{
						Bands: [2]string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:array", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Emptyness{
						Artists: [0][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Array{
						Artists: [2][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:slice", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Emptyness{
						Songs: [0][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("array:map", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Emptyness{
						Albums: [0]map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int
				Bands   []string
				Artists [][2]string
				Artist  [][0]string
				Songs   [][]string
				Albums  []map[int]string
			}

			g.Describe("slice:numeric", func() {
				g.It("success when the element value is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 30},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element value reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 25},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 15},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 10},
						},
					}

					result := filter.IsValid(Slice{
						Pages: []int{5, 15},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:string", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 18},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 9},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 9},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 10},
						},
					}

					result := filter.IsValid(Slice{
						Bands: []string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:array", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Artist: [][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Slice{
						Artists: [][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:slice", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("slice:map", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int
				Bands   map[int]string
				Artists map[string][2]string
				Artist  map[string][0]string
				Songs   map[string][]string
				Albums  map[string]map[int]string
			}

			g.Describe("map:numeric", func() {
				g.It("success when the element value is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 30},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element value reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 25},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 15},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"eachMax", 10},
						},
					}

					result := filter.IsValid(Map{
						Pages: map[string]int{
							"Title":    5,
							"Prologue": 15,
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:string", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 18},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Aerosmith",
							2: "Scorpions",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 9},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Metallica",
							2: "Nirvana",
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 9},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"eachMax", 10},
						},
					}

					result := filter.IsValid(Map{
						Bands: map[int]string{
							1: "Led Zeppelin",
							2: "Pink Floyd",
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:array", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Aerosmith": {
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Artist: map[string][0]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Map{
						Artists: map[string][2]string{
							"Nirvana": {
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							"Metallica": {
								"Lars Ulrich",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:slice", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Aerosmith": {
								"Walk This Way",
								"Dream On",
							},
							"Scorpions": {
								"Rock You Like a Hurricane",
								"Still Loving You",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Led Zeppelin": {
								"Kashmir",
								"Stairway To Heaven",
							},
							"Pink Floyd": {
								"Wish You Were Here",
								"High Hopes",
							},
						},
					})
					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Map{
						Songs: map[string][]string{
							"Nirvana": {
								"Lithium",
								"Smells Like Teen Spirit",
							},
							"Metallica": {
								"One",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})

			// ...

			g.Describe("map:map", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 3},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsTrue()
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 2},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(result).IsTrue()
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"eachMax", 1},
						},
					}

					result := filter.IsValid(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(result).IsFalse()
				})
			})
		})

	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidEachMatch ./pkg/validation/...

func TestIsValidEachMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eachMatch"`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Hash  [2]string
				Empty [0]string
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty list", func() {
				filter := Filter{
					{
						Field: "Empty",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Array{
					Empty: [0]string{},
				})

				g.Assert(result).IsTrue()
			})

			g.It("failure when at least 1 value does not match", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"Z0zZ0z19711zZz3z73z41z909z66zZzZ",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch"},
					},
				}

				result := filter.IsValid(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when given an empty mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", ``},
					},
				}

				result := filter.IsValid(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when given an empty rule", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{},
					},
				}

				result := filter.IsValid(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Hash []string
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty list", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{},
				})

				g.Assert(result).IsTrue()
			})

			g.It("failure when at least 1 value does not match", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"Z0zZ0z19711zZz3z73z41z909z66zZzZ",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch"},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when given an empty mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", ``},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when given an empty rule", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{},
					},
				}

				result := filter.IsValid(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Hash map[int]string
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty list", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{},
				})

				g.Assert(result).IsTrue()
			})

			g.It("failure when at least 1 value does not match", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "Z0zZ0z19711zZz3z73z41z909z66zZzZ",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch"},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when given an empty mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", ``},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})

			g.It("failure when given an empty rule", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{},
					},
				}

				result := filter.IsValid(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(result).IsFalse()
			})
		})

	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidMinFields ./pkg/validation/...

func TestIsValidMinFields(t *testing.T) {
	type Article struct {
		Id        uint16 `json:"id"`
		Status    uint8  `json:"status"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	g := Goblin(t)

	g.Describe(`Rule "minFields"`, func() {
		g.It("success when passing required fields", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
				{
					Field: "Status",
					Check: Range{1, 5},
				},
				{
					Field: "FirstName",
					Check: Range{3, 15},
				},
				{
					Field: "LastName",
					Check: Range{3, 15},
				},
				{
					Check: Rule{"minFields", 4},
				},
			}

			result := filter.IsValid(Article{
				Id:        10,
				Status:    5,
				FirstName: "John",
				LastName:  "Doe",
			})

			g.Assert(result).IsTrue()
		})

		g.It("failure when missing required fields", func() {
			filter := Filter{
				{
					Field: "Id",
					Check: NON_ZERO,
				},
				{
					Field: "Status",
					Check: Range{1, 5},
				},
				{
					Field: "FirstName",
					Check: Range{3, 15},
				},
				{
					Field: "LastName",
					Check: Range{3, 15},
				},
				{
					Check: Rule{"minFields", 4},
				},
			}

			result := filter.IsValid(Article{})
			g.Assert(result).IsFalse()
		})
	})
}

// go test ./pkg/validation/...
// go test -v -run TestIsValidNonZero ./pkg/validation/...

func TestIsValidNonZero(t *testing.T) {
	type Article struct {
		Int8       int8
		Int16      int16
		Int32      int32
		Int64      int64
		Int        int
		Uint8      uint8
		Uint16     uint16
		Uint32     uint32
		Uint64     uint64
		Uint       uint
		Float32    float32
		Float64    float64
		Complex64  complex64
		Complex128 complex128
		Bool       bool
		String     string
		Struct     time.Time
	}

	g := Goblin(t)

	g.Describe(`Rule NON_ZERO`, func() {
		g.It("success when given a non-zero int8 value", func() {
			filter := Filter{
				{
					Field: "Int8",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int8: math.MinInt8})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero int16 value", func() {
			filter := Filter{
				{
					Field: "Int16",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int16: math.MinInt16})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero int32 value", func() {
			filter := Filter{
				{
					Field: "Int32",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int32: math.MinInt32})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero int64 value", func() {
			filter := Filter{
				{
					Field: "Int64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int64: math.MinInt64})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero int value", func() {
			filter := Filter{
				{
					Field: "Int",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int: math.MinInt})
			g.Assert(result).IsTrue()
		})

		// ...

		g.It("success when given a non-zero uint8 value", func() {
			filter := Filter{
				{
					Field: "Uint8",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint8: 1})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero uint16 value", func() {
			filter := Filter{
				{
					Field: "Uint16",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint16: 1})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero uint32 value", func() {
			filter := Filter{
				{
					Field: "Uint32",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint32: 1})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero uint64 value", func() {
			filter := Filter{
				{
					Field: "Uint64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint64: 1})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero uint value", func() {
			filter := Filter{
				{
					Field: "Uint",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint: 1})
			g.Assert(result).IsTrue()
		})

		// ...

		g.It("success when given a non-zero float32 value", func() {
			filter := Filter{
				{
					Field: "Float32",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Float32: math.MaxFloat32})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero float64 value", func() {
			filter := Filter{
				{
					Field: "Float64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Float64: math.MaxFloat64})
			g.Assert(result).IsTrue()
		})

		// ...

		g.It("success when given a non-zero complex64 value", func() {
			filter := Filter{
				{
					Field: "Complex64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Complex64: -1})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero complex128 value", func() {
			filter := Filter{
				{
					Field: "Complex128",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Complex128: 1})
			g.Assert(result).IsTrue()
		})

		// ...

		g.It("success when given a non-zero boolean value", func() {
			filter := Filter{
				{
					Field: "Bool",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Bool: true})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero string value", func() {
			filter := Filter{
				{
					Field: "String",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{String: "ok"})
			g.Assert(result).IsTrue()
		})

		g.It("success when given a non-zero struct value", func() {
			filter := Filter{
				{
					Field: "Struct",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Struct: time.Now()})
			g.Assert(result).IsTrue()
		})

		// failure

		g.It("failure when given a zero int8 value", func() {
			filter := Filter{
				{
					Field: "Int8",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int8: *new(int8)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero int16 value", func() {
			filter := Filter{
				{
					Field: "Int16",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int16: *new(int16)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero int32 value", func() {
			filter := Filter{
				{
					Field: "Int32",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int32: *new(int32)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero int64 value", func() {
			filter := Filter{
				{
					Field: "Int64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int64: *new(int64)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero int value", func() {
			filter := Filter{
				{
					Field: "Int",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Int: *new(int)})
			g.Assert(result).IsFalse()
		})

		// ...

		g.It("failure when given a zero uint8 value", func() {
			filter := Filter{
				{
					Field: "Uint8",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint8: *new(uint8)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero uint16 value", func() {
			filter := Filter{
				{
					Field: "Uint16",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint16: *new(uint16)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero uint32 value", func() {
			filter := Filter{
				{
					Field: "Uint32",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint32: *new(uint32)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero uint64 value", func() {
			filter := Filter{
				{
					Field: "Uint64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint64: *new(uint64)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero uint value", func() {
			filter := Filter{
				{
					Field: "Uint",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Uint: *new(uint)})
			g.Assert(result).IsFalse()
		})

		// ...

		g.It("failure when given a zero float32 value", func() {
			filter := Filter{
				{
					Field: "Float32",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Float32: *new(float32)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero float64 value", func() {
			filter := Filter{
				{
					Field: "Float64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Float64: *new(float64)})
			g.Assert(result).IsFalse()
		})

		// ...

		g.It("failure when given a zero complex64 value", func() {
			filter := Filter{
				{
					Field: "Complex64",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Complex64: *new(complex64)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero complex128 value", func() {
			filter := Filter{
				{
					Field: "Complex128",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Complex128: *new(complex128)})
			g.Assert(result).IsFalse()
		})

		// ...

		g.It("failure when given a zero boolean value", func() {
			filter := Filter{
				{
					Field: "Bool",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Bool: *new(bool)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero string value", func() {
			filter := Filter{
				{
					Field: "String",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{String: *new(string)})
			g.Assert(result).IsFalse()
		})

		g.It("failure when given a zero struct value", func() {
			filter := Filter{
				{
					Field: "Struct",
					Check: NON_ZERO,
				},
			}

			result := filter.IsValid(Article{Struct: *new(time.Time)})
			g.Assert(result).IsFalse()
		})
	})
}
