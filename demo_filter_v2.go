package main

import (
	"fmt"
	"time"
	"yu/golang/src/validation"
)

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
	Images   []string
	Phones   []string
	Date     time.Time
}

func main() {
	filter := validation.FilterMap{
		"Id":       {"min": uint64(1)},
		"RegionId": {"min": uint8(1), "max": uint8(25)},
		"Hash":     {"minLen": uint8(1)},
		"Link":     {"minLen": uint8(1)},
		"Title":    {"minLen": uint8(5)},
		"Message":  {"minLen": uint8(10)},
		"Sex":      {"eq": uint8(2)},
		"Age":      {"min": uint8(18), "max": uint8(45)},
		"Height":   {"min": uint8(150), "max": uint8(200)},
		"Weight":   {"min": uint8(45), "max": uint8(80)},
		"Images":   {"minLen": uint8(1)},
		"Phones":   {"minLen": uint8(1)},
		"Date":     {"year": uint16(2023)},
	}

	article := &Article{
		Id:       100,
		RegionId: 10,
		Hash:     "zJszhq8ck9qtZhh1IhyqCqrwxLUyjJsu",
		Link:     "https://domain.com/data/",
		Title:    "Title",
		Message:  "Hello World!",
		Sex:      2,
		Age:      28,
		Height:   165,
		Weight:   60,
		Images:   []string{"img1", "img2"},
		Phones:   []string{"0001234567"},
		Date:     time.Now(),
	}

	if !filter.IsValid(article) {
		// panic(errors.New("not valid"))
		fmt.Println("Not valid!")
	} else {
		fmt.Println("Valid!")
	}

	if errs := filter.Validate(article); len(errs) > 0 {
		for n, field := range errs {
			fmt.Printf("%d. Wrong %s\n", n+1, field)
		}
	}
}
