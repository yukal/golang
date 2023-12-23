package app

import (
	"encoding/json"
	"fmt"
)

func MustConvertToJson(data any) string {
	// if bytes, err := json.Marshal(data); err != nil {
	if bytes, err := json.MarshalIndent(data, "", "  "); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", bytes)
	}
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func SliceContains[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item == value {
			return true
		}
	}

	return false
}

func SliceDifference[T comparable](sliceA, sliceB []T) []T {
	if len(sliceA) == 0 && len(sliceB) == 0 {
		return []T{}
	}

	biggest := sliceA
	smallest := sliceB

	if len(sliceB) > len(sliceA) {
		biggest = sliceB
		smallest = sliceA
	}

	diff := make([]T, 0, len(biggest))
	smallestMap := make(map[T]any, len(smallest))

	for _, str := range smallest {
		smallestMap[str] = nil
	}

	for _, str := range biggest {
		if _, exist := smallestMap[str]; !exist {
			diff = append(diff, str)
		}
	}

	return diff
}
