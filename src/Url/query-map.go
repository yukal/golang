package Url

import (
	"net/url"
	"regexp"
	"strings"
)

type QueryMap map[string]any
type QueryVal struct {
	key string
	val any
}

func NewQueryMap(urlQuery string) QueryMap {
	params := strings.Split(urlQuery, "?")
	paramsLength := len(params)

	if paramsLength == 1 {
		urlQuery = params[0]
	}
	if paramsLength > 1 {
		urlQuery = params[1]
	}

	params = strings.Split(urlQuery, "&")
	tree := make(QueryMap, paramsLength)

	for _, param := range params {
		key, val := queryChunks(param, 0)

		if item, ok := tree[key]; !ok {
			if val != nil {

				tree[key] = QueryMap{
					val.(QueryVal).key: val.(QueryVal).val,
				}

			}
		} else {
			k := val.(QueryVal).key
			v := val.(QueryVal).val

			if subItem, ok := item.(QueryMap)[k]; !ok {

				item.(QueryMap)[k] = v

			} else {
				for subKey, subVal := range v.(QueryMap) {

					subItem.(QueryMap)[subKey] = subVal

				}
			}
		}
	}

	return tree
}

func queryChunks(param string, depth uint16) (key string, val any) {
	key, val = unpackQueryKeyVal(param)

	if len(val.(string)) > 0 {
		if val.(string)[0] == '=' {
			val = val.(string)[1:]
		}

		if val.(string)[0] == '[' {
			k, v := queryChunks(val.(string), depth+1)

			if depth == 0 {
				val = QueryVal{key: k, val: v}
			} else {
				val = QueryMap{k: v}
			}
		}
	} else {
		val = nil
	}

	return key, val
}

func unpackQueryKeyVal(subParam string) (key string, val string) {
	if reg, err := regexp.Compile(`\[?(\w+)\]?`); err == nil {
		if matches := reg.FindStringSubmatch(subParam); len(matches) > 1 {
			key = matches[1]
			// val = subParam[len(matches[0]):]
			val, _ = url.QueryUnescape(subParam[len(matches[0]):])
		}
	}

	return
}
