package test

import (
	"net/url"
	"testing"
	"yu/golang/internal/app/Url"
	"yu/golang/test"
)

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

			if res, err := test.ConvertToJson(payload); err != nil {
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

			if res, err := test.ConvertToJson(payload); err != nil {
				t.Error(err)
			} else if res != expect {
				t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
			}
		})

	})

	t.Run("parse nested tree", func(t *testing.T) {
		t.Parallel()

		const query = `filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6`
		const expect = `{"filter":{"data":[{"regionId":1,"phone":"380001234567","age":{"min":18,"max":28},"nes":{"nes2":{"isActual":true,"isUniq":false}}},{"regionId":2,"phone":"380007654321","age":{"min":21,"max":41},"nes":{"nes2":{"isActual":false,"isUniq":true}}}],"text":"привіт!","orderBy":["user_id DESC","created_at DESC"],"limit":12,"offset":6}}`

		var payload Payload

		u, err := url.ParseQuery(query)
		if err != nil {
			t.Fatal(err)
		}

		Url.UnmarshalQuery(u, &payload)

		if res, err := test.ConvertToJson(payload); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})
}
