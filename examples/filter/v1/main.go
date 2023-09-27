package main

import (
	"fmt"
	"time"
	"yu/golang/internal/app"
	"yu/golang/internal/app/validation"
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
	filter := app.ArticleFilter{
		Id:       map[string]uint64{"min": 1},
		RegionId: map[string]uint8{"min": 1, "max": 25},
		Hash:     map[string]uint8{"minLen": 1},
		Link:     map[string]uint8{"minLen": 1},
		Title:    map[string]uint8{"minLen": 5},
		Message:  map[string]uint8{"minLen": 10},
		Sex:      map[string]uint8{"eq": 2},
		Age:      map[string]uint8{"min": 18, "max": 45},
		Height:   map[string]uint8{"min": 150, "max": 200},
		Weight:   map[string]uint8{"min": 45, "max": 80},
		Images:   map[string]uint8{"minLen": 1},
		Phones:   map[string]uint8{"minLen": 1},
		Date:     map[string]uint16{"year": 2023},
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

	if !validation.IsValid(filter, article) {
		// panic(errors.New("not valid"))
		fmt.Println("Not valid!")
	} else {
		fmt.Println("Valid!")
	}

	if errs := validation.Validate(filter, article); len(errs) > 0 {
		for n, field := range errs {
			fmt.Printf("%d. Wrong %s\n", n+1, field)
		}
	}
}
