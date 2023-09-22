package test

import (
	"encoding/json"
	"testing"
	"yu/golang/src/Url"
)

func convertToJson(data any) (string, error) {
	if bytes, err := json.Marshal(data); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func TestNewQueryMap(t *testing.T) {

	t.Run("query with empty value [1]", func(t *testing.T) {
		t.Parallel()
		const query = "data"
		const expect = `{"data":""}`

		if res, err := convertToJson(Url.NewQueryMap(query)); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})

	t.Run("query with empty value [2]", func(t *testing.T) {
		t.Parallel()
		const query = "data="
		const expect = `{"data":""}`

		if res, err := convertToJson(Url.NewQueryMap(query)); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})

	t.Run("rewrite existed value [1]", func(t *testing.T) {
		t.Parallel()
		const query = "data=1&data"
		const expect = `{"data":""}`

		if res, err := convertToJson(Url.NewQueryMap(query)); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})

	t.Run("rewrite existed value [2]", func(t *testing.T) {
		t.Parallel()
		const query = "data&data=1"
		const expect = `{"data":"1"}`

		if res, err := convertToJson(Url.NewQueryMap(query)); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})

	t.Run("parse object", func(t *testing.T) {
		t.Parallel()
		const query = "filter[fname]=Alex&filter[sname]=Yu&filter[bdate]=2020-01-01"
		const expect = `{"filter":{"bdate":"2020-01-01","fname":"Alex","sname":"Yu"}}`

		if res, err := convertToJson(Url.NewQueryMap(query)); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})

	t.Run("parse nested tree", func(t *testing.T) {
		t.Parallel()
		const query = "text=hello+world!&filter[slice][0][regionId]=2&filter[slice][0][phone]=380001234567&filter[slice][0][phone]=&filter[slice][0][data][min]=21&filter[slice][0][data][max]=41&filter[slice][0][data][conf][isActual]=true&filter[slice][0][data][conf][isUniq]=0&filter[slice][0][data][conf][min]=1&filter[slice][0][data][conf][max]=9&filter[slice][1][regionId]=4&filter[slice][1][phone]=380001234567&filter[slice][1][data][min]=4&filter[slice][1][data][max]=8&filter[slice][1][data][conf][isActual]=&filter[slice][1][data][conf][isUniq]=true&filter[text]=%D0%BF%D1%80%D0%B8%D0%B2%D1%96%D1%82!&filter[orderBy][0]=user_id+DESC&filter[orderBy][1]=created_at+DESC&filter[limit]=12&filter[offset]=6"
		const expect = `{"filter":{"limit":"12","offset":"6","orderBy":{"0":"user_id DESC","1":"created_at DESC"},"slice":{"0":{"data":{"conf":{"isActual":"true","isUniq":"0","max":"9","min":"1"},"max":"41","min":"21"},"phone":"380001234567","regionId":"2"},"1":{"data":{"conf":{"isActual":"","isUniq":"true"},"max":"8","min":"4"},"phone":"380001234567","regionId":"4"}},"text":"привіт!"},"text":"hello world!"}`

		if res, err := convertToJson(Url.NewQueryMap(query)); err != nil {
			t.Error(err)
		} else if res != expect {
			t.Errorf("Expect( %[1]v ) => Got( %[2]v )", expect, res)
		}
	})
}
