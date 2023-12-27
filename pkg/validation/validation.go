package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	NON_ZERO  = "NonZero"
	NON_EMPTY = "NonEmpty"

	MsgMinFields   = "must contain at least %d valid fields"
	MsgMaxFields   = "must contain up to %d valid fields"
	MsgMinStrLen   = "must contain at least %d characters"
	MsgMaxStrLen   = "must contain up to %d characters"
	MsgEqStrLen    = "must contain exactly %d characters"
	MsgRangeStrLen = "must contain %d..%d characters"
	MsgMinSetLen   = "must contain at least %d items"
	MsgMaxSetLen   = "must contain up to %d items"
	MsgEqSetLen    = "must contain exactly %d items"
	MsgRangeSetLen = "must contain %d..%d items"
	MsgMin         = "must be at least %d"
	MsgMax         = "must be up to %d"
	MsgEq          = "must be exactly %d"
	MsgRange       = "must be in the range %d..%d"
	MsgNotValid    = "is not valid"
	MsgEmpty       = "is empty"
	// MsgNotMatch  = "is not match"
	MsgInvalidValue = "invalid value"
)

type Group []any
type Range [2]any
type Rule [2]any

type FilterItem struct {
	Field    string
	Check    any
	Optional bool
	// Unsafe   bool
}

type Filter []FilterItem

// data should be a type of struct{ ... }
func (filter Filter) IsValid(data any) bool {
	return len(checkIsValid(filter, data)) == 0
}

// data should be a type of struct{ ... }
func (filter Filter) Validate(data any) []string {
	return checkIsValid(filter, data)
}

func checkIsValid(filter Filter, data any) []string {
	refValData := reflect.Indirect(reflect.ValueOf(data))
	refTypData := refValData.Type()

	hints := make([]string, 0, refTypData.NumField())
	successFields := 0

	for _, filterStruct := range filter {
		// field := refTypData.Field(i)
		// field, fieldExist := refTypData.FieldByName(filterStruct.Field)
		rules := reflect.Indirect(reflect.ValueOf(filterStruct.Check))

		if field, exist := refTypData.FieldByName(filterStruct.Field); exist {
			tagName := filterStruct.Field
			// tagName = field.Tag.Get("json")

			if tagNameJson, exist := field.Tag.Lookup("json"); exist {
				tagName = tagNameJson
			}

			value := refValData.FieldByName(filterStruct.Field)

			if filterStruct.Optional && value.IsZero() {
				continue
			}

			if hint := checkField(rules, value); hint != "" {
				hints = append(hints, tagName+" "+hint)
			} else {
				successFields++
			}

			continue
		}

		var (
			action = ""
			value  = reflect.ValueOf(nil)
			proto  reflect.Value
		)

		switch rules.Type().String() {
		// case "string":
		// case "validation.Group":
		case "validation.Rule":
			action = rules.Index(0).Elem().String()
			proto = rules.Index(1).Elem()

			if strings.Contains(action, "-fields") {
				value = reflect.ValueOf(successFields)
			}

			if hint := Compare(action, proto, value); hint != "" {
				hints = append(hints, "body "+hint)
			}
		}
	}

	return hints
}

func checkField(rules, value reflect.Value) string {
	switch rules.Type().String() {
	case "validation.Group":

		for n := 0; n < rules.Len(); n++ {
			item := rules.Index(n)

			if item.Kind() == reflect.Interface {
				item = reflect.Indirect(reflect.ValueOf(
					item.Interface(),
				))
			}

			switch item.Type().String() {
			case "validation.Range":
				if hint := Compare("range", rules, value); hint != "" {
					return hint
				}

			case "validation.Rule":
				action := item.Index(0).Elem().String()
				proto := item.Index(1).Elem()

				if hint := Compare(action, proto, value); hint != "" {
					return hint
				}

			case "string":
				action := item.Elem().String()
				proto := reflect.ValueOf(nil)

				if hint := Compare(action, proto, value); hint != "" {
					return hint
				}
			}
		}

	case "validation.Range":
		return Compare("range", rules, value)

	case "validation.Rule":
		action := rules.Index(0).Elem().String()
		proto := rules.Index(1).Elem()

		return Compare(action, proto, value)

	case "string":
		action := rules.String()
		proto := reflect.ValueOf(nil)

		return Compare(action, proto, value)
	}

	return ""
}

