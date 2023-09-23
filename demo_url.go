package main

import (
	"fmt"
	"net/url"
	"yu/golang/src"
	"yu/golang/src/Url"
)

func main() {
	type FilterNested2 struct {
		Val string `query:"val"`
	}

	type FilterNested1 struct {
		Nested2 FilterNested2 `query:"nes2"`
	}

	type FilterMinMax struct {
		Min uint8 `query:"min"`
		Max uint8 `query:"max"`
	}

	type Filter struct {
		RegionId uint8         `query:"regionId"`
		Age      FilterMinMax  `query:"age"`
		Text     string        `query:"text"`
		Phone    string        `query:"phone"`
		IsActual bool          `query:"isActual"`
		IsUniq   bool          `query:"isUniq"`
		OrderBy  []string      `query:"orderBy"`
		Offset   int64         `query:"offset"`
		Limit    int64         `query:"limit"`
		Nested1  FilterNested1 `query:"nes1"`
	}

	type Payload struct {
		Filter Filter `json:"filter" query:"filter"`
	}

	var payload Payload

	rawURL := "http://localhost:50598/news/?filter[regionId]=2&filter[age][min]=21&filter[age][max]=41&filter[phone]=380001234567&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[isActual]=true&filter[isUniq]=0&filter[nes1][nes2][val]=tadam!&filter[orderBy][0]=created_at+DESC&filter[orderBy][1]=user_id+DESC&filter[limit]=16&filter[offset]=0"
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	queryMap := Url.NewQueryMapR(u.RawQuery)
	fmt.Printf("%#v\n\n", queryMap)

	// Unmarshal raw query to a struct
	Url.NewSearchParams(u.RawQuery, &payload)

	// fmt.Printf("%#v\n", payload)
	fmt.Println(src.InspectData(payload))
}
