package validation

import (
	"reflect"
	"strings"
	"time"
	"yu/golang/src"
)

type IFilterType interface {
	FilterMap | src.ArticleFilter
}

type IFilterInstance interface {
	IsValid(anyStruct any) bool
	Validate(anyStruct any) []string
}

func IsValid[T IFilterType](filter T, anyStruct any) bool {
	flag, _ := checkIsValid(filter, anyStruct, true)
	return flag
}

func Validate[T IFilterType](filter T, anyStruct any) []string {
	_, wrongFields := checkIsValid(filter, anyStruct, false)
	return wrongFields
}

func checkIsValid[T IFilterType](filter T, anyStruct any, strict bool) (bool, []string) {
	var errList []string

	filterVal := reflect.Indirect(reflect.ValueOf(filter))
	structVal := reflect.Indirect(reflect.ValueOf(anyStruct))
	structType := structVal.Type()

	if !strict {
		errList = make([]string, 0, structType.NumField())
	}

	for i := 0; i < structType.NumField(); i++ {
		var filterItem reflect.Value

		field := structType.Field(i)
		value := structVal.FieldByName(field.Name)

		switch filterVal.Kind() {
		case reflect.Struct:
			filterItem = filterVal.FieldByName(field.Name)

		case reflect.Map:
			filterItem = filterVal.MapIndex(reflect.ValueOf(field.Name))
		}

		switch filterItem.Kind() {
		case reflect.Invalid:
			continue

		case reflect.Map:
			for _, rule := range filterItem.MapKeys() {
				ruleVal := filterItem.MapIndex(rule)
				// fmt.Println(field.Name, rule, ruleVal)

				if !Compare(rule.String(), ruleVal.Interface(), value.Interface()) {
					if strict {
						return false, errList
					}

					errList = append(errList, strings.ToLower(field.Name))
					break
				}
			}
			break
		}

	}

	return true, errList
}

func Compare(rule string, filterVal, comparableVal any) bool {
	switch rule {

	case "min":
		return isMin(filterVal, comparableVal)

	case "max":
		return isMax(filterVal, comparableVal)

	case "eq":
		return isEqual(filterVal, comparableVal)

	case "strLen":
		return isStrLenEqual(filterVal, comparableVal)

	case "minLen":
		return isStrLenMin(filterVal, comparableVal)

	case "maxLen":
		return isStrLenMax(filterVal, comparableVal)

	case "year":
		return isYearEqual(filterVal, comparableVal)

	}

	return true
}

func isEachStrLenEqual(filterVal, val any) bool {
	var result bool = len(val.([]string)) > 0

	for _, str := range val.([]string) {
		result = result && isEqual(filterVal, len(str))
	}

	return result
}

func isStrLenEqual(filterVal, val any) bool {
	switch reflect.TypeOf(val).String() {

	case "[]string":
		length := len(val.([]string))
		return isEqual(filterVal, length)

	case "string":
		length := len(val.(string))
		return isEqual(filterVal, length)

	}

	return false
}

func isStrLenMin(filterVal, val any) bool {
	switch reflect.TypeOf(val).String() {

	case "[]string":
		length := len(val.([]string))
		return isMin(filterVal, length)

	case "string":
		length := len(val.(string))
		return isMin(filterVal, length)

	}

	return false
}

func isStrLenMax(filterVal, val any) bool {
	switch reflect.TypeOf(val).String() {

	case "[]string":
		length := len(val.([]string))
		return isMax(filterVal, length)

	case "string":
		length := len(val.(string))
		return isMax(filterVal, length)

	}

	return false
}

func isYearEqual(filterVal, val any) bool {
	if reflect.TypeOf(val).String() != "time.Time" {
		return false
	}

	year := val.(time.Time).Year()
	return isEqual(filterVal, year)
}