func Compare(action string, proto, value reflect.Value) string {
	if !value.IsValid() {
		return MsgInvalidValue
	}

	switch action {
	case NON_ZERO:
		if value.IsZero() {
			return MsgEmpty
		}
		return ""
	}

	if !proto.IsValid() {
		return MsgInvalidValue
	}

	switch action {
	case "min-fields":
		if !IsMin(proto.Interface(), value.Interface()) {
			return fmt.Sprintf(MsgMinFields, proto.Interface())
		}

	case "range":
		return filterRange(proto, value)

	case "min":
		return filterMin(proto, value)

	case "max":
		return filterMax(proto, value)

	case "eq":
		return filterEq(proto, value)

	case "match":
		if !IsMatch(proto.Interface(), value.Interface()) {
			return MsgNotValid
		}

	case "eachMatch":
		if !IsEachMatches(proto.Interface(), value.Interface()) {
			return MsgNotValid
		}

	case "eachMin":
		success := true
		for n := 0; n < value.Len(); n++ {
			success = success && IsMin(proto.Interface(), value.Index(n).Interface())
		}

		if !success {
			return fmt.Sprintf(MsgMin, proto.Interface())
		}

	case "eachMax":
		success := true
		for n := 0; n < value.Len(); n++ {
			success = success && IsMax(proto.Interface(), value.Index(n).Interface())
		}

		if !success {
			return fmt.Sprintf(MsgMax, proto.Interface())
		}

	case "eachEq":
		success := true
		for n := 0; n < value.Len(); n++ {
			success = success && IsEqual(proto.Interface(), value.Index(n).Interface())
		}

		if !success {
			return fmt.Sprintf(MsgEq, proto.Interface())
		}

	case "eachMinLen":
		success := true
		for n := 0; n < value.Len(); n++ {
			success = success && IsMin(proto.Interface(), value.Index(n).Interface())
		}

		if !success {
			return fmt.Sprintf(MsgMinStrLen, proto.Interface())
		}

	case "eachMaxLen":
		success := true
		for n := 0; n < value.Len(); n++ {
			success = success && IsMax(proto.Interface(), value.Index(n).Interface())
		}

		if !success {
			return fmt.Sprintf(MsgMaxStrLen, proto.Interface())
		}

	case "eachLen":
		success := true
		for n := 0; n < value.Len(); n++ {
			success = success && IsEqual(proto.Interface(), value.Index(n).Interface())
		}

		if !success {
			return fmt.Sprintf(MsgEqStrLen, proto.Interface())
		}

	case "year":
		if !IsYearEqual(proto.Interface(), value.Interface()) {
			return fmt.Sprintf(MsgEq, proto.Interface())
		}
	}

	return ""
}

