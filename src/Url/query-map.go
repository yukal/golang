package Url

import (
	"net/url"
	"reflect"
	"strings"
)

const MAX_STEPS = 128
const MAX_DEPTH = 16

type QueryMap map[string]any

var querySeparators = map[rune]bool{
	'[': true,
	']': true,
	'=': true,
}

func NewQueryMap(urlQuery string) QueryMap {
	params := strings.Split(urlQuery, "&")

	cache := make(QueryMap, len(params))
	data := make(QueryMap, len(params))

	for step, param := range params {
		qUnescaped, _ := url.QueryUnescape(param)
		chunks, value := divideQueryParam(qUnescaped)

		if length := len(chunks); length > 1 {

			buildTree(chunks, value, data, cache, 0)

		} else {

			cache[chunks[0]] = value
			data[chunks[0]] = value

		}

		if step > MAX_STEPS {
			return data
		}
	}

	return data
}

func buildTree(chunks []string, value any, data, cache QueryMap, depth int) {
	length := len(chunks)
	nextKey := strings.Join(chunks[:length-1], "/")

	if reflect.ValueOf(cache[nextKey]).Kind() == reflect.Map {
		currKey := chunks[length-1]
		fullKey := nextKey + "/" + currKey

		cache[fullKey] = value
		cache[nextKey].(QueryMap)[currKey] = cache[fullKey]

		return
	}

	currKey := strings.Join(chunks, "/")
	cache[currKey] = value

	if length > 1 {
		if depth <= MAX_DEPTH {
			buildTree(chunks[:length-1], QueryMap{chunks[length-1]: value}, data, cache, depth+1)
		}

		return
	}

	data[currKey] = cache[currKey]
}

func divideQueryParam(param string) ([]string, string) {
	value := ""

	if index := strings.Index(param, "="); index > -1 {
		value = param[index+1:]
		param = param[:index]
	}

	return strings.FieldsFunc(param, queryDivider), value
}

func queryDivider(r rune) bool {
	_, ok := querySeparators[r]
	return ok
}
