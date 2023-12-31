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

func NewQueryMapR(urlQuery string) QueryMap {
	params := strings.Split(urlQuery, "&")

	if length := len(params); length > 0 && length < MAX_STEPS {
		cache := make(QueryMap, len(params))
		data := make(QueryMap, len(params))

		buildTree(params, length, data, cache)

		return data
	}

	return QueryMap{}
}

// TODO:
//
// 1.
// To increase the performance of the algorithm twice, we have to build a tree parsing the
// path from the left side at the same time as parsing the path from the right side.
// The algorithm should end in the middle of the path, in that case, this will mean that
// we are halfway through building the full path (one branch for a tree).

func buildTree(params []string, paramsLen int, data, cache QueryMap) {
	qUnescaped, _ := url.QueryUnescape(params[0])
	chunks, value := divideQueryParam(qUnescaped)
	chunksLen := len(chunks)

	switch {
	case chunksLen > 1 && chunksLen < MAX_DEPTH:
		buildBranch(chunks, value, data, cache, chunksLen, "")

	case chunksLen == 1:
		cache[chunks[0]] = value
		data[chunks[0]] = value
	}

	if paramsLen > 1 {
		buildTree(params[1:], paramsLen-1, data, cache)
	}
}

func buildBranch(chunks []string, value any, data, cache QueryMap, length int, currKey string) {
	nextKey := strings.Join(chunks[:length-1], "/")

	// 1. The shortest way of recording parameters using an already built path that is in the cache

	if reflect.ValueOf(cache[nextKey]).Kind() == reflect.Map {
		currKey := chunks[length-1]
		fullKey := nextKey + "/" + currKey

		cache[fullKey] = value
		cache[nextKey].(QueryMap)[currKey] = cache[fullKey]

		return
	}

	// 2. If the given path is not in the cache, then it would recursively build

	cache[currKey] = value

	if length > 1 {
		buildBranch(chunks[:length-1], QueryMap{chunks[length-1]: value}, data, cache, length-1, nextKey)
		return
	}

	// 3. A tree branch has been created and will be stored in the root of the data object

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