func filterRange(proto, value reflect.Value) string {
	hint := ""

	valMin := proto.Index(0)
	valMax := proto.Index(1)

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgRangeStrLen, valMin.Interface(), valMax.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		value = reflect.ValueOf(value.Len())
		hint = fmt.Sprintf(MsgRangeSetLen, valMin.Interface(), valMax.Interface())

	default:
		hint = fmt.Sprintf(MsgRange, valMin.Interface(), valMax.Interface())
	}

	if !IsMin(valMin.Interface(), value.Interface()) {
		return hint
	}

	if !IsMax(valMax.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func filterMin(proto, value reflect.Value) string {
	hint := ""

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgMinStrLen, proto.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		if value.IsZero() {
			value = reflect.ValueOf(0)
		} else {
			value = reflect.ValueOf(value.Len())
		}

		hint = fmt.Sprintf(MsgMinSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgMin, proto.Interface())
	}

	if !IsMin(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func filterMax(proto, value reflect.Value) string {
	hint := ""

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgMaxStrLen, proto.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		if value.IsZero() {
			value = reflect.ValueOf(0)
		} else {
			value = reflect.ValueOf(value.Len())
		}

		hint = fmt.Sprintf(MsgMaxSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgMax, proto.Interface())
	}

	if !IsMax(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func filterEq(proto, value reflect.Value) string {
	hint := ""

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgEqStrLen, proto.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		if value.IsZero() {
			value = reflect.ValueOf(0)
		} else {
			value = reflect.ValueOf(value.Len())
		}

		hint = fmt.Sprintf(MsgEqSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgEq, proto.Interface())
	}

	if !IsEqual(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func IsMatch(reg, value any) (flag bool) {
	if reflect.ValueOf(reg).Kind() == reflect.String &&
		reflect.ValueOf(value).Kind() == reflect.String {
		flag, _ = regexp.MatchString(reg.(string), value.(string))
	}

	return
}

func IsEachMatches(reg, value any) bool {
	refReg := reflect.ValueOf(reg)
	refVal := reflect.ValueOf(value)

	if refReg.Kind() == reflect.Invalid || refVal.Kind() == reflect.Invalid {
		return false
	}

	isValid := refReg.Kind() == reflect.String &&
		strings.HasSuffix(refVal.Type().String(), "string")

	switch refVal.Kind() {

	case reflect.Array, reflect.Slice:
		for n := 0; n < refVal.Len(); n++ {
			isValid = isValid && IsMatch(reg, refVal.Index(n).String())
		}
		break

	case reflect.Map:
		iter := refVal.MapRange()

		for iter.Next() {
			// k := iter.Key()
			// v := iter.Value()
			isValid = isValid && IsMatch(reg, iter.Value().String())
		}

		break
	}

	return isValid
}

func IsYearEqual(proto, value any) bool {
	refVal := reflect.ValueOf(value)

	if refVal.Kind() != reflect.Invalid {
		if refVal.Type().String() == "time.Time" {

			return IsEqual(proto, value.(time.Time).Year())

		}
	}

	return false
}

// https://go.dev/ref/spec#Numeric_types
func IsMin(proto, value any) bool {
	types := reflect.ValueOf(value).Kind().String() + ":" + reflect.ValueOf(proto).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return value.(int8) >= proto.(int8)

	case "int8:int16":
		return int16(value.(int8)) >= proto.(int16)

	case "int8:int32", "int8:rune":
		return int32(value.(int8)) >= proto.(int32)

	case "int8:int64":
		return int64(value.(int8)) >= proto.(int64)

	case "int8:int":
		return int(value.(int8)) >= proto.(int)

	case "int16:int8":
		return value.(int16) >= int16(proto.(int8))

	case "int16:int16":
		return value.(int16) >= proto.(int16)

	case "int16:int32", "int16:rune":
		return int32(value.(int16)) >= proto.(int32)

	case "int16:int64":
		return int64(value.(int16)) >= proto.(int64)

	case "int16:int":
		return int(value.(int16)) >= proto.(int)

	case "int32:int8", "rune:int8":
		return value.(int32) >= int32(proto.(int8))

	case "int32:int16", "rune:int16":
		return value.(int32) >= int32(proto.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return value.(int32) >= proto.(int32)

	case "int32:int64", "rune:int64":
		return int64(value.(int32)) >= proto.(int64)

	case "int32:int", "rune:int":
		return int(value.(int32)) >= proto.(int)

	case "int64:int8":
		return value.(int64) >= int64(proto.(int8))

	case "int64:int16":
		return value.(int64) >= int64(proto.(int16))

	case "int64:int32", "int64:rune":
		return value.(int64) >= int64(proto.(int32))

	case "int64:int64":
		return value.(int64) >= proto.(int64)

	case "int64:int":
		return value.(int64) >= int64(proto.(int))

	case "int:int8":
		return value.(int) >= int(proto.(int8))

	case "int:int16":
		return value.(int) >= int(proto.(int16))

	case "int:int32", "int:rune":
		return value.(int) >= int(proto.(int32))

	case "int:int64":
		return int64(value.(int)) >= proto.(int64)

	case "int:int":
		return value.(int) >= proto.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if value.(int8) >= 0 {
			return uint8(value.(int8)) >= proto.(uint8)
		}
		return false

	case "int8:uint16":
		if value.(int8) >= 0 {
			return uint16(value.(int8)) >= proto.(uint16)
		}
		return false

	case "int8:uint32":
		if value.(int8) >= 0 {
			return uint32(value.(int8)) >= proto.(uint32)
		}
		return false

	case "int8:uint64":
		if value.(int8) >= 0 {
			return uint64(value.(int8)) >= proto.(uint64)
		}
		return false

	case "int8:uint":
		if value.(int8) >= 0 {
			return uint(value.(int8)) >= proto.(uint)
		}
		return false

	// ...

	case "int16:uint8", "int16:byte":
		return value.(int16) >= int16(proto.(uint8))

	case "int16:uint16":
		if value.(int16) >= 0 {
			return uint16(value.(int16)) >= proto.(uint16)
		}
		return false

	case "int16:uint32":
		if value.(int16) >= 0 {
			return uint32(value.(int16)) >= proto.(uint32)
		}
		return false

	case "int16:uint64":
		if value.(int16) >= 0 {
			return uint64(value.(int16)) >= proto.(uint64)
		}
		return false

	case "int16:uint":
		if value.(int16) >= 0 {
			return uint(value.(int16)) >= proto.(uint)
		}
		return false

	// ...

	case "int32:uint8", "int32:byte":
		return value.(int32) >= int32(proto.(uint8))

	case "int32:uint16":
		return value.(int32) >= int32(proto.(uint16))

	case "int32:uint32":
		if value.(int32) >= 0 {
			return uint32(value.(int32)) >= proto.(uint32)
		}
		return false

	case "int32:uint64":
		if value.(int32) >= 0 {
			return uint64(value.(int32)) >= proto.(uint64)
		}
		return false

	case "int32:uint":
		if value.(int32) >= 0 {
			return uint(value.(int32)) >= proto.(uint)
		}
		return false

	// ...

	case "int64:uint8", "int64:byte":
		return value.(int64) >= int64(proto.(uint8))

	case "int64:uint16":
		return value.(int64) >= int64(proto.(uint16))

	case "int64:uint32":
		return value.(int64) >= int64(proto.(uint32))

	case "int64:uint64":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) >= uint64(proto.(uint64))
		}
		return false

	case "int64:uint":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) >= uint64(proto.(uint))
		}
		return false

	// ...

	case "int:uint8", "int:byte":
		return value.(int) >= int(proto.(uint8))

	case "int:uint16":
		return value.(int) >= int(proto.(uint16))

	case "int:uint32":
		if value.(int) >= 0 {
			return uint64(value.(int)) >= uint64(proto.(uint32))
		}
		return false

	case "int:uint64":
		if value.(int) >= 0 {
			return uint64(value.(int)) >= proto.(uint64)
		}
		return false

	case "int:uint":
		if value.(int) >= 0 {
			return uint(value.(int)) >= proto.(uint)
		}
		return false

	// ...

	case "rune:uint8", "rune:byte":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) >= uint32(proto.(uint8))
		}
		return false

	case "rune:uint16":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) >= uint32(proto.(uint16))
		}
		return false

	case "rune:uint32":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) >= proto.(uint32)
		}
		return false

	case "rune:uint64":
		if value.(rune) >= 0 {
			return uint64(value.(rune)) >= proto.(uint64)
		}
		return false

	case "rune:uint":
		if value.(rune) >= 0 {
			return uint(value.(rune)) >= proto.(uint)
		}
		return false

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return value.(uint8) >= proto.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(value.(uint8)) >= proto.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(value.(uint8)) >= proto.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(value.(uint8)) >= proto.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(value.(uint8)) >= proto.(uint)

	case "uint16:uint8", "uint16:byte":
		return value.(uint16) >= uint16(proto.(uint8))

	case "uint16:uint16":
		return value.(uint16) >= proto.(uint16)
	case "uint16:uint32":
		return uint32(value.(uint16)) >= proto.(uint32)
	case "uint16:uint64":
		return uint64(value.(uint16)) >= proto.(uint64)
	case "uint16:uint":
		return uint(value.(uint16)) >= proto.(uint)

	case "uint32:uint8", "uint32:byte":
		return value.(uint32) >= uint32(proto.(uint8))
	case "uint32:uint16":
		return value.(uint32) >= uint32(proto.(uint16))
	case "uint32:uint32":
		return value.(uint32) >= proto.(uint32)
	case "uint32:uint64":
		return uint64(value.(uint32)) >= proto.(uint64)
	case "uint32:uint":
		return uint(value.(uint32)) >= proto.(uint)

	case "uint64:uint8", "uint64:byte":
		return value.(uint64) >= uint64(proto.(uint8))
	case "uint64:uint16":
		return value.(uint64) >= uint64(proto.(uint16))
	case "uint64:uint32":
		return value.(uint64) >= uint64(proto.(uint32))
	case "uint64:uint64":
		return value.(uint64) >= proto.(uint64)
	case "uint64:uint":
		return value.(uint64) >= uint64(proto.(uint))

	case "uint:uint8", "uint:byte":
		return value.(uint) >= uint(proto.(uint8))
	case "uint:uint16":
		return value.(uint) >= uint(proto.(uint16))
	case "uint:uint32":
		return value.(uint) >= uint(proto.(uint32))
	case "uint:uint64":
		return uint64(value.(uint)) >= proto.(uint64)
	case "uint:uint":
		return value.(uint) >= proto.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if proto.(int8) >= 0 {
			return value.(uint8) >= uint8(proto.(int8))
		}
		return true

	case "uint8:int16", "byte:int16":
		return int16(value.(uint8)) >= proto.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(value.(uint8)) >= proto.(int32)

	case "uint8:int64", "byte:int64":
		return int64(value.(uint8)) >= proto.(int64)

	case "uint8:int", "byte:int":
		return int(value.(uint8)) >= proto.(int)

	// ...

	case "uint16:int8":
		if proto.(int8) >= 0 {
			return value.(uint16) >= uint16(proto.(int8))
		}
		return true

	case "uint16:int16":
		if proto.(int16) >= 0 {
			return value.(uint16) >= uint16(proto.(int16))
		}
		return true

	case "uint16:int32", "uint16:rune":
		return int32(value.(uint16)) >= proto.(int32)

	case "uint16:int64":
		return int64(value.(uint16)) >= proto.(int64)

	case "uint16:int":
		return int(value.(uint16)) >= proto.(int)

	// ...

	case "uint32:int8":
		if proto.(int8) >= 0 {
			return value.(uint32) >= uint32(proto.(int8))
		}
		return true

	case "uint32:int16":
		if proto.(int16) >= 0 {
			return value.(uint32) >= uint32(proto.(int16))
		}
		return true

	case "uint32:int32", "uint32:rune":
		if proto.(int32) >= 0 {
			return value.(uint32) >= uint32(proto.(int32))
		}
		return true

	case "uint32:int64":
		return int64(value.(uint32)) >= proto.(int64)

	case "uint32:int":
		return int64(value.(uint32)) >= int64(proto.(int))

	// ...

	case "uint64:int8":
		if proto.(int8) >= 0 {
			return value.(uint64) >= uint64(proto.(int8))
		}
		return true

	case "uint64:int16":
		if proto.(int16) >= 0 {
			return value.(uint64) >= uint64(proto.(int16))
		}
		return true

	case "uint64:int32", "uint64:rune":
		if proto.(int32) >= 0 {
			return value.(uint64) >= uint64(proto.(int32))
		}
		return true

	case "uint64:int64":
		if proto.(int64) >= 0 {
			return value.(uint64) >= uint64(proto.(int64))
		}
		return true

	case "uint64:int":
		if proto.(int) >= 0 {
			return value.(uint64) >= uint64(proto.(int))
		}
		return true

	// ...

	case "uint:int8":
		if proto.(int8) >= 0 {
			return value.(uint) >= uint(proto.(int8))
		}
		return true

	case "uint:int16":
		if proto.(int16) >= 0 {
			return value.(uint) >= uint(proto.(int16))
		}
		return true

	case "uint:int32":
		if proto.(int32) >= 0 {
			return value.(uint) >= uint(proto.(int32))
		}
		return true

	case "uint:int64":
		if proto.(int64) >= 0 {
			return uint64(value.(uint)) >= uint64(proto.(int64))
		}
		return true

	case "uint:int":
		if proto.(int) >= 0 {
			return value.(uint) >= uint(proto.(int))
		}
		return true

	}

	return false
}

