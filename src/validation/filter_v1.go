package validation

import (
	"reflect"
	"strings"
)

type Filter struct {
	Id       map[string]uint64 `json:"Id"`
	RegionId map[string]uint8  `json:"RegionId"`
	Hash     map[string]uint8  `json:"Hash"`
	Link     map[string]uint8  `json:"Link"`
	Title    map[string]uint8  `json:"Title"`
	Message  map[string]uint8  `json:"Message"`
	Sex      map[string]uint8  `json:"Sex"`
	Age      map[string]uint8  `json:"Age"`
	Height   map[string]uint8  `json:"Height"`
	Weight   map[string]uint8  `json:"Weight"`
	Images   map[string]uint8  `json:"Images"`
	Phones   map[string]uint8  `json:"Phones"`
	Date     map[string]uint16 `json:"Date"`
}

func IsValid(filter *Filter, anyStruct any) bool {
	refVal := reflect.ValueOf(anyStruct)
	refType := refVal.Elem().Type()

	filterRefVal := reflect.ValueOf(filter)

	for i := 0; i < refType.NumField(); i++ {

		field := refType.Field(i)
		value := reflect.Indirect(refVal).FieldByName(field.Name)

		filterRefItem := reflect.Indirect(filterRefVal).FieldByName(field.Name)

		switch filterRefItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			for _, rule := range filterRefItem.MapKeys() {
				filterVal := filterRefItem.MapIndex(rule)

				if !Compare(rule.String(), filterVal.Interface(), value.Interface()) {
					return false
				}
			}
			break
		}

	}

	return true
}

func Validate(filter *Filter, anyStruct any) (errList []string) {
	refVal := reflect.ValueOf(anyStruct)
	refType := refVal.Elem().Type()

	filterRefVal := reflect.ValueOf(filter)

	for i := 0; i < refType.NumField(); i++ {

		field := refType.Field(i)
		value := reflect.Indirect(refVal).
			FieldByName(field.Name)

		filterRefItem := reflect.Indirect(filterRefVal).
			FieldByName(field.Name)

		switch filterRefItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			res := true

			for _, rule := range filterRefItem.MapKeys() {
				filterVal := filterRefItem.MapIndex(rule)
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
