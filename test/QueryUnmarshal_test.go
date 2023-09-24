package test

import (
	"net/url"
	"testing"
	"yu/golang/src/Url"
)

// go test -run TestQueryUnmarshal
// go test -bench=. -benchtime=1000x -benchmem

func TestQueryUnmarshal(t *testing.T) {
	t.Run("query with empty value", func(t *testing.T) {
		type Payload struct {
			Data string `json:"data" query:"data"`
		}

		t.Run(`without sign "="`, func(t *testing.T) {
			t.Parallel()

			const query = "data"
			const expect = `{"data":""}`

			var payload Payload

			u, err := url.ParseQuery(query)
			if err != nil {
				t.Fatal(err)
			}

			Url.UnmarshalQuery(u, &payload)

			if res, err := convertToJson(payload); err != nil {
				t.Error(err)
			} else if res != expect {
				t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
			}
		})

		t.Run(`with sign "="`, func(t *testing.T) {
			t.Parallel()

			const query = "data="
			const expect = `{"data":""}`

			var payload Payload

			u, err := url.ParseQuery(query)
			if err != nil {
				t.Fatal(err)
			}

			Url.UnmarshalQuery(u, &payload)

			if res, err := convertToJson(payload); err != nil {
				t.Error(err)
			} else if res != expect {
				t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
			}
		})

	})

	t.Run("parse nested tree", func(t *testing.T) {
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
			Filter Filter `json:"filter" query:"filter"`
		}

		const query = `filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6`
		const expect = `{"filter":{"data":[{"regionId":1,"phone":"380001234567","age":{"min":18,"max":28},"nes":{"nes2":{"isActual":true,"isUniq":false}}},{"regionId":2,"phone":"380007654321","age":{"min":21,"max":41},"nes":{"nes2":{"isActual":false,"isUniq":true}}}],"text":"привіт!","orderBy":["user_id DESC","created_at DESC"],"limit":12,"offset":6}}`

		var payload Payload

		u, err := url.ParseQuery(query)
		if err != nil {
			t.Fatal(err)
		}

		Url.UnmarshalQuery(u, &payload)

		if res, err := convertToJson(payload); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})
}

func BenchmarkQueryUnmarshal(b *testing.B) {
	type FilterNested2 struct {
		IsActual bool `query:"isActual"`
		IsUniq   bool `query:"isUniq"`
	}

	type FilterNested1 struct {
		Nested FilterNested2 `query:"nes2"`
	}

	type FilterMinMax struct {
		Min uint8 `query:"min"`
		Max uint8 `query:"max"`
	}

	type FilterData struct {
		RegionId uint8         `query:"regionId"`
		Phone    string        `query:"phone"`
		Age      FilterMinMax  `query:"age"`
		Nested   FilterNested1 `query:"nes"`
	}

	type Filter struct {
		Data    []FilterData `query:"data"`
		Text    string       `query:"text"`
		OrderBy []string     `query:"orderBy"`
		Limit   int64        `query:"limit"`
		Offset  int64        `query:"offset"`
	}

	type Payload struct {
		Filter Filter `query:"filter"`
	}

	var payload Payload
	const query = `filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6`

	u, err := url.ParseQuery(query)
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		Url.UnmarshalQuery(u, &payload)
	}
}