func IsMax(proto, value any) bool {
	types := reflect.ValueOf(value).Kind().String() + ":" + reflect.ValueOf(proto).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return value.(int8) <= proto.(int8)

	case "int8:int16":
		return int16(value.(int8)) <= proto.(int16)

	case "int8:int32", "int8:rune":
		return int32(value.(int8)) <= proto.(int32)

	case "int8:int64":
		return int64(value.(int8)) <= proto.(int64)

	case "int8:int":
		return int(value.(int8)) <= proto.(int)

	case "int16:int8":
		return value.(int16) <= int16(proto.(int8))

	case "int16:int16":
		return value.(int16) <= proto.(int16)

	case "int16:int32", "int16:rune":
		return int32(value.(int16)) <= proto.(int32)

	case "int16:int64":
		return int64(value.(int16)) <= proto.(int64)

	case "int16:int":
		return int(value.(int16)) <= proto.(int)

	case "int32:int8", "rune:int8":
		return value.(int32) <= int32(proto.(int8))

	case "int32:int16", "rune:int16":
		return value.(int32) <= int32(proto.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return value.(int32) <= proto.(int32)

	case "int32:int64", "rune:int64":
		return int64(value.(int32)) <= proto.(int64)

	case "int32:int", "rune:int":
		return int(value.(int32)) <= proto.(int)

	case "int64:int8":
		return value.(int64) <= int64(proto.(int8))

	case "int64:int16":
		return value.(int64) <= int64(proto.(int16))

	case "int64:int32", "int64:rune":
		return value.(int64) <= int64(proto.(int32))

	case "int64:int64":
		return value.(int64) <= proto.(int64)

	case "int64:int":
		return value.(int64) <= int64(proto.(int))

	case "int:int8":
		return value.(int) <= int(proto.(int8))

	case "int:int16":
		return value.(int) <= int(proto.(int16))

	case "int:int32", "int:rune":
		return value.(int) <= int(proto.(int32))

	case "int:int64":
		return int64(value.(int)) <= proto.(int64)

	case "int:int":
		return value.(int) <= proto.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if value.(int8) >= 0 {
			return uint8(value.(int8)) <= proto.(uint8)
		}
		return true

	case "int8:uint16":
		if value.(int8) >= 0 {
			return uint16(value.(int8)) <= proto.(uint16)
		}
		return true

	case "int8:uint32":
		if value.(int8) >= 0 {
			return uint32(value.(int8)) <= proto.(uint32)
		}
		return true

	case "int8:uint64":
		if value.(int8) >= 0 {
			return uint64(value.(int8)) <= proto.(uint64)
		}
		return true

	case "int8:uint":
		if value.(int8) >= 0 {
			return uint(value.(int8)) <= proto.(uint)
		}
		return true

	// ...

	case "int16:uint8", "int16:byte":
		return value.(int16) <= int16(proto.(uint8))

	case "int16:uint16":
		if value.(int16) >= 0 {
			return uint16(value.(int16)) <= proto.(uint16)
		}
		return true

	case "int16:uint32":
		if value.(int16) >= 0 {
			return uint32(value.(int16)) <= proto.(uint32)
		}
		return true

	case "int16:uint64":
		if value.(int16) >= 0 {
			return uint64(value.(int16)) <= proto.(uint64)
		}
		return true

	case "int16:uint":
		if value.(int16) >= 0 {
			return uint(value.(int16)) <= proto.(uint)
		}
		return true

	// ...

	case "int32:uint8", "int32:byte":
		return value.(int32) <= int32(proto.(uint8))

	case "int32:uint16":
		return value.(int32) <= int32(proto.(uint16))

	case "int32:uint32":
		if value.(int32) >= 0 {
			return uint32(value.(int32)) <= proto.(uint32)
		}
		return true

	case "int32:uint64":
		if value.(int32) >= 0 {
			return uint64(value.(int32)) <= proto.(uint64)
		}
		return true

	case "int32:uint":
		if value.(int32) >= 0 {
			return uint(value.(int32)) <= proto.(uint)
		}
		return true

	// ...

	case "int64:uint8", "int64:byte":
		return value.(int64) <= int64(proto.(uint8))

	case "int64:uint16":
		return value.(int64) <= int64(proto.(uint16))

	case "int64:uint32":
		return value.(int64) <= int64(proto.(uint32))

	case "int64:uint64":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) <= uint64(proto.(uint64))
		}
		return true

	case "int64:uint":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) <= uint64(proto.(uint))
		}
		return true

	// ...

	case "int:uint8", "int:byte":
		return value.(int) <= int(proto.(uint8))

	case "int:uint16":
		return value.(int) <= int(proto.(uint16))

	case "int:uint32":
		if value.(int) >= 0 {
			return uint(value.(int)) <= uint(proto.(uint32))
		}
		return true

	case "int:uint64":
		if value.(int) >= 0 {
			return uint64(value.(int)) <= proto.(uint64)
		}
		return true

	case "int:uint":
		if value.(int) >= 0 {
			return uint(value.(int)) <= proto.(uint)
		}
		return true

	// ...

	case "rune:uint8", "rune:byte":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) <= uint32(proto.(uint8))
		}
		return false

	case "rune:uint16":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) <= uint32(proto.(uint16))
		}
		return false

	case "rune:uint32":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) <= proto.(uint32)
		}
		return false

	case "rune:uint64":
		if value.(rune) >= 0 {
			return uint64(value.(rune)) <= proto.(uint64)
		}
		return false

	case "rune:uint":
		if value.(rune) >= 0 {
			return uint(value.(rune)) <= proto.(uint)
		}
		return false

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return value.(uint8) <= proto.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(value.(uint8)) <= proto.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(value.(uint8)) <= proto.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(value.(uint8)) <= proto.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(value.(uint8)) <= proto.(uint)

	case "uint16:uint8", "uint16:byte":
		return value.(uint16) <= uint16(proto.(uint8))

	case "uint16:uint16":
		return value.(uint16) <= proto.(uint16)
	case "uint16:uint32":
		return uint32(value.(uint16)) <= proto.(uint32)
	case "uint16:uint64":
		return uint64(value.(uint16)) <= proto.(uint64)
	case "uint16:uint":
		return uint(value.(uint16)) <= proto.(uint)

	case "uint32:uint8", "uint32:byte":
		return value.(uint32) <= uint32(proto.(uint8))
	case "uint32:uint16":
		return value.(uint32) <= uint32(proto.(uint16))
	case "uint32:uint32":
		return value.(uint32) <= proto.(uint32)
	case "uint32:uint64":
		return uint64(value.(uint32)) <= proto.(uint64)
	case "uint32:uint":
		return uint(value.(uint32)) <= proto.(uint)

	case "uint64:uint8", "uint64:byte":
		return value.(uint64) <= uint64(proto.(uint8))
	case "uint64:uint16":
		return value.(uint64) <= uint64(proto.(uint16))
	case "uint64:uint32":
		return value.(uint64) <= uint64(proto.(uint32))
	case "uint64:uint64":
		return value.(uint64) <= proto.(uint64)
	case "uint64:uint":
		return value.(uint64) <= uint64(proto.(uint))

	case "uint:uint8", "uint:byte":
		return value.(uint) <= uint(proto.(uint8))
	case "uint:uint16":
		return value.(uint) <= uint(proto.(uint16))
	case "uint:uint32":
		return value.(uint) <= uint(proto.(uint32))
	case "uint:uint64":
		return uint64(value.(uint)) <= proto.(uint64)
	case "uint:uint":
		return value.(uint) <= proto.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if proto.(int8) >= 0 {
			return value.(uint8) <= uint8(proto.(int8))
		}
		return false

	case "uint8:int16", "byte:int16":
		return int16(value.(uint8)) <= proto.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(value.(uint8)) <= proto.(int32)

	case "uint8:int64", "byte:int64":
		return int64(value.(uint8)) <= proto.(int64)

	case "uint8:int", "byte:int":
		return int(value.(uint8)) <= proto.(int)

	// ...

	case "uint16:int8":
		if proto.(int8) >= 0 {
			return value.(uint16) <= uint16(proto.(int8))
		}
		return false

	case "uint16:int16":
		if proto.(int16) >= 0 {
			return value.(uint16) <= uint16(proto.(int16))
		}
		return false

	case "uint16:int32", "uint16:rune":
		return int32(value.(uint16)) <= proto.(int32)

	case "uint16:int64":
		return int64(value.(uint16)) <= proto.(int64)

	case "uint16:int":
		return int(value.(uint16)) <= proto.(int)

	// ...

	case "uint32:int8":
		if proto.(int8) >= 0 {
			return value.(uint32) <= uint32(proto.(int8))
		}
		return false

	case "uint32:int16":
		if proto.(int16) >= 0 {
			return value.(uint32) <= uint32(proto.(int16))
		}
		return false

	case "uint32:int32", "uint32:rune":
		if proto.(int32) >= 0 {
			return value.(uint32) <= uint32(proto.(int32))
		}
		return false

	case "uint32:int64":
		return int64(value.(uint32)) <= proto.(int64)

	case "uint32:int":
		return int64(value.(uint32)) <= int64(proto.(int))

	// ...

	case "uint64:int8":
		if proto.(int8) >= 0 {
			return value.(uint64) <= uint64(proto.(int8))
		}
		return false

	case "uint64:int16":
		if proto.(int16) >= 0 {
			return value.(uint64) <= uint64(proto.(int16))
		}
		return false

	case "uint64:int32", "uint64:rune":
		if proto.(int32) >= 0 {
			return value.(uint64) <= uint64(proto.(int32))
		}
		return false

	case "uint64:int64":
		if proto.(int64) >= 0 {
			return value.(uint64) <= uint64(proto.(int64))
		}
		return false

	case "uint64:int":
		if proto.(int) >= 0 {
			return value.(uint64) <= uint64(proto.(int))
		}
		return false

	// ...

	case "uint:int8":
		if proto.(int8) >= 0 {
			return value.(uint) <= uint(proto.(int8))
		}
		return false

	case "uint:int16":
		if proto.(int16) >= 0 {
			return value.(uint) <= uint(proto.(int16))
		}
		return false

	case "uint:int32", "uint:rune":
		if proto.(int32) >= 0 {
			return value.(uint) <= uint(proto.(int32))
		}
		return false

	case "uint:int64":
		if proto.(int64) >= 0 {
			return uint64(value.(uint)) <= uint64(proto.(int64))
		}
		return false

	case "uint:int":
		if proto.(int) >= 0 {
			return value.(uint) <= uint(proto.(int))
		}
		return false

	}

	return false
}

