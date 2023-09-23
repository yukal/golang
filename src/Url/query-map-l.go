package Url

import (
	"net/url"
	"reflect"
	"strings"
)

func NewQueryMapL(urlQuery string) QueryMap {
	params := strings.Split(urlQuery, "&")
	paramsLen := len(params)

	if paramsLen < 1 {
		return QueryMap{}
	}

	cache := make(QueryMap, paramsLen)
	data := make(QueryMap, paramsLen)

	var steps uint
	var chunks []string
	var value any

	qUnescaped, _ := url.QueryUnescape(params[0])
	chunks, value = divideQueryParam(qUnescaped)
	currKey := strings.Join(chunks, "/")
	depth := len(chunks)
	params = params[1:]

	for {
		steps++

		if depth > 1 {
			nextKey := strings.Join(chunks[:depth-1], "/")

			// 1. The shortest way of recording parameters using an already built path that is in the cache

			if reflect.ValueOf(cache[nextKey]).Kind() == reflect.Map {

				currKey := chunks[depth-1]
				fullKey := nextKey + "/" + currKey

				cache[fullKey] = value
				cache[nextKey].(QueryMap)[currKey] = cache[fullKey]

				depth = 1 // should assign the value at the next iteration
				chunks = chunks[:depth:depth]

			} else {

				// 2. If the given path is not in the cache, then it would recursively build

				cache[currKey] = value
				value = QueryMap{chunks[depth-1]: value}

				depth--
				chunks = chunks[:depth:depth]

				currKey = nextKey
			}

			// 3. A tree branch has been created and will be stored in the root of the data object

		} else {

			if currKey == chunks[0] {
				cache[chunks[0]] = value
				data[chunks[0]] = value
			} else {
				cache[currKey] = value
			}

			if len(params) > 0 {
				qUnescaped, _ := url.QueryUnescape(params[0])
				chunks, value = divideQueryParam(qUnescaped)
				currKey = strings.Join(chunks, "/")
				depth = len(chunks)
				params = params[1:]
			} else {
				return data
			}

		}

		if steps > MAX_STEPS || depth > MAX_DEPTH {
			break
		}
	}

	return data
}
