package validation

import (
	"reflect"
	"strings"
)

type IFilterType interface {
	FilterMap | interface{}
}

type IFilterInstance interface {
	IsValid(anyStruct any) bool
	Validate(anyStruct any) []string
}

func IsValid[T IFilterType](filter T, anyStruct any) bool {
	refVal := reflect.ValueOf(anyStruct)
	refType := refVal.Elem().Type()

	filterVal := reflect.ValueOf(filter)
	filterKind := filterVal.Kind()

	if filterKind == reflect.Ptr {
		filterKind = filterVal.Elem().Kind()
	}

	for i := 0; i < refType.NumField(); i++ {
		var filterItem reflect.Value

		field := refType.Field(i)
		value := reflect.Indirect(refVal).FieldByName(field.Name)

		switch filterKind {
		case reflect.Struct:
			filterItem = reflect.Indirect(filterVal).FieldByName(field.Name)

		case reflect.Map:
			filterItem = filterVal.MapIndex(reflect.ValueOf(field.Name))
		}

		switch filterItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			for _, rule := range filterItem.MapKeys() {
				filterVal := filterItem.MapIndex(rule)

				if !Compare(rule.String(), filterVal.Interface(), value.Interface()) {
					return false
				}
			}
			break
		}

	}

	return true
}

func Validate[T IFilterType](filter T, anyStruct any) (errList []string) {
	refVal := reflect.ValueOf(anyStruct)
	refType := refVal.Elem().Type()

	filterVal := reflect.ValueOf(filter)
	filterKind := filterVal.Kind()

	if filterKind == reflect.Ptr {
		filterKind = filterVal.Elem().Kind()
	}

	for i := 0; i < refType.NumField(); i++ {
		var filterItem reflect.Value

		field := refType.Field(i)
		value := reflect.Indirect(refVal).
			FieldByName(field.Name)

		switch filterKind {
		case reflect.Struct:
			filterItem = reflect.Indirect(filterVal).FieldByName(field.Name)

		case reflect.Map:
			filterItem = filterVal.MapIndex(reflect.ValueOf(field.Name))
		}

		switch filterItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			res := true

			for _, rule := range filterItem.MapKeys() {
				filterVal := filterItem.MapIndex(rule)
				res = res && Compare(rule.String(), filterVal.Interface(), value.Interface())
			}

			if !res {
				errList = append(errList, strings.ToLower(field.Name))
			}

			break
		}

	}

	return
}
