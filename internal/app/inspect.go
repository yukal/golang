package app

import (
	"fmt"
	"reflect"
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
			text = inspectStruct(data, fieldsCount-1, depth)
		}

	case reflect.Slice:
		if length := data.Len(); length > 0 {
			text = inspectSlice(data, length-1, depth)
		}

	case reflect.Map:
		keys := data.MapKeys()

		if length := len(keys); length > 0 {
			text = inspectMap(data, keys, length-1, depth)
		}

	case reflect.String:
		text = fmt.Sprintf("%#v\n", data.String())

	case reflect.Interface:
		if value := reflect.ValueOf(data.Interface()); value.IsValid() {
			text = inspectRecursively(value, depth)
		} else {
			text = fmt.Sprintf("%#v", data.Interface())
		}
	}

	return text
}

func inspectStruct(data reflect.Value, index, depth int) (text string) {
	indent := strings.Repeat(" ", depth*2)

	field := data.Type().Field(index)
	value := reflect.Indirect(data).FieldByName(field.Name)

	if value.Kind() == reflect.Interface {
		value = reflect.ValueOf(value.Interface())
	}

	if index > 0 {
		text += inspectStruct(data, index-1, depth)
	}

	switch {
	case !value.IsValid():
		text += fmt.Sprintf("%s%s: %#v\n", indent, field.Name, value)

	case value.Kind() == reflect.String:
		text += fmt.Sprintf("%s%s: %#v\n", indent, field.Name, value.String())

	case IsPrimitive(value):
		text += fmt.Sprintf("%s%s: %v\n", indent, field.Name, value)

	default:
		text += fmt.Sprintf("%s%s:\n%s", indent, field.Name,
			inspectRecursively(value, depth+1))
	}

	return
}

func inspectSlice(data reflect.Value, index, depth int) (text string) {
	indent := strings.Repeat(" ", depth*2)
	value := data.Index(index)

	if value.Kind() == reflect.Interface {
		value = reflect.ValueOf(value.Interface())
	}

	if index > 0 {
		text += inspectSlice(data, index-1, depth)
	}

	switch {
	case !value.IsValid():
		text += fmt.Sprintf("%s%v: %#v\n", indent, index, value)

	case value.Kind() == reflect.String:
		text += fmt.Sprintf("%s%v: %#v\n", indent, index, value.String())

	case IsPrimitive(value):
		text += fmt.Sprintf("%s%v: %v\n", indent, index, value.Interface())

	default:
		text += fmt.Sprintf("%s%v:\n%s", indent, index,
			inspectRecursively(value, depth+1))
	}

	return
}

func inspectMap(data reflect.Value, keys []reflect.Value, index, depth int) (text string) {
	indent := strings.Repeat(" ", depth*2)
	value := data.MapIndex(keys[index])

	if value.Kind() == reflect.Interface {
		value = reflect.ValueOf(value.Interface())
	}

	if index > 0 {
		text += inspectMap(data, keys, index-1, depth)
	}

	switch {
	case !value.IsValid():
		text += fmt.Sprintf("%s%s: %#v\n", indent, keys[index], value)

	case value.Kind() == reflect.String:
		text += fmt.Sprintf("%s%s: %#v\n", indent, keys[index], value.String())

	case IsPrimitive(value):
		text += fmt.Sprintf("%s%s: %v\n", indent, keys[index], value.Interface())

	default:
		text += fmt.Sprintf("%s%s:\n%s", indent, keys[index],
			inspectRecursively(value, depth+1))
	}

	return
}
