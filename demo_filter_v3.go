package main

import (
	"fmt"
	"os"
	"path"
	"time"
	"yu/golang/src"
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
	appPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	settings := src.MustLoadSettings(
		path.Join(appPath, "data/settings.json"))

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

	if !validation.IsValid(settings.Task.Filter, article) {
		fmt.Println("Not valid!")
	} else {
		fmt.Println("Valid!")
	}

	if errs := validation.Validate(settings.Task.Filter, article); len(errs) > 0 {
		for n, field := range errs {
			fmt.Printf("%d. Wrong %s\n", n+1, field)
		}
	}
}
