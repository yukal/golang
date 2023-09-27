package main

import (
	"fmt"
	"yu/golang/internal/app"
)

func main() {
	settings := app.MustLoadSettings("../../data/settings.json")

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

	fmt.Println(settings)
	fmt.Println()
}
