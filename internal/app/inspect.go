package app

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const MAX_DEPTH = 16

func InspectData(data any) string {
	key := ""
	value := reflect.Indirect(reflect.ValueOf(data))

	if value.IsValid() {
		key = value.Type().String()
	} else {
		key = value.Kind().String()
	}

	return inspectRecursively(key, value, 0)
}

func inspectRecursively(key string, value reflect.Value, depth int) (text string) {
	separ := " "

	if value.Kind() == reflect.Interface {
		value = reflect.ValueOf(value.Interface())
	}

	switch value.Kind() {
	// case reflect.Interface:
	// 	text = fmt.Sprintf("%v\n", reflect.ValueOf(value.Interface()))
	// 	text = inspectRecursively("", reflect.ValueOf(value.Interface()), cmap, depth)

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Bool,
		reflect.Invalid:

		text = fmt.Sprintf("%v\n", value)

	case reflect.String, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		text = fmt.Sprintf("%#v\n", value)

	case reflect.Uintptr:
		text = fmt.Sprintf("uintptr(%#v)\n", value)

	case reflect.Pointer:
		// TODO: avoid infinite tree
		text = inspectRecursively(key, value.Elem(), depth)

	case reflect.Struct:
		// TODO: avoid infinite tree

		if value.Type().NumField() > 0 && depth < MAX_DEPTH {
			text = inspectStruct(value, 0, depth+1)
			separ = "\n"
		} else {
			text = value.Type().String() + "{}\n"
		}

	case reflect.Slice, reflect.Array:
		if value.Len() > 0 && depth < MAX_DEPTH {
			text = inspectSlice(value, 0, depth+1)
			separ = "\n"
		} else {
			text = value.Type().String() + "{}\n"
		}

	case reflect.Map:
		// TODO: avoid infinite tree

		if keys := value.MapKeys(); len(keys) > 0 && depth < MAX_DEPTH {
			text = inspectMap(value, sortMapKeys(keys), depth+1)
			separ = "\n"
		} else {
			text = value.Type().String() + "{}\n"
		}
	}

	indent := strings.Repeat(" ", depth*2)
	text = fmt.Sprintf("%s%s:%s%s", indent, key, separ, text)

	return
}

func inspectStruct(data reflect.Value, index, depth int) (text string) {
	field := data.Type().Field(index)
	value := reflect.Indirect(data).FieldByName(field.Name)

	text = inspectRecursively(field.Name, value, depth)

	if index < data.Type().NumField()-1 {
		text += inspectStruct(data, index+1, depth)
	}

	return
}

func inspectSlice(data reflect.Value, index, depth int) (text string) {
	text = inspectRecursively(strconv.Itoa(index), data.Index(index), depth)

	if index < data.Len()-1 {
		text += inspectSlice(data, index+1, depth)
	}

	return
}

func inspectMap(data reflect.Value, keys []string, depth int) (text string) {
	value := data.MapIndex(reflect.ValueOf(keys[0]))
	text = inspectRecursively(keys[0], value, depth)

	if len(keys) > 1 {
		text += inspectMap(data, keys[1:], depth)
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
