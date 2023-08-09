package main

import (
	"fmt"
	"strings"
	"yu/golang/src"
)

func DemoSettings() {
	settings := src.MustLoadSettings("data/settings.json")
	separator := strings.Repeat("=", 12)

	// settings.Paths.Database
	// settings.Paths.Images

	// settings.Server.Host
	// settings.Server.Protocol
	// settings.Server.Urls

	// settings.Task.Filter.Id
	// settings.Task.Filter.RegionId
	// settings.Task.Filter.Sex
	// settings.Task.Filter.Age
	// settings.Task.Filter.Height
	// settings.Task.Filter.Weight
	// settings.Task.Filter.Title
	// settings.Task.Filter.Message
	// settings.Task.Filter.Images
	// settings.Task.Filter.Phones
	// settings.Task.Filter.Link
	// settings.Task.Filter.Hash
	// settings.Task.Filter.Date

	fmt.Printf("%s\n%s\n%[1]s\n\n", separator, settings.String())
}
