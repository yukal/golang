package main

import (
	"fmt"
	"yu/golang/src"
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
	Images   uint8
	Phones   uint8
	Date     uint16
}

func DemoFilter() {
	filter := &src.Filter{
		Id:       map[string]uint64{"min": 1},
		RegionId: map[string]uint8{"min": 1, "max": 25},
		// Hash:     map[string]uint8{"minLen": 1},
		// Link:     map[string]uint8{"minLen": 1},
		// Title:    map[string]uint8{"minLen": 5},
		// Message:  map[string]uint8{"minLen": 10},
		// Sex:      map[string]uint8{"eq": 2},
		Age:      map[string]uint8{"min": 18, "max": 45},
		Height:   map[string]uint8{"min": 150, "max": 200},
		Weight:   map[string]uint8{"min": 45, "max": 80},
		// Images:   map[string]uint8{"minLen": 1},
		// Phones:   map[string]uint8{"minLen": 1},
		// Date:     map[string]uint16{"year": 2023},
	}

	article := &Article{
		Id:       0,
		RegionId: 0,
		Hash:     "",
		Link:     "",
		Title:    "",
		Message:  "",
		Sex:      0,
		Age:      17,
		Height:   145,
		Weight:   35,
		Images:   0,
		Phones:   0,
		Date:     0,
	}

	if !src.IsValid(filter, article) {
		// panic(errors.New("not valid"))
		fmt.Println("Not valid!")
	}

	if errs := src.Validate(filter, article); len(errs) > 0 {
		for n, field := range errs {
			fmt.Printf("%d. Wrong %s\n", n+1, field)
		}
	}
}
