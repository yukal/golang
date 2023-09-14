package Url

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
	"yu/golang/src"
)

// https://github.com/gin-gonic/gin/issues/1310

func NewSearchParams(rawUrlQuery string, data any) {
	qmap := NewQueryMap(rawUrlQuery)

	ref := reflect.Indirect(reflect.ValueOf(data))
	InspectQuery(ref, qmap, 0)
}

func InspectQuery(val reflect.Value, qmap QueryMap, depth int) {
	switch val.Kind() {
	case reflect.Struct:
		urlQueryInspectStruct(val, qmap, depth)

	case reflect.Map:
		urlQueryInspectMap(val, qmap, depth)

	case reflect.Slice:
		urlQueryInspectSlice(val, qmap, depth)
	}
}

func urlQueryInspectStruct(mval reflect.Value, qmap QueryMap, depth int) {
	for n := 0; n < mval.Type().NumField(); n++ {
		var (
			ok    bool
			qm    QueryMap = QueryMap{}
			value any
		)

		typ := mval.Type().Field(n)
		mval := mval.Field(n)

		keyname := typ.Tag.Get("query")
		if value, ok = qmap[keyname]; ok {
			if reflect.ValueOf(value).Kind() == reflect.Map {
				qm = value.(QueryMap)
			}
		}

		if !mval.IsValid() {
			return
		}

		if src.IsPrimitive(mval) {

			urlQuerySetValue(mval, value)

		} else {

			// Unmarshal complex data
			InspectQuery(mval, qm, depth+1)

		}
	}
}

func urlQueryInspectMap(mval reflect.Value, qmap QueryMap, depth int) {
	for _, key := range mval.MapKeys() {
		value := mval.MapIndex(key)

		if src.IsPrimitive(value.Interface()) {

			urlQuerySetValue(value, value)

		} else {

			// TODO: Unmarshal complex data

		}
	}
}

func urlQueryInspectSlice(mval reflect.Value, qmap QueryMap, depth int) {
	keys := sortMapKeys(qmap)

	for _, key := range keys {
		value := qmap[key]

		if src.IsPrimitive(value) {

			newSlice := reflect.Append(mval, reflect.ValueOf(value))
			mval.Set(newSlice)

		} else {

			// TODO: Unmarshal complex data

		}
	}
}

func urlQuerySetValue(mval reflect.Value, value any) (any, error) {
	if !(mval.CanSet() && reflect.ValueOf(value).Kind() == reflect.String) {
		return nil, errors.New("cannot set | wrong input param type")
	}

	switch mval.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		if v, err := strconv.ParseInt(value.(string), 10, mval.Type().Bits()); err != nil {
			return nil, err
		} else {
			mval.SetInt(v)
		}

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		if v, err := strconv.ParseUint(value.(string), 10, mval.Type().Bits()); err != nil {
			return nil, err
		} else {
			mval.SetUint(v)
		}

	case reflect.Float32, reflect.Float64:
		if v, err := strconv.ParseFloat(value.(string), mval.Type().Bits()); err != nil {
			return nil, err
		} else {
			mval.SetFloat(v)
		}

	case reflect.Complex64, reflect.Complex128:
		if v, err := strconv.ParseComplex(value.(string), mval.Type().Bits()); err != nil {
			return nil, err
		} else {
			mval.SetComplex(v)
		}

	case reflect.Bool:
		switch value.(string) {
		case "true", "1":
			mval.SetBool(true)

		case "false", "0":
			mval.SetBool(false)
		}

	case reflect.String:
		mval.SetString(value.(string))

	case reflect.Slice:
		for _, vl := range value.(QueryMap) {
			newSlice := reflect.Append(mval, reflect.ValueOf(vl))
			mval.Set(newSlice)
		}
	}

	return nil, nil
}

func sortMapKeys(qmap QueryMap) []string {
	keys := make([]string, 0, len(qmap))

	for k := range qmap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func sortMap(qmap QueryMap) QueryMap {
	keys := make([]string, 0, len(qmap))
	data := make(QueryMap, len(qmap))

	for k := range qmap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		data[k] = qmap[k]
	}

	return data
}
