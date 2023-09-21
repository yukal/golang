package Url

import (
	"net/url"
	"reflect"
	"regexp"
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
	depth := 0

	for step, param := range params {
		qUnescaped, _ := url.QueryUnescape(param)
		chunks, value := divideQueryParam(qUnescaped)

		if length := len(chunks); length > 1 {

			buildTree(data, cache, chunks, value, depth)

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

func buildTree(data, cache QueryMap, chunks []string, value any, depth int) {
	length := len(chunks)
	ckey := strings.Join(chunks[:length-1], "/")

	if reflect.ValueOf(cache[ckey]).Kind() == reflect.Map {
		key := chunks[length-1]
		subCkey := ckey + "/" + key

		if _, ok := cache[subCkey]; !ok {
			cache[subCkey] = value
		}

		cache[ckey].(QueryMap)[key] = cache[subCkey]
		return
	}

	ckey = strings.Join(chunks, "/")

	if length > 1 {
		cache[ckey] = value

		if depth <= MAX_DEPTH {
			buildTree(data, cache, chunks[:length-1], QueryMap{chunks[length-1]: value}, depth+1)
		}

		return
	}

	cache[ckey] = value
	data[ckey] = cache[ckey]
}

func divideQueryParam(param string) ([]string, string) {
	reg, _ := regexp.Compile(`=(.*)$`)
	value := ""

	if matches := reg.FindStringSubmatch(param); len(matches) > 0 {
		if matches[1] != "" {
			value = matches[1]
			param = param[:len(param)-len(matches[0])]
		}
	}

	chunks := strings.FieldsFunc(param, func(r rune) bool {
		_, ok := querySeparators[r]
		return ok
	})

	return chunks, value
}
