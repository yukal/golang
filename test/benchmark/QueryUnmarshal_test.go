package test

import (
	"net/url"
	"testing"
	"yu/golang/pkg/Url"
)

// go test -bench=^BenchmarkQuery -benchtime=1000x -benchmem ./test/benchmark

func BenchmarkQueryUnmarshal(b *testing.B) {
	const query = `filter[data][0][regionId]=1&filter[data][0][phone]=380001234567&filter[data][0][age][min]=18&filter[data][0][age][max]=28&filter[data][0][nes][nes2][isActual]=true&filter[data][0][nes][nes2][isUniq]=0&filter[data][1][regionId]=2&filter[data][1][phone]=380007654321&filter[data][1][age][min]=21&filter[data][1][age][max]=41&filter[data][1][nes][nes2][isActual]=&filter[data][1][nes][nes2][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6`
	var payload Payload

	u, err := url.ParseQuery(query)
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		Url.UnmarshalQuery(u, &payload)
	}
}
