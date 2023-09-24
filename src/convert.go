package src

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
