package test

import (
	"errors"
	"fmt"
	"testing"
	"yu/golang/internal/app"
	"yu/golang/test"
)

var result any

func TestInspectData(t *testing.T) {
	expect, err := test.ReadFile("./data/settings.yaml")
	if err != nil {
		t.Fatal(err)
	}

	wrongItems := 0

	for n := 0; n < 8; n++ {
		settings := app.MustLoadSettings("./data/settings.json")
		yaml := app.InspectData(settings)

		if yaml != expect {
			wrongItems++
			// t.Fatalf("Expect(\n%s)\n\nGot(\n%s)", expect, yaml)

			// if err := os.WriteFile("./1-expect.yaml", []byte(expect), 0644); err != nil {
			// 	t.Fatal(err)
			// }

			// if err := os.WriteFile("./2-got.yaml", []byte(yaml), 0644); err != nil {
			// 	t.Fatal(err)
			// }
		}
	}

	if wrongItems > 0 {
		msg := fmt.Sprintf("Generated settings are not match %d times", wrongItems)
		t.Error(errors.New(msg))
	}
}
