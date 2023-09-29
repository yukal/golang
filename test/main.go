package test

import (
	"encoding/json"
	"os"
)

func ConvertToJson(data any) (string, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ConvertToJsonI(data any) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ReadFile(path string) (string, error) {
	if b, err := os.ReadFile(path); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
