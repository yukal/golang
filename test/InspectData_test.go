package test

import (
	"net/url"
	"os"
	"testing"
	"yu/golang/internal/app"
	"yu/golang/internal/app/Url"
)

// go test -run TestInspectData
// go test -bench=. -benchtime=1000x -benchmem

func TestInspectData(t *testing.T) {
	expectInspectData, err := os.ReadFile("./InspectData_res.yml")
	if err != nil {
		t.Fatal(err)
	}

	t.Parallel()

	type FilterNested2 struct {
		IsActual bool `query:"isActual" json:"isActual"`
		IsUniq   bool `query:"isUniq" json:"isUniq"`
	}

	type FilterNested1 struct {
		Nested FilterNested2 `query:"nes2" json:"nes2"`
	}

	type FilterMinMax struct {
		Min uint8 `query:"min" json:"min"`
		Max uint8 `query:"max" json:"max"`
	}

	type FilterData struct {
		RegionId uint8         `query:"regionId" json:"regionId"`
		Phone    string        `query:"phone" json:"phone"`
		Age      FilterMinMax  `query:"age" json:"age"`
		Nested   FilterNested1 `query:"nes" json:"nes"`
	}

	type Filter struct {
		Data    []FilterData `query:"data" json:"data"`
		Text    string       `query:"text" json:"text"`
		OrderBy []string     `query:"orderBy" json:"orderBy"`
		Limit   int64        `query:"limit" json:"limit"`
		Offset  int64        `query:"offset" json:"offset"`
	}

	type Payload struct {
		Filter Filter `query:"filter" json:"filter"`
	}

	var payload Payload
	var expect = string(expectInspectData)

	const query = `filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6`

	u, err := url.ParseQuery(query)
	if err != nil {
		t.Fatal(err)
	}

	Url.UnmarshalQuery(u, &payload)

	if res := app.InspectData(payload); res != expect {
		t.Errorf("Expect(\n%[1]v\n)\nGot(\n%[2]v\n)", expect, res)
	}
}

func BenchmarkInspectData(b *testing.B) {
	type FilterNested2 struct {
		IsActual bool `query:"isActual" json:"isActual"`
		IsUniq   bool `query:"isUniq" json:"isUniq"`
	}

	type FilterNested1 struct {
		Nested FilterNested2 `query:"nes2" json:"nes2"`
	}

	type FilterMinMax struct {
		Min uint8 `query:"min" json:"min"`
		Max uint8 `query:"max" json:"max"`
	}

	type FilterData struct {
		RegionId uint8         `query:"regionId" json:"regionId"`
		Phone    string        `query:"phone" json:"phone"`
		Age      FilterMinMax  `query:"age" json:"age"`
		Nested   FilterNested1 `query:"nes" json:"nes"`
	}

	type Filter struct {
		Data    []FilterData `query:"data" json:"data"`
		Text    string       `query:"text" json:"text"`
		OrderBy []string     `query:"orderBy" json:"orderBy"`
		Limit   int64        `query:"limit" json:"limit"`
		Offset  int64        `query:"offset" json:"offset"`
	}

	type Payload struct {
		Filter Filter `query:"filter" json:"filter"`
	}

	var payload Payload
	var res string
	const query = `filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6`

	u, err := url.ParseQuery(query)
	if err != nil {
		b.Fatal(err)
	}

	Url.UnmarshalQuery(u, &payload)

	for n := 0; n < b.N; n++ {
		res = app.InspectData(payload)
	}

	result = res
}
