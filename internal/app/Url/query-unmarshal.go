package Url

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"yu/golang/internal/app"
)

// https://github.com/gin-gonic/gin/issues/1310

func UnmarshalQuery(queryValues url.Values, data any) {
	ref := reflect.Indirect(reflect.ValueOf(data))
	fillTree(ref, queryValues, "")
}

func fillTree(val reflect.Value, queries url.Values, path string) {
	switch val.Kind() {
	case reflect.Struct:
		if fieldsCount := val.Type().NumField(); fieldsCount > 0 {
			fillStruct(val, queries, path, fieldsCount-1)
		}

	case reflect.Slice:
		itemType := getItemType(val.Interface())
		item := reflect.Indirect(reflect.New(itemType))

		fillSlice(val, item, queries, path, 0)
	}
}

func fillStruct(mval reflect.Value, queries url.Values, path string, index int) {
	fieldVal := mval.Field(index)
	fieldType := mval.Type().Field(index)
	ckey := getQueryPath(path, fieldType.Tag.Get("query"))

	if app.IsPrimitive(fieldVal.Interface()) {
		if value, ok := queries[ckey]; ok {

			SetValue(fieldVal, value[0])

		}
	} else {

		fillTree(fieldVal, queries, ckey)

	}

	if index > 0 {
		fillStruct(mval, queries, path, index-1)
	}
}

func fillSlice(val reflect.Value, item reflect.Value, queries url.Values, path string, index int) {
	ckey := getQueryPath(path, strconv.Itoa(index))

	if app.IsPrimitive(item) {
		if value, ok := queries[ckey]; ok {

			SetValue(val, value[0])
			fillSlice(val, item, queries, path, index+1)

		}
	} else {

		proto := fmt.Sprint(item)
		fillTree(item, queries, ckey)

		if proto != fmt.Sprint(item) {
			val.Set(reflect.Append(val, item))
			fillSlice(val, reflect.New(item.Type()).Elem(), queries, path, index+1)
		}

	}
}

func SetValue(mval reflect.Value, value any) (any, error) {
	if !mval.CanSet() {
		return nil, errors.New("cannot set variable")
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
		case "1", "true":
			mval.SetBool(true)

		case "0", "false", "":
			mval.SetBool(false)
		}

	case reflect.String:
		mval.SetString(value.(string))

	case reflect.Slice:
		mval.Set(reflect.Append(mval, reflect.ValueOf(value)))

	}

	return nil, nil
}

func getItemType(a interface{}) reflect.Type {
	for t := reflect.TypeOf(a); ; {
		switch t.Kind() {
		case reflect.Ptr, reflect.Slice:
			t = t.Elem()
		default:
			return t
		}
	}
}

func getQueryPath(queryPath, queryChunk string) string {
	if queryPath == "" {
		return queryChunk
	}

	return queryPath + "[" + queryChunk + "]"
	// return queryPath + "/" + queryChunk
}
