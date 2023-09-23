package Url

import (
	"net/url"
	"reflect"
	"strings"
)

func NewQueryMapL(urlQuery string) QueryMap {
	params := strings.Split(urlQuery, "&")

	if len(params) < 1 {
		return QueryMap{}
	}

	cache := make(QueryMap, len(params))
	data := make(QueryMap, len(params))

	var depth uint
	var steps uint
	var chunks []string
	var value any

	qUnescaped, _ := url.QueryUnescape(params[0])
	chunks, value = divideQueryParam(qUnescaped)
	currKey := strings.Join(chunks, "/")

	for {
		steps++

		if length := len(chunks); length > 1 {
			nextKey := strings.Join(chunks[:length-1], "/")

			// 1. The shortest way of recording parameters using an already built path that is in the cache

			if reflect.ValueOf(cache[nextKey]).Kind() == reflect.Map {
				depth += 2

				currKey := chunks[length-1]
				fullKey := nextKey + "/" + currKey

				cache[fullKey] = value
				cache[nextKey].(QueryMap)[currKey] = cache[fullKey]

				chunks = chunks[length-1:]

			} else {

				// 2. If the given path is not in the cache, then it would recursively build

				depth++

				cache[currKey] = value

				value = QueryMap{chunks[length-1]: value}
				chunks = chunks[:length-1]

			}

			currKey = nextKey

			// 3. A tree branch has been created and will be stored in the root of the data object

		} else {
			depth++

			if currKey == chunks[0] {
				cache[currKey] = value
				data[currKey] = value
			} else {
				cache[chunks[0]] = value
			}

			if params = params[1:]; len(params) > 0 {
				qUnescaped, _ := url.QueryUnescape(params[0])
				chunks, value = divideQueryParam(qUnescaped)
				currKey = strings.Join(chunks, "/")
				depth = 0
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
