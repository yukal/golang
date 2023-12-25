package test

import (
	"testing"
	"yu/golang/pkg/Url"
	"yu/golang/test"

	. "github.com/franela/goblin"
)

// go test ./test/unit/url...
// go test -v -run TestNewQueryMapR ./test/unit/url
// go test -v -run TestNewQueryMapL ./test/unit/url

func TestNewQueryMapR(t *testing.T) {
	g := Goblin(t)

	g.Describe("NewQueryMapR", func() {
		g.Describe("Parsing a non-empty query", func() {
			g.It("overwrite a duplicate key with an empty value", func() {
				const expect = `{"data":""}`
				const query = "data=1&data"

				res, err := test.ConvertToJson(Url.NewQueryMapR(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It("overwrite a duplicate key with a filled value", func() {
				const expect = `{"data":"1"}`
				const query = "data&data=1"

				res, err := test.ConvertToJson(Url.NewQueryMapR(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It("parsing a simple object", func() {
				const expect = `{"filter":{"bdate":"2020-01-01","fname":"Alex","sname":"Yu"}}`
				const query = "filter[fname]=Alex&filter[sname]=Yu&filter[bdate]=2020-01-01"

				res, err := test.ConvertToJson(Url.NewQueryMapR(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It("parsing a nested tree", func() {
				const expect = `{"filter":{"limit":"12","offset":"6","orderBy":{"0":"user_id DESC","1":"created_at DESC"},"slice":{"0":{"data":{"conf":{"isActual":"true","isUniq":"0","max":"9","min":"1"},"max":"41","min":"21"},"phone":"","regionId":"2"},"1":{"data":{"conf":{"isActual":"","isUniq":"true"},"max":"8","min":"4"},"phone":"380001234567","regionId":"4"}},"text":"привіт!"},"text":"hello world!"}`
				const query = "text=hello+world!&filter[slice][0][regionId]=2&filter[slice][0][phone]=380001234567&filter[slice][0][phone]=&filter[slice][0][data][min]=21&filter[slice][0][data][max]=41&filter[slice][0][data][conf][isActual]=true&filter[slice][0][data][conf][isUniq]=0&filter[slice][0][data][conf][min]=1&filter[slice][0][data][conf][max]=9&filter[slice][1][regionId]=4&filter[slice][1][phone]=380001234567&filter[slice][1][data][min]=4&filter[slice][1][data][max]=8&filter[slice][1][data][conf][isActual]=&filter[slice][1][data][conf][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6"

				res, err := test.ConvertToJson(Url.NewQueryMapR(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

		})

		g.Describe("Parsing an empty query", func() {
			g.It(`parsing without an equal sign "="`, func() {
				const expect = `{"data":""}`
				const query = "data"

				res, err := test.ConvertToJson(Url.NewQueryMapR(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It(`parsing with an equal sign "="`, func() {
				const expect = `{"data":""}`
				const query = "data="

				res, err := test.ConvertToJson(Url.NewQueryMapR(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})
		})
	})
}

func TestNewQueryMapL(t *testing.T) {
	g := Goblin(t)

	g.Describe("NewQueryMapL", func() {
		g.Describe("Parsing a non-empty query", func() {
			g.It("overwrite a duplicate key with an empty value", func() {
				const expect = `{"data":""}`
				const query = "data=1&data"

				res, err := test.ConvertToJson(Url.NewQueryMapL(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It("overwrite a duplicate key with a filled value", func() {
				const expect = `{"data":"1"}`
				const query = "data&data=1"

				res, err := test.ConvertToJson(Url.NewQueryMapL(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It("parsing a simple object", func() {
				const expect = `{"filter":{"bdate":"2020-01-01","fname":"Alex","sname":"Yu"}}`
				const query = "filter[fname]=Alex&filter[sname]=Yu&filter[bdate]=2020-01-01"

				res, err := test.ConvertToJson(Url.NewQueryMapL(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It("parsing a nested tree", func() {
				const expect = `{"filter":{"limit":"12","offset":"6","orderBy":{"0":"user_id DESC","1":"created_at DESC"},"slice":{"0":{"data":{"conf":{"isActual":"true","isUniq":"0","max":"9","min":"1"},"max":"41","min":"21"},"phone":"","regionId":"2"},"1":{"data":{"conf":{"isActual":"","isUniq":"true"},"max":"8","min":"4"},"phone":"380001234567","regionId":"4"}},"text":"привіт!"},"text":"hello world!"}`
				const query = "text=hello+world!&filter[slice][0][regionId]=2&filter[slice][0][phone]=380001234567&filter[slice][0][phone]=&filter[slice][0][data][min]=21&filter[slice][0][data][max]=41&filter[slice][0][data][conf][isActual]=true&filter[slice][0][data][conf][isUniq]=0&filter[slice][0][data][conf][min]=1&filter[slice][0][data][conf][max]=9&filter[slice][1][regionId]=4&filter[slice][1][phone]=380001234567&filter[slice][1][data][min]=4&filter[slice][1][data][max]=8&filter[slice][1][data][conf][isActual]=&filter[slice][1][data][conf][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6"

				res, err := test.ConvertToJson(Url.NewQueryMapL(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

		})

		g.Describe("Parsing an empty query", func() {
			g.It(`parsing without an equal sign "="`, func() {
				const expect = `{"data":""}`
				const query = "data"

				res, err := test.ConvertToJson(Url.NewQueryMapL(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})

			g.It(`parsing with an equal sign "="`, func() {
				const expect = `{"data":""}`
				const query = "data="

				res, err := test.ConvertToJson(Url.NewQueryMapL(query))

				g.Assert(err).IsNil(err)
				g.Assert(res).Equal(expect)
			})
		})
	})
}
