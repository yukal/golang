package main

import (
	"fmt"
	"os"
	"path"
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

func (article *Article) IsValid(filter any) bool {
	flag, _ := validation.CheckIsValid(filter, article, true)
	return flag
}

func (article *Article) Validate(filter any) []string {
	_, wrongFields := validation.CheckIsValid(filter, article, false)
	return wrongFields
}

func main() {
	appPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	settings := app.MustLoadSettings(
		path.Join(appPath, "../../../data/settings.json"))

	article := &Article{
		Id:       100,
		RegionId: 10,
		Hash:     "b0fb0c19711bcf3b73f41c909f66bded",
		Link:     "https://domain.com/data/",
		Title:    "Title",
		Message:  "Hello World!",
		Sex:      2,
		Age:      28,
		Height:   165,
		Weight:   60,
		Images: []string{
			"https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
			"https://img.domain.com/4792592a98f8b9143de71d1db403d163.jpg",
			"https://img.domain.com/92f2b876b8ea94f711d2173539e73802.png",
		},
		Phones: []string{"380001234567"},
		Date:   time.Now(),
	}

	if !article.IsValid(settings.Task.Filter) {
		fmt.Println("Not valid!")
	} else {
		fmt.Println("Valid!")
	}

	if errs := article.Validate(settings.Task.Filter); len(errs) > 0 {
		for n, field := range errs {
			fmt.Printf("%d. Wrong %s\n", n+1, field)
		}
	}
}
