package test

import (
	"errors"
	"fmt"
	"testing"
	"yu/golang/internal/app"
	"yu/golang/internal/app/Url"
	"yu/golang/test"
)

var result any

// go test ./test/unit/inspect...
// go test -run TestInspect ./test/unit/inspect

func TestInspectData(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
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
	})

	t.Run("map", func(t *testing.T) {
		expect, err := test.ReadFile("./data/query.yaml")
		if err != nil {
			t.Fatal(err)
		}

		wrongItems := 0
		const rawQuery = "filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6"

		for n := 0; n < 8; n++ {
			queryMap := Url.NewQueryMapR(rawQuery)
			yaml := app.InspectData(queryMap)

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
	})
}