// https://go.dev/ref/spec#Numeric_types
func isMin(filterVal, val any) bool {
	types := reflect.TypeOf(val).Kind().String() + ":" + reflect.TypeOf(filterVal).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return val.(int8) >= filterVal.(int8)

	case "int8:int16":
		return int16(val.(int8)) >= filterVal.(int16)

	case "int8:int32", "int8:rune":
		return int32(val.(int8)) >= filterVal.(int32)

	case "int8:int64":
		return int64(val.(int8)) >= filterVal.(int64)

	case "int8:int":
		return int(val.(int8)) >= filterVal.(int)

	case "int16:int8":
		return val.(int16) >= int16(filterVal.(int8))

	case "int16:int16":
		return val.(int16) >= filterVal.(int16)

	case "int16:int32", "int16:rune":
		return int32(val.(int16)) >= filterVal.(int32)

	case "int16:int64":
		return int64(val.(int16)) >= filterVal.(int64)

	case "int16:int":
		return int(val.(int16)) >= filterVal.(int)

	case "int32:int8", "rune:int8":
		return val.(int32) >= int32(filterVal.(int8))

	case "int32:int16", "rune:int16":
		return val.(int32) >= int32(filterVal.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return val.(int32) >= filterVal.(int32)

	case "int32:int64", "rune:int64":
		return int64(val.(int32)) >= filterVal.(int64)

	case "int32:int", "rune:int":
		return int(val.(int32)) >= filterVal.(int)

	case "int64:int8":
		return val.(int64) >= int64(filterVal.(int8))

	case "int64:int16":
		return val.(int64) >= int64(filterVal.(int16))

	case "int64:int32", "int64:rune":
		return val.(int64) >= int64(filterVal.(int32))

	case "int64:int64":
		return val.(int64) >= filterVal.(int64)

	case "int64:int":
		return val.(int64) >= int64(filterVal.(int))

	case "int:int8":
		return val.(int) >= int(filterVal.(int8))

	case "int:int16":
		return val.(int) >= int(filterVal.(int16))

	case "int:int32", "int:rune":
		return val.(int) >= int(filterVal.(int32))

	case "int:int64":
		return int64(val.(int)) >= filterVal.(int64)

	case "int:int":
		return val.(int) >= filterVal.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if val.(int8) >= 0 {
			return uint8(val.(int8)) >= filterVal.(uint8)
		}
		return false

	case "int8:uint16":
		if val.(int8) >= 0 {
			return uint16(val.(int8)) >= filterVal.(uint16)
		}
		return false

	case "int8:uint32":
		if val.(int8) >= 0 {
			return uint32(val.(int8)) >= filterVal.(uint32)
		}
		return false

	case "int8:uint64":
		if val.(int8) >= 0 {
			return uint64(val.(int8)) >= filterVal.(uint64)
		}
		return false

	case "int8:uint":
		if val.(int8) >= 0 {
			return uint(val.(int8)) >= filterVal.(uint)
		}
		return false

	// ...

	case "int16:uint8", "int16:byte":
		return val.(int16) >= int16(filterVal.(uint8))

	case "int16:uint16":
		if val.(int16) >= 0 {
			return uint16(val.(int16)) >= filterVal.(uint16)
		}
		return false

	case "int16:uint32":
		if val.(int16) >= 0 {
			return uint32(val.(int16)) >= filterVal.(uint32)
		}
		return false

	case "int16:uint64":
		if val.(int16) >= 0 {
			return uint64(val.(int16)) >= filterVal.(uint64)
		}
		return false

	case "int16:uint":
		if val.(int16) >= 0 {
			return uint(val.(int16)) >= filterVal.(uint)
		}
		return false

	// ...

	case "int32:uint8", "int32:byte":
		return val.(int32) >= int32(filterVal.(uint8))

	case "int32:uint16":
		return val.(int32) >= int32(filterVal.(uint16))

	case "int32:uint32":
		if val.(int32) >= 0 {
			return uint32(val.(int32)) >= filterVal.(uint32)
		}
		return false

	case "int32:uint64":
		if val.(int32) >= 0 {
			return uint64(val.(int32)) >= filterVal.(uint64)
		}
		return false

	case "int32:uint":
		if val.(int32) >= 0 {
			return uint(val.(int32)) >= filterVal.(uint)
		}
		return false

	// ...

	case "int64:uint8", "int64:byte":
		return val.(int64) >= int64(filterVal.(uint8))

	case "int64:uint16":
		return val.(int64) >= int64(filterVal.(uint16))

	case "int64:uint32":
		return val.(int64) >= int64(filterVal.(uint32))

	case "int64:uint64":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) >= uint64(filterVal.(uint64))
		}
		return false

	case "int64:uint":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) >= uint64(filterVal.(uint))
		}
		return false

	// ...

	case "int:uint8", "int:byte":
		return val.(int) >= int(filterVal.(uint8))

	case "int:uint16":
		return val.(int) >= int(filterVal.(uint16))

	case "int:uint32":
		if val.(int) >= 0 {
			return uint64(val.(int)) >= uint64(filterVal.(uint32))
		}
		return false

	case "int:uint64":
		if val.(int) >= 0 {
			return uint64(val.(int)) >= filterVal.(uint64)
		}
		return false

	case "int:uint":
		if val.(int) >= 0 {
			return uint(val.(int)) >= filterVal.(uint)
		}
		return false

	// ...

	case "rune:uint8", "rune:byte":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) >= uint32(filterVal.(uint8))
		}
		return false

	case "rune:uint16":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) >= uint32(filterVal.(uint16))
		}
		return false

	case "rune:uint32":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) >= filterVal.(uint32)
		}
		return false

	case "rune:uint64":
		if val.(rune) >= 0 {
			return uint64(val.(rune)) >= filterVal.(uint64)
		}
		return false

	case "rune:uint":
		if val.(rune) >= 0 {
			return uint(val.(rune)) >= filterVal.(uint)
		}
		return false

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return val.(uint8) >= filterVal.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(val.(uint8)) >= filterVal.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(val.(uint8)) >= filterVal.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(val.(uint8)) >= filterVal.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(val.(uint8)) >= filterVal.(uint)

	case "uint16:uint8", "uint16:byte":
		return val.(uint16) >= uint16(filterVal.(uint8))

	case "uint16:uint16":
		return val.(uint16) >= filterVal.(uint16)
	case "uint16:uint32":
		return uint32(val.(uint16)) >= filterVal.(uint32)
	case "uint16:uint64":
		return uint64(val.(uint16)) >= filterVal.(uint64)
	case "uint16:uint":
		return uint(val.(uint16)) >= filterVal.(uint)

	case "uint32:uint8", "uint32:byte":
		return val.(uint32) >= uint32(filterVal.(uint8))
	case "uint32:uint16":
		return val.(uint32) >= uint32(filterVal.(uint16))
	case "uint32:uint32":
		return val.(uint32) >= filterVal.(uint32)
	case "uint32:uint64":
		return uint64(val.(uint32)) >= filterVal.(uint64)
	case "uint32:uint":
		return uint(val.(uint32)) >= filterVal.(uint)

	case "uint64:uint8", "uint64:byte":
		return val.(uint64) >= uint64(filterVal.(uint8))
	case "uint64:uint16":
		return val.(uint64) >= uint64(filterVal.(uint16))
	case "uint64:uint32":
		return val.(uint64) >= uint64(filterVal.(uint32))
	case "uint64:uint64":
		return val.(uint64) >= filterVal.(uint64)
	case "uint64:uint":
		return val.(uint64) >= uint64(filterVal.(uint))

	case "uint:uint8", "uint:byte":
		return val.(uint) >= uint(filterVal.(uint8))
	case "uint:uint16":
		return val.(uint) >= uint(filterVal.(uint16))
	case "uint:uint32":
		return val.(uint) >= uint(filterVal.(uint32))
	case "uint:uint64":
		return uint64(val.(uint)) >= filterVal.(uint64)
	case "uint:uint":
		return val.(uint) >= filterVal.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint8) >= uint8(filterVal.(int8))
		}
		return true

	case "uint8:int16", "byte:int16":
		return int16(val.(uint8)) >= filterVal.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(val.(uint8)) >= filterVal.(int32)

	case "uint8:int64", "byte:int64":
		return int64(val.(uint8)) >= filterVal.(int64)

	case "uint8:int", "byte:int":
		return int(val.(uint8)) >= filterVal.(int)

	// ...

	case "uint16:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint16) >= uint16(filterVal.(int8))
		}
		return true

	case "uint16:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint16) >= uint16(filterVal.(int16))
		}
		return true

	case "uint16:int32", "uint16:rune":
		return int32(val.(uint16)) >= filterVal.(int32)

	case "uint16:int64":
		return int64(val.(uint16)) >= filterVal.(int64)

	case "uint16:int":
		return int(val.(uint16)) >= filterVal.(int)

	// ...

	case "uint32:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint32) >= uint32(filterVal.(int8))
		}
		return true

	case "uint32:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint32) >= uint32(filterVal.(int16))
		}
		return true

	case "uint32:int32", "uint32:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint32) >= uint32(filterVal.(int32))
		}
		return true

	case "uint32:int64":
		return int64(val.(uint32)) >= filterVal.(int64)

	case "uint32:int":
		return int64(val.(uint32)) >= int64(filterVal.(int))

	// ...

	case "uint64:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint64) >= uint64(filterVal.(int8))
		}
		return true

	case "uint64:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint64) >= uint64(filterVal.(int16))
		}
		return true

	case "uint64:int32", "uint64:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint64) >= uint64(filterVal.(int32))
		}
		return true

	case "uint64:int64":
		if filterVal.(int64) >= 0 {
			return val.(uint64) >= uint64(filterVal.(int64))
		}
		return true

	case "uint64:int":
		if filterVal.(int) >= 0 {
			return val.(uint64) >= uint64(filterVal.(int))
		}
		return true

	// ...

	case "uint:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint) >= uint(filterVal.(int8))
		}
		return true

	case "uint:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint) >= uint(filterVal.(int16))
		}
		return true

	case "uint:int32":
		if filterVal.(int32) >= 0 {
			return val.(uint) >= uint(filterVal.(int32))
		}
		return true

	case "uint:int64":
		if filterVal.(int64) >= 0 {
			return uint64(val.(uint)) >= uint64(filterVal.(int64))
		}
		return true

	case "uint:int":
		if filterVal.(int) >= 0 {
			return val.(uint) >= uint(filterVal.(int))
		}
		return true

	}

	return false
}

