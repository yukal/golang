package validation

import (
	"reflect"
	"strings"
)

type FilterMap map[string]map[string]any

func (filter FilterMap) IsValid(anyStruct any) bool {
	refVal := reflect.ValueOf(anyStruct)
	refType := refVal.Elem().Type()

	filterVal := reflect.ValueOf(filter)

	// v := filterVal.MapIndex(reflect.ValueOf("Id"))
	// fmt.Println("Id", v)

	// for _, e := range filterVal.MapKeys() {
	// 	v := filterVal.MapIndex(e)
	// 	fmt.Println("keys", e, v)
	// }

	for i := 0; i < refType.NumField(); i++ {

		field := refType.Field(i)
		value := reflect.Indirect(refVal).FieldByName(field.Name)

		// filterRefItem := reflect.Indirect(filterVal).FieldByName(field.Name)
		filterItem := filterVal.MapIndex(reflect.ValueOf(field.Name))
		// fmt.Println(field.Name, filterItem)

		switch filterItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			for _, rule := range filterItem.MapKeys() {
				filterVal := filterItem.MapIndex(rule)
				// fmt.Println("--", rule, filterVal)

				if !Compare(rule.String(), filterVal.Interface(), value.Interface()) {
					return false
				}
			}
			break
		}

	}

	return true
}

func (filter FilterMap) Validate(anyStruct any) (errList []string) {
	refVal := reflect.ValueOf(anyStruct)
	refType := refVal.Elem().Type()

	filterVal := reflect.ValueOf(filter)

	for i := 0; i < refType.NumField(); i++ {

		field := refType.Field(i)
		value := reflect.Indirect(refVal).
			FieldByName(field.Name)

		// filterItem := reflect.Indirect(filterVal).FieldByName(field.Name)
		filterItem := filterVal.MapIndex(reflect.ValueOf(field.Name))

		switch filterItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			res := true

			for _, rule := range filterItem.MapKeys() {
				filterVal := filterItem.MapIndex(rule)
				// fmt.Println("--", rule, filterVal)
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
