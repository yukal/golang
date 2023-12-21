package main

// To see the old examples, please visit:
// https://github.com/yukal/golang/tree/e8ebd290b0b8e59072ccbb5a509fcefeb43bc10a/examples/filter

import (
	"fmt"
	"time"
	"yu/golang/pkg/validation"
)

type Article struct {
	Id       uint64    `json:"id"`
	RegionId uint8     `json:"regionId"`
	Hash     string    `json:"hash"`
	Link     string    `json:"link"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
	Sex      uint8     `json:"sex"`
	Age      uint8     `json:"age"`
	Height   uint16    `json:"height"`
	Weight   uint16    `json:"weight"`
	Images   []string  `json:"images"`
	Phones   []string  `json:"phones"`
	Date     time.Time `json:"date"`
}

func main() {
	filter := validation.Filter{
		{
			Field: "Id",
			// means the initial value for each type, e.g. string(""), int(0)
			Check: validation.NON_ZERO,
		},
		{
			Field: "RegionId",
			Check: validation.Range{1, 25},
		},
		{
			Field: "Hash",
			Check: validation.Rule{"match", `(?i)^[0-9a-f]{32}$`},
		},
		{
			Field: "Link",
			Check: validation.Rule{"match", `^https://domain.com/data/[/\w-]*$`},
		},
		{
			Field: "Title",
			// the range for string type means the length in a range of n1..n2
			Check: validation.Range{5, 25},
		},
		{
			Field: "Message",
			Check: validation.Range{10, 400},
		},
		{
			Field:    "Sex",
			Check:    validation.Rule{"eq", 2},
			Optional: true,
		},
		{
			Field: "Age",
			Check: validation.Range{18, 45},
		},
		{
			Field:    "Height",
			Check:    validation.Range{150, 200},
			Optional: true,
		},
		{
			Field:    "Weight",
			Check:    validation.Range{45, 55},
			Optional: true,
		},
		{
			Field: "Images",
			Check: validation.Group{
				// at least 1 item in a collection
				validation.Rule{"min", 1},
				// each must match the regex mask
				validation.Rule{"matchEach", `(?i)[0-9a-f]{32}\.(?:png|jpe?g)$`},
			},
		},
		{
			Field: "Phones",
			Check: validation.Group{
				// at least 1 item
				validation.Rule{"min", 1},
				// each must contain 10 digits
				validation.Rule{"matchEach", `^\d{10}$`},
			},
		},
		{
			Field: "Date",
			// must contain exactly 2024 year
			Check: validation.Rule{"year", 2024},
		},
		{
			// This rule is useful in cases where the count of optional elements prevails
			// and you do not know exactly which will be filled (non-zero).
			// It should be at the last position of the validation rules
			Check: validation.Rule{"min-fields", 2},
		},
	}

	article := &Article{
		Id:       100,
		RegionId: 10,
		Hash:     "zJszhq8ck9qtZhh1IhyqCqrwxLUyjJsu",
		Link:     "https://domain.com/data/somebody-to-love",
		Title:    "Love is...",
		Message:  "Somebody to love",
		// Sex:      2,
		// Age:      28,
		// Height:   165,
		// Weight:   60,
		Images: []string{
			"https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
			"https://img.domain.com/4792592a98f8b9143de71d1db403d163.jpg",
			"https://img.domain.com/92f2b876b8ea94f711d2173539e73802.png",
		},
		Phones: []string{"0001234567"},
		Date:   time.Now(),
	}

	if !filter.IsValid(article) {
		fmt.Println("Not valid!")
	} else {
		fmt.Println("Valid!")
	}

	if hints := filter.Validate(article); len(hints) > 0 {
		for n, field := range hints {
			fmt.Printf("%d. %s\n", n+1, field)
		}
	}
}
