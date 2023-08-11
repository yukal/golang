package main

import (
	"fmt"
	"yu/golang/src"
)

func main() {
	settings := src.MustLoadSettings("data/settings.json")
	fmt.Println(src.InspectData(settings))
}