func isMax(filterVal, val any) bool {
	types := reflect.TypeOf(val).Kind().String() + ":" + reflect.TypeOf(filterVal).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return val.(int8) <= filterVal.(int8)

	case "int8:int16":
		return int16(val.(int8)) <= filterVal.(int16)

	case "int8:int32", "int8:rune":
		return int32(val.(int8)) <= filterVal.(int32)

	case "int8:int64":
		return int64(val.(int8)) <= filterVal.(int64)

	case "int8:int":
		return int(val.(int8)) <= filterVal.(int)

	case "int16:int8":
		return val.(int16) <= int16(filterVal.(int8))

	case "int16:int16":
		return val.(int16) <= filterVal.(int16)

	case "int16:int32", "int16:rune":
		return int32(val.(int16)) <= filterVal.(int32)

	case "int16:int64":
		return int64(val.(int16)) <= filterVal.(int64)

	case "int16:int":
		return int(val.(int16)) <= filterVal.(int)

	case "int32:int8", "rune:int8":
		return val.(int32) <= int32(filterVal.(int8))

	case "int32:int16", "rune:int16":
		return val.(int32) <= int32(filterVal.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return val.(int32) <= filterVal.(int32)

	case "int32:int64", "rune:int64":
		return int64(val.(int32)) <= filterVal.(int64)

	case "int32:int", "rune:int":
		return int(val.(int32)) <= filterVal.(int)

	case "int64:int8":
		return val.(int64) <= int64(filterVal.(int8))

	case "int64:int16":
		return val.(int64) <= int64(filterVal.(int16))

	case "int64:int32", "int64:rune":
		return val.(int64) <= int64(filterVal.(int32))

	case "int64:int64":
		return val.(int64) <= filterVal.(int64)

	case "int64:int":
		return val.(int64) <= int64(filterVal.(int))

	case "int:int8":
		return val.(int) <= int(filterVal.(int8))

	case "int:int16":
		return val.(int) <= int(filterVal.(int16))

	case "int:int32", "int:rune":
		return val.(int) <= int(filterVal.(int32))

	case "int:int64":
		return int64(val.(int)) <= filterVal.(int64)

	case "int:int":
		return val.(int) <= filterVal.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if val.(int8) >= 0 {
			return uint8(val.(int8)) <= filterVal.(uint8)
		}
		return true

	case "int8:uint16":
		if val.(int8) >= 0 {
			return uint16(val.(int8)) <= filterVal.(uint16)
		}
		return true

	case "int8:uint32":
		if val.(int8) >= 0 {
			return uint32(val.(int8)) <= filterVal.(uint32)
		}
		return true

	case "int8:uint64":
		if val.(int8) >= 0 {
			return uint64(val.(int8)) <= filterVal.(uint64)
		}
		return true

	case "int8:uint":
		if val.(int8) >= 0 {
			return uint(val.(int8)) <= filterVal.(uint)
		}
		return true

	// ...

	case "int16:uint8", "int16:byte":
		return val.(int16) <= int16(filterVal.(uint8))

	case "int16:uint16":
		if val.(int16) >= 0 {
			return uint16(val.(int16)) <= filterVal.(uint16)
		}
		return true

	case "int16:uint32":
		if val.(int16) >= 0 {
			return uint32(val.(int16)) <= filterVal.(uint32)
		}
		return true

	case "int16:uint64":
		if val.(int16) >= 0 {
			return uint64(val.(int16)) <= filterVal.(uint64)
		}
		return true

	case "int16:uint":
		if val.(int16) >= 0 {
			return uint(val.(int16)) <= filterVal.(uint)
		}
		return true

	// ...

	case "int32:uint8", "int32:byte":
		return val.(int32) <= int32(filterVal.(uint8))

	case "int32:uint16":
		return val.(int32) <= int32(filterVal.(uint16))

	case "int32:uint32":
		if val.(int32) >= 0 {
			return uint32(val.(int32)) <= filterVal.(uint32)
		}
		return true

	case "int32:uint64":
		if val.(int32) >= 0 {
			return uint64(val.(int32)) <= filterVal.(uint64)
		}
		return true

	case "int32:uint":
		if val.(int32) >= 0 {
			return uint(val.(int32)) <= filterVal.(uint)
		}
		return true

	// ...

	case "int64:uint8", "int64:byte":
		return val.(int64) <= int64(filterVal.(uint8))

	case "int64:uint16":
		return val.(int64) <= int64(filterVal.(uint16))

	case "int64:uint32":
		return val.(int64) <= int64(filterVal.(uint32))

	case "int64:uint64":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) <= uint64(filterVal.(uint64))
		}
		return true

	case "int64:uint":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) <= uint64(filterVal.(uint))
		}
		return true

	// ...

	case "int:uint8", "int:byte":
		return val.(int) <= int(filterVal.(uint8))

	case "int:uint16":
		return val.(int) <= int(filterVal.(uint16))

	case "int:uint32":
		if val.(int) >= 0 {
			return uint(val.(int)) <= uint(filterVal.(uint32))
		}
		return true

	case "int:uint64":
		if val.(int) >= 0 {
			return uint64(val.(int)) <= filterVal.(uint64)
		}
		return true

	case "int:uint":
		if val.(int) >= 0 {
			return uint(val.(int)) <= filterVal.(uint)
		}
		return true

	// ...

	case "rune:uint8", "rune:byte":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) <= uint32(filterVal.(uint8))
		}
		return false

	case "rune:uint16":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) <= uint32(filterVal.(uint16))
		}
		return false

	case "rune:uint32":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) <= filterVal.(uint32)
		}
		return false

	case "rune:uint64":
		if val.(rune) >= 0 {
			return uint64(val.(rune)) <= filterVal.(uint64)
		}
		return false

	case "rune:uint":
		if val.(rune) >= 0 {
			return uint(val.(rune)) <= filterVal.(uint)
		}
		return false

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return val.(uint8) <= filterVal.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(val.(uint8)) <= filterVal.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(val.(uint8)) <= filterVal.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(val.(uint8)) <= filterVal.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(val.(uint8)) <= filterVal.(uint)

	case "uint16:uint8", "uint16:byte":
		return val.(uint16) <= uint16(filterVal.(uint8))

	case "uint16:uint16":
		return val.(uint16) <= filterVal.(uint16)
	case "uint16:uint32":
		return uint32(val.(uint16)) <= filterVal.(uint32)
	case "uint16:uint64":
		return uint64(val.(uint16)) <= filterVal.(uint64)
	case "uint16:uint":
		return uint(val.(uint16)) <= filterVal.(uint)

	case "uint32:uint8", "uint32:byte":
		return val.(uint32) <= uint32(filterVal.(uint8))
	case "uint32:uint16":
		return val.(uint32) <= uint32(filterVal.(uint16))
	case "uint32:uint32":
		return val.(uint32) <= filterVal.(uint32)
	case "uint32:uint64":
		return uint64(val.(uint32)) <= filterVal.(uint64)
	case "uint32:uint":
		return uint(val.(uint32)) <= filterVal.(uint)

	case "uint64:uint8", "uint64:byte":
		return val.(uint64) <= uint64(filterVal.(uint8))
	case "uint64:uint16":
		return val.(uint64) <= uint64(filterVal.(uint16))
	case "uint64:uint32":
		return val.(uint64) <= uint64(filterVal.(uint32))
	case "uint64:uint64":
		return val.(uint64) <= filterVal.(uint64)
	case "uint64:uint":
		return val.(uint64) <= uint64(filterVal.(uint))

	case "uint:uint8", "uint:byte":
		return val.(uint) <= uint(filterVal.(uint8))
	case "uint:uint16":
		return val.(uint) <= uint(filterVal.(uint16))
	case "uint:uint32":
		return val.(uint) <= uint(filterVal.(uint32))
	case "uint:uint64":
		return uint64(val.(uint)) <= filterVal.(uint64)
	case "uint:uint":
		return val.(uint) <= filterVal.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint8) <= uint8(filterVal.(int8))
		}
		return false

	case "uint8:int16", "byte:int16":
		return int16(val.(uint8)) <= filterVal.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(val.(uint8)) <= filterVal.(int32)

	case "uint8:int64", "byte:int64":
		return int64(val.(uint8)) <= filterVal.(int64)

	case "uint8:int", "byte:int":
		return int(val.(uint8)) <= filterVal.(int)

	// ...

	case "uint16:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint16) <= uint16(filterVal.(int8))
		}
		return false

	case "uint16:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint16) <= uint16(filterVal.(int16))
		}
		return false

	case "uint16:int32", "uint16:rune":
		return int32(val.(uint16)) <= filterVal.(int32)

	case "uint16:int64":
		return int64(val.(uint16)) <= filterVal.(int64)

	case "uint16:int":
		return int(val.(uint16)) <= filterVal.(int)

	// ...

	case "uint32:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint32) <= uint32(filterVal.(int8))
		}
		return false

	case "uint32:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint32) <= uint32(filterVal.(int16))
		}
		return false

	case "uint32:int32", "uint32:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint32) <= uint32(filterVal.(int32))
		}
		return false

	case "uint32:int64":
		return int64(val.(uint32)) <= filterVal.(int64)

	case "uint32:int":
		return int64(val.(uint32)) <= int64(filterVal.(int))

	// ...

	case "uint64:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint64) <= uint64(filterVal.(int8))
		}
		return false

	case "uint64:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint64) <= uint64(filterVal.(int16))
		}
		return false

	case "uint64:int32", "uint64:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint64) <= uint64(filterVal.(int32))
		}
		return false

	case "uint64:int64":
		if filterVal.(int64) >= 0 {
			return val.(uint64) <= uint64(filterVal.(int64))
		}
		return false

	case "uint64:int":
		if filterVal.(int) >= 0 {
			return val.(uint64) <= uint64(filterVal.(int))
		}
		return false

	// ...

	case "uint:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint) <= uint(filterVal.(int8))
		}
		return false

	case "uint:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint) <= uint(filterVal.(int16))
		}
		return false

	case "uint:int32", "uint:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint) <= uint(filterVal.(int32))
		}
		return false

	case "uint:int64":
		if filterVal.(int64) >= 0 {
			return uint64(val.(uint)) <= uint64(filterVal.(int64))
		}
		return false

	case "uint:int":
		if filterVal.(int) >= 0 {
			return val.(uint) <= uint(filterVal.(int))
		}
		return false

	}

	return false
}

