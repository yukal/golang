package test

import (
	"testing"
	"yu/golang/internal/app"
	"yu/golang/pkg/Url"
)

const query = "text=hello+world!&filter[slice][0][regionId]=2&filter[slice][0][phone]=380001234567&filter[slice][0][phone]=&filter[slice][0][data][min]=21&filter[slice][0][data][max]=41&filter[slice][0][data][conf][isActual]=true&filter[slice][0][data][conf][isUniq]=0&filter[slice][0][data][conf][min]=1&filter[slice][0][data][conf][max]=9&filter[slice][1][regionId]=4&filter[slice][1][phone]=380001234567&filter[slice][1][data][min]=4&filter[slice][1][data][max]=8&filter[slice][1][data][conf][isActual]=&filter[slice][1][data][conf][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6"

// go test -bench=^BenchmarkInspectData -benchtime=1000x -benchmem ./test/benchmark

func BenchmarkInspectData(b *testing.B) {
	var r string
	qmap := Url.NewQueryMapR(query)

	for n := 0; n < b.N; n++ {
		r = app.InspectData(qmap)
	}

	result = r
}
