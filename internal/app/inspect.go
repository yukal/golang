package app

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func InspectData(data any) string {
	return strings.TrimSpace(
		inspectRecursively(reflect.Indirect(reflect.ValueOf(data)), 0),
	)
}

func inspectRecursively(data reflect.Value, depth int) (text string) {
	switch data.Kind() {
	case reflect.Struct:
		if fieldsCount := data.Type().NumField(); fieldsCount > 0 {
			text = inspectStruct(data, 0, depth)
		}

	case reflect.Slice:
		if length := data.Len(); length > 0 {
			text = inspectSlice(data, 0, depth)
		}

	case reflect.Map:
		if keys := sortMapKeys(data.MapKeys()); len(keys) > 0 {
			text = inspectMap(data, keys, depth)
		}

	case reflect.Interface:
		text = inspectRecursively(reflect.ValueOf(data.Interface()), depth)

	default:
		text = toText("", data, depth)
	}

	return
}

func inspectStruct(data reflect.Value, index, depth int) (text string) {
	field := data.Type().Field(index)
	value := reflect.Indirect(data).FieldByName(field.Name)

	text = toText(field.Name, value, depth)

	if index < data.Type().NumField()-1 {
		text += inspectStruct(data, index+1, depth)
	}

	return
}

func inspectSlice(data reflect.Value, index, depth int) (text string) {
	text = toText(strconv.Itoa(index), data.Index(index), depth)

	if index < data.Len()-1 {
		text += inspectSlice(data, index+1, depth)
	}

	return
}

func inspectMap(data reflect.Value, keys []string, depth int) (text string) {
	value := data.MapIndex(reflect.ValueOf(keys[0]))
	text = toText(keys[0], value, depth)

	if len(keys) > 1 {
		text += inspectMap(data, keys[1:], depth)
	}

	return
}

func toText(key string, value reflect.Value, depth int) (text string) {
	indent := strings.Repeat(" ", depth*2)

	switch value.Kind() {
	case reflect.Invalid:
		text = fmt.Sprintf("%s%s: %#v\n", indent, key, value)

	case reflect.Bool:
		text = fmt.Sprintf("%s%s: %v\n", indent, key, value.Bool())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		text = fmt.Sprintf("%s%s: %v\n", indent, key, value.Int())

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		text = fmt.Sprintf("%s%s: %v\n", indent, key, value.Uint())

	case reflect.Float32, reflect.Float64:
		text = fmt.Sprintf("%s%s: %v\n", indent, key, value.Float())

	case reflect.Complex64, reflect.Complex128:
		text = fmt.Sprintf("%s%s: %v\n", indent, key, value.Complex())

	case reflect.String:
		text = fmt.Sprintf("%s%s: %#v\n", indent, key, value.String())

	case reflect.Chan:
		text = fmt.Sprintf("%s%s: [chan]\n", indent, key)

	case reflect.Func:
		text = fmt.Sprintf("%s%s: [func]\n", indent, key)

	case reflect.Uintptr, reflect.Pointer, reflect.UnsafePointer:
		text = fmt.Sprintf("%s%s: %#v\n", indent, key, value)

	case reflect.Interface:
		// text = fmt.Sprintf("%s%s: %v\n", indent, key, reflect.ValueOf(value.Interface()))
		text = toText(key, reflect.ValueOf(value.Interface()), depth)

	case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map:
		text = fmt.Sprintf("%s%s:\n%s", indent, key, inspectRecursively(value, depth+1))
	}

	return
}

func sortMapKeys(refKeys []reflect.Value) []string {
	sortKeys := make([]string, 0, len(refKeys))

	for _, v := range refKeys {
		sortKeys = append(sortKeys, v.String())
	}

	sort.Strings(sortKeys)
	return sortKeys
}