func IsEqual(proto, value any) bool {
	types := reflect.ValueOf(value).Kind().String() + ":" + reflect.ValueOf(proto).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return value.(int8) == proto.(int8)

	case "int8:int16":
		return int16(value.(int8)) == proto.(int16)

	case "int8:int32", "int8:rune":
		return int32(value.(int8)) == proto.(int32)

	case "int8:int64":
		return int64(value.(int8)) == proto.(int64)

	case "int8:int":
		return int(value.(int8)) == proto.(int)

	case "int16:int8":
		return value.(int16) == int16(proto.(int8))

	case "int16:int16":
		return value.(int16) == proto.(int16)

	case "int16:int32", "int16:rune":
		return int32(value.(int16)) == proto.(int32)

	case "int16:int64":
		return int64(value.(int16)) == proto.(int64)

	case "int16:int":
		return int(value.(int16)) == proto.(int)

	case "int32:int8", "rune:int8":
		return value.(int32) == int32(proto.(int8))

	case "int32:int16", "rune:int16":
		return value.(int32) == int32(proto.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return value.(int32) == proto.(int32)

	case "int32:int64", "rune:int64":
		return int64(value.(int32)) == proto.(int64)

	case "int32:int", "rune:int":
		return int(value.(int32)) == proto.(int)

	case "int64:int8":
		return value.(int64) == int64(proto.(int8))

	case "int64:int16":
		return value.(int64) == int64(proto.(int16))

	case "int64:int32", "int64:rune":
		return value.(int64) == int64(proto.(int32))

	case "int64:int64":
		return value.(int64) == proto.(int64)

	case "int64:int":
		return value.(int64) == int64(proto.(int))

	case "int:int8":
		return value.(int) == int(proto.(int8))

	case "int:int16":
		return value.(int) == int(proto.(int16))

	case "int:int32", "int:rune":
		return value.(int) == int(proto.(int32))

	case "int:int64":
		return int64(value.(int)) == proto.(int64)

	case "int:int":
		return value.(int) == proto.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if value.(int8) >= 0 {
			return uint8(value.(int8)) == proto.(uint8)
		}

	case "int8:uint16":
		if value.(int8) >= 0 {
			return uint16(value.(int8)) == proto.(uint16)
		}

	case "int8:uint32":
		if value.(int8) >= 0 {
			return uint32(value.(int8)) == proto.(uint32)
		}

	case "int8:uint64":
		if value.(int8) >= 0 {
			return uint64(value.(int8)) == proto.(uint64)
		}

	case "int8:uint":
		if value.(int8) >= 0 {
			return uint(value.(int8)) == proto.(uint)
		}

	// ...

	case "int16:uint8", "int16:byte":
		return value.(int16) == int16(proto.(uint8))

	case "int16:uint16":
		if value.(int16) >= 0 {
			return uint16(value.(int16)) == proto.(uint16)
		}

	case "int16:uint32":
		if value.(int16) >= 0 {
			return uint32(value.(int16)) == proto.(uint32)
		}

	case "int16:uint64":
		if value.(int16) >= 0 {
			return uint64(value.(int16)) == proto.(uint64)
		}

	case "int16:uint":
		if value.(int16) >= 0 {
			return uint(value.(int16)) == proto.(uint)
		}

	// ...

	case "int32:uint8", "int32:byte":
		return value.(int32) == int32(proto.(uint8))

	case "int32:uint16":
		return value.(int32) == int32(proto.(uint16))

	case "int32:uint32":
		if value.(int32) >= 0 {
			return uint32(value.(int32)) == proto.(uint32)
		}

	case "int32:uint64":
		if value.(int32) >= 0 {
			return uint64(value.(int32)) == proto.(uint64)
		}

	case "int32:uint":
		if value.(int32) >= 0 {
			return uint(value.(int32)) == proto.(uint)
		}

	// ...

	case "int64:uint8", "int64:byte":
		return value.(int64) == int64(proto.(uint8))

	case "int64:uint16":
		return value.(int64) == int64(proto.(uint16))

	case "int64:uint32":
		return value.(int64) == int64(proto.(uint32))

	case "int64:uint64":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) == uint64(proto.(uint64))
		}

	case "int64:uint":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) == uint64(proto.(uint))
		}

	// ...

	case "int:uint8", "int:byte":
		return value.(int) == int(proto.(uint8))

	case "int:uint16":
		return value.(int) == int(proto.(uint16))

	case "int:uint32":
		if value.(int) >= 0 {
			return uint(value.(int)) == uint(proto.(uint32))
		}

	case "int:uint64":
		if value.(int) >= 0 {
			return uint64(value.(int)) == proto.(uint64)
		}

	case "int:uint":
		if value.(int) >= 0 {
			return uint(value.(int)) == proto.(uint)
		}

	// ...

	case "rune:uint8", "rune:byte":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) == uint32(proto.(uint8))
		}

	case "rune:uint16":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) == uint32(proto.(uint16))
		}

	case "rune:uint32":
		if value.(rune) >= 0 {
			return uint32(value.(rune)) == proto.(uint32)
		}

	case "rune:uint64":
		if value.(rune) >= 0 {
			return uint64(value.(rune)) == proto.(uint64)
		}

	case "rune:uint":
		if value.(rune) >= 0 {
			return uint(value.(rune)) == proto.(uint)
		}

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return value.(uint8) == proto.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(value.(uint8)) == proto.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(value.(uint8)) == proto.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(value.(uint8)) == proto.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(value.(uint8)) == proto.(uint)

	case "uint16:uint8", "uint16:byte":
		return value.(uint16) == uint16(proto.(uint8))

	case "uint16:uint16":
		return value.(uint16) == proto.(uint16)
	case "uint16:uint32":
		return uint32(value.(uint16)) == proto.(uint32)
	case "uint16:uint64":
		return uint64(value.(uint16)) == proto.(uint64)
	case "uint16:uint":
		return uint(value.(uint16)) == proto.(uint)

	case "uint32:uint8", "uint32:byte":
		return value.(uint32) == uint32(proto.(uint8))
	case "uint32:uint16":
		return value.(uint32) == uint32(proto.(uint16))
	case "uint32:uint32":
		return value.(uint32) == proto.(uint32)
	case "uint32:uint64":
		return uint64(value.(uint32)) == proto.(uint64)
	case "uint32:uint":
		return uint(value.(uint32)) == proto.(uint)

	case "uint64:uint8", "uint64:byte":
		return value.(uint64) == uint64(proto.(uint8))
	case "uint64:uint16":
		return value.(uint64) == uint64(proto.(uint16))
	case "uint64:uint32":
		return value.(uint64) == uint64(proto.(uint32))
	case "uint64:uint64":
		return value.(uint64) == proto.(uint64)
	case "uint64:uint":
		return value.(uint64) == uint64(proto.(uint))

	case "uint:uint8", "uint:byte":
		return value.(uint) == uint(proto.(uint8))
	case "uint:uint16":
		return value.(uint) == uint(proto.(uint16))
	case "uint:uint32":
		return value.(uint) == uint(proto.(uint32))
	case "uint:uint64":
		return uint64(value.(uint)) == proto.(uint64)
	case "uint:uint":
		return value.(uint) == proto.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if proto.(int8) >= 0 {
			return value.(uint8) == uint8(proto.(int8))
		}

	case "uint8:int16", "byte:int16":
		return int16(value.(uint8)) == proto.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(value.(uint8)) == proto.(int32)

	case "uint8:int64", "byte:int64":
		return int64(value.(uint8)) == proto.(int64)

	case "uint8:int", "byte:int":
		return int(value.(uint8)) == proto.(int)

	// ...

	case "uint16:int8":
		if proto.(int8) >= 0 {
			return value.(uint16) == uint16(proto.(int8))
		}

	case "uint16:int16":
		if proto.(int16) >= 0 {
			return value.(uint16) == uint16(proto.(int16))
		}

	case "uint16:int32", "uint16:rune":
		return int32(value.(uint16)) == proto.(int32)

	case "uint16:int64":
		return int64(value.(uint16)) == proto.(int64)

	case "uint16:int":
		return int(value.(uint16)) == proto.(int)

	// ...

	case "uint32:int8":
		if proto.(int8) >= 0 {
			return value.(uint32) == uint32(proto.(int8))
		}

	case "uint32:int16":
		if proto.(int16) >= 0 {
			return value.(uint32) == uint32(proto.(int16))
		}

	case "uint32:int32", "uint32:rune":
		if proto.(int32) >= 0 {
			return value.(uint32) == uint32(proto.(int32))
		}

	case "uint32:int64":
		return int64(value.(uint32)) == proto.(int64)

	case "uint32:int":
		return int64(value.(uint32)) == int64(proto.(int))

	// ...

	case "uint64:int8":
		if proto.(int8) >= 0 {
			return value.(uint64) == uint64(proto.(int8))
		}

	case "uint64:int16":
		if proto.(int16) >= 0 {
			return value.(uint64) == uint64(proto.(int16))
		}

	case "uint64:int32", "uint64:rune":
		if proto.(int32) >= 0 {
			return value.(uint64) == uint64(proto.(int32))
		}

	case "uint64:int64":
		if proto.(int64) >= 0 {
			return value.(uint64) == uint64(proto.(int64))
		}

	case "uint64:int":
		if proto.(int) >= 0 {
			return value.(uint64) == uint64(proto.(int))
		}

	// ...

	case "uint:int8":
		if proto.(int8) >= 0 {
			return value.(uint) == uint(proto.(int8))
		}

	case "uint:int16":
		if proto.(int16) >= 0 {
			return value.(uint) == uint(proto.(int16))
		}

	case "uint:int32", "uint:rune":
		if proto.(int32) >= 0 {
			return value.(uint) == uint(proto.(int32))
		}

	case "uint:int64":
		if proto.(int64) >= 0 {
			return uint64(value.(uint)) == uint64(proto.(int64))
		}

	case "uint:int":
		if proto.(int) >= 0 {
			return value.(uint) == uint(proto.(int))
		}

	}

	return false
}