func isEqual(filterVal, val any) bool {
	types := reflect.TypeOf(val).Kind().String() + ":" + reflect.TypeOf(filterVal).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return val.(int8) == filterVal.(int8)

	case "int8:int16":
		return int16(val.(int8)) == filterVal.(int16)

	case "int8:int32", "int8:rune":
		return int32(val.(int8)) == filterVal.(int32)

	case "int8:int64":
		return int64(val.(int8)) == filterVal.(int64)

	case "int8:int":
		return int(val.(int8)) == filterVal.(int)

	case "int16:int8":
		return val.(int16) == int16(filterVal.(int8))

	case "int16:int16":
		return val.(int16) == filterVal.(int16)

	case "int16:int32", "int16:rune":
		return int32(val.(int16)) == filterVal.(int32)

	case "int16:int64":
		return int64(val.(int16)) == filterVal.(int64)

	case "int16:int":
		return int(val.(int16)) == filterVal.(int)

	case "int32:int8", "rune:int8":
		return val.(int32) == int32(filterVal.(int8))

	case "int32:int16", "rune:int16":
		return val.(int32) == int32(filterVal.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return val.(int32) == filterVal.(int32)

	case "int32:int64", "rune:int64":
		return int64(val.(int32)) == filterVal.(int64)

	case "int32:int", "rune:int":
		return int(val.(int32)) == filterVal.(int)

	case "int64:int8":
		return val.(int64) == int64(filterVal.(int8))

	case "int64:int16":
		return val.(int64) == int64(filterVal.(int16))

	case "int64:int32", "int64:rune":
		return val.(int64) == int64(filterVal.(int32))

	case "int64:int64":
		return val.(int64) == filterVal.(int64)

	case "int64:int":
		return val.(int64) == int64(filterVal.(int))

	case "int:int8":
		return val.(int) == int(filterVal.(int8))

	case "int:int16":
		return val.(int) == int(filterVal.(int16))

	case "int:int32", "int:rune":
		return val.(int) == int(filterVal.(int32))

	case "int:int64":
		return int64(val.(int)) == filterVal.(int64)

	case "int:int":
		return val.(int) == filterVal.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if val.(int8) >= 0 {
			return uint8(val.(int8)) == filterVal.(uint8)
		}

	case "int8:uint16":
		if val.(int8) >= 0 {
			return uint16(val.(int8)) == filterVal.(uint16)
		}

	case "int8:uint32":
		if val.(int8) >= 0 {
			return uint32(val.(int8)) == filterVal.(uint32)
		}

	case "int8:uint64":
		if val.(int8) >= 0 {
			return uint64(val.(int8)) == filterVal.(uint64)
		}

	case "int8:uint":
		if val.(int8) >= 0 {
			return uint(val.(int8)) == filterVal.(uint)
		}

	// ...

	case "int16:uint8", "int16:byte":
		return val.(int16) == int16(filterVal.(uint8))

	case "int16:uint16":
		if val.(int16) >= 0 {
			return uint16(val.(int16)) == filterVal.(uint16)
		}

	case "int16:uint32":
		if val.(int16) >= 0 {
			return uint32(val.(int16)) == filterVal.(uint32)
		}

	case "int16:uint64":
		if val.(int16) >= 0 {
			return uint64(val.(int16)) == filterVal.(uint64)
		}

	case "int16:uint":
		if val.(int16) >= 0 {
			return uint(val.(int16)) == filterVal.(uint)
		}

	// ...

	case "int32:uint8", "int32:byte":
		return val.(int32) == int32(filterVal.(uint8))

	case "int32:uint16":
		return val.(int32) == int32(filterVal.(uint16))

	case "int32:uint32":
		if val.(int32) >= 0 {
			return uint32(val.(int32)) == filterVal.(uint32)
		}

	case "int32:uint64":
		if val.(int32) >= 0 {
			return uint64(val.(int32)) == filterVal.(uint64)
		}

	case "int32:uint":
		if val.(int32) >= 0 {
			return uint(val.(int32)) == filterVal.(uint)
		}

	// ...

	case "int64:uint8", "int64:byte":
		return val.(int64) == int64(filterVal.(uint8))

	case "int64:uint16":
		return val.(int64) == int64(filterVal.(uint16))

	case "int64:uint32":
		return val.(int64) == int64(filterVal.(uint32))

	case "int64:uint64":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) == uint64(filterVal.(uint64))
		}

	case "int64:uint":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) == uint64(filterVal.(uint))
		}

	// ...

	case "int:uint8", "int:byte":
		return val.(int) == int(filterVal.(uint8))

	case "int:uint16":
		return val.(int) == int(filterVal.(uint16))

	case "int:uint32":
		if val.(int) >= 0 {
			return uint(val.(int)) == uint(filterVal.(uint32))
		}

	case "int:uint64":
		if val.(int) >= 0 {
			return uint64(val.(int)) == filterVal.(uint64)
		}

	case "int:uint":
		if val.(int) >= 0 {
			return uint(val.(int)) == filterVal.(uint)
		}

	// ...

	case "rune:uint8", "rune:byte":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) == uint32(filterVal.(uint8))
		}

	case "rune:uint16":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) == uint32(filterVal.(uint16))
		}

	case "rune:uint32":
		if val.(rune) >= 0 {
			return uint32(val.(rune)) == filterVal.(uint32)
		}

	case "rune:uint64":
		if val.(rune) >= 0 {
			return uint64(val.(rune)) == filterVal.(uint64)
		}

	case "rune:uint":
		if val.(rune) >= 0 {
			return uint(val.(rune)) == filterVal.(uint)
		}

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return val.(uint8) == filterVal.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(val.(uint8)) == filterVal.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(val.(uint8)) == filterVal.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(val.(uint8)) == filterVal.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(val.(uint8)) == filterVal.(uint)

	case "uint16:uint8", "uint16:byte":
		return val.(uint16) == uint16(filterVal.(uint8))

	case "uint16:uint16":
		return val.(uint16) == filterVal.(uint16)
	case "uint16:uint32":
		return uint32(val.(uint16)) == filterVal.(uint32)
	case "uint16:uint64":
		return uint64(val.(uint16)) == filterVal.(uint64)
	case "uint16:uint":
		return uint(val.(uint16)) == filterVal.(uint)

	case "uint32:uint8", "uint32:byte":
		return val.(uint32) == uint32(filterVal.(uint8))
	case "uint32:uint16":
		return val.(uint32) == uint32(filterVal.(uint16))
	case "uint32:uint32":
		return val.(uint32) == filterVal.(uint32)
	case "uint32:uint64":
		return uint64(val.(uint32)) == filterVal.(uint64)
	case "uint32:uint":
		return uint(val.(uint32)) == filterVal.(uint)

	case "uint64:uint8", "uint64:byte":
		return val.(uint64) == uint64(filterVal.(uint8))
	case "uint64:uint16":
		return val.(uint64) == uint64(filterVal.(uint16))
	case "uint64:uint32":
		return val.(uint64) == uint64(filterVal.(uint32))
	case "uint64:uint64":
		return val.(uint64) == filterVal.(uint64)
	case "uint64:uint":
		return val.(uint64) == uint64(filterVal.(uint))

	case "uint:uint8", "uint:byte":
		return val.(uint) == uint(filterVal.(uint8))
	case "uint:uint16":
		return val.(uint) == uint(filterVal.(uint16))
	case "uint:uint32":
		return val.(uint) == uint(filterVal.(uint32))
	case "uint:uint64":
		return uint64(val.(uint)) == filterVal.(uint64)
	case "uint:uint":
		return val.(uint) == filterVal.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint8) == uint8(filterVal.(int8))
		}

	case "uint8:int16", "byte:int16":
		return int16(val.(uint8)) == filterVal.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(val.(uint8)) == filterVal.(int32)

	case "uint8:int64", "byte:int64":
		return int64(val.(uint8)) == filterVal.(int64)

	case "uint8:int", "byte:int":
		return int(val.(uint8)) == filterVal.(int)

	// ...

	case "uint16:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint16) == uint16(filterVal.(int8))
		}

	case "uint16:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint16) == uint16(filterVal.(int16))
		}

	case "uint16:int32", "uint16:rune":
		return int32(val.(uint16)) == filterVal.(int32)

	case "uint16:int64":
		return int64(val.(uint16)) == filterVal.(int64)

	case "uint16:int":
		return int(val.(uint16)) == filterVal.(int)

	// ...

	case "uint32:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint32) == uint32(filterVal.(int8))
		}

	case "uint32:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint32) == uint32(filterVal.(int16))
		}

	case "uint32:int32", "uint32:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint32) == uint32(filterVal.(int32))
		}

	case "uint32:int64":
		return int64(val.(uint32)) == filterVal.(int64)

	case "uint32:int":
		return int64(val.(uint32)) == int64(filterVal.(int))

	// ...

	case "uint64:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint64) == uint64(filterVal.(int8))
		}

	case "uint64:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint64) == uint64(filterVal.(int16))
		}

	case "uint64:int32", "uint64:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint64) == uint64(filterVal.(int32))
		}

	case "uint64:int64":
		if filterVal.(int64) >= 0 {
			return val.(uint64) == uint64(filterVal.(int64))
		}

	case "uint64:int":
		if filterVal.(int) >= 0 {
			return val.(uint64) == uint64(filterVal.(int))
		}

	// ...

	case "uint:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint) == uint(filterVal.(int8))
		}

	case "uint:int16":
		if filterVal.(int16) >= 0 {
			return val.(uint) == uint(filterVal.(int16))
		}

	case "uint:int32", "uint:rune":
		if filterVal.(int32) >= 0 {
			return val.(uint) == uint(filterVal.(int32))
		}

	case "uint:int64":
		if filterVal.(int64) >= 0 {
			return uint64(val.(uint)) == uint64(filterVal.(int64))
		}

	case "uint:int":
		if filterVal.(int) >= 0 {
			return val.(uint) == uint(filterVal.(int))
		}

	}

	return false
}
