package test

import (
	"encoding/json"
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
