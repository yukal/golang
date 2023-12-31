package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"unicode/utf8"
)

const (
	NON_ZERO  = "NonZero"
	NON_EMPTY = "NonEmpty"

	MsgMinFields      = "must contain at least %d valid fields"
	MsgMaxFields      = "must contain up to %d valid fields"
	MsgMinStrLen      = "must contain at least %d characters"
	MsgMaxStrLen      = "must contain up to %d characters"
	MsgEqStrLen       = "must contain exactly %d characters"
	MsgRangeStrLen    = "must contain %d..%d characters"
	MsgMinSetLen      = "must contain at least %d items"
	MsgMaxSetLen      = "must contain up to %d items"
	MsgEqSetLen       = "must contain exactly %d items"
	MsgRangeSetLen    = "must contain %d..%d items"
	MsgMin            = "must be at least %d"
	MsgMax            = "must be up to %d"
	MsgEq             = "must be exactly %d"
	MsgRange          = "must be in the range %d..%d"
	MsgNotValid       = "is not valid"
	MsgEmpty          = "is empty"
	MsgInvalidValue   = "has invalid value"
	MsgInvalidRule    = "has invalid rule"
	MsgInvalidBodyVal = "invalid body value"
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
	return len(filter.Validate(data)) == 0
}

// data should be a type of struct{ ... }
func (filter Filter) Validate(data any) []string {
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

			if action == "minFields" {
				value = reflect.ValueOf(successFields)
			}

			if hint := compare(action, proto, value); hint != "" {
				hints = append(hints, MsgInvalidBodyVal)
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

			if hint := checkField(item, value); hint != "" {
				return hint
			}
		}

	case "validation.Range":
		return compare("range", rules, value)

	case "validation.Rule":
		action := rules.Index(0).Elem().String()
		proto := rules.Index(1).Elem()

		return compare(action, proto, value)

	case "string":
		action := rules.String()
		proto := reflect.ValueOf(nil)

		return compare(action, proto, value)
	}

	return ""
}

func compare(action string, proto, value reflect.Value) string {
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
		return MsgInvalidRule
	}

	switch action {
	case "minFields":
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
		if !IsMatch(proto, value) {
			return MsgNotValid
		}

	case "eachMatch":
		if (proto.Kind() != reflect.String) || (proto.Len() == 0) {
			return MsgInvalidRule
		}

		if !IsEachMatch(proto.String(), value) {
			return MsgNotValid
		}

	case "eachMin":
		return IsEachMin(proto, value)

	case "eachMax":
		return IsEachMax(proto, value)

	case "eachEq":
		return IsEachEq(proto, value)

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

	if valMin.IsNil() || valMax.IsNil() {
		return MsgInvalidRule
	}

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
		value = reflect.ValueOf(value.Len())
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
		value = reflect.ValueOf(value.Len())
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
		value = reflect.ValueOf(value.Len())
		hint = fmt.Sprintf(MsgEqSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgEq, proto.Interface())
	}

	if !IsEqual(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func IsMatch(reg, value reflect.Value) (flag bool) {
	if reg.Kind() == reflect.String && value.Kind() == reflect.String {
		flag, _ = regexp.MatchString(reg.String(), value.String())
	}

	return
}

func IsEachMatch(reg string, value reflect.Value) bool {
	isValid := false

	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		if isValid = value.Type().Elem().Kind() == reflect.String; !isValid {
			return false
		}

		for n := 0; n < value.Len(); n++ {
			matched, _ := regexp.MatchString(reg, value.Index(n).String())
			isValid = isValid && matched
		}

	case reflect.Map:
		if isValid = value.Type().Elem().Kind() == reflect.String; !isValid {
			return false
		}

		iter := value.MapRange()

		for iter.Next() {
			// k := iter.Key()
			// v := iter.Value()

			matched, _ := regexp.MatchString(reg, iter.Value().String())
			isValid = isValid && matched
		}
	}

	return isValid
}

func IsEachEq(proto, value reflect.Value) string {
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for n := 0; n < value.Len(); n++ {
			if hint := filterEq(proto, value.Index(n)); len(hint) > 0 {
				return fmt.Sprintf("item[%v] "+hint, n)
			}
		}

		return ""

	case reflect.Map:
		iter := value.MapRange()

		for iter.Next() {
			// k := iter.Key()
			// v := iter.Value()

			if hint := filterEq(proto, iter.Value()); len(hint) > 0 {
				return fmt.Sprintf("item[%v] "+hint, iter.Key())
			}
		}

		return ""
	}

	return MsgInvalidValue
}

func IsEachMax(proto, value reflect.Value) string {
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for n := 0; n < value.Len(); n++ {
			if hint := filterMax(proto, value.Index(n)); len(hint) > 0 {
				return fmt.Sprintf("item[%v] "+hint, n)
			}
		}

		return ""

	case reflect.Map:
		iter := value.MapRange()

		for iter.Next() {
			// k := iter.Key()
			// v := iter.Value()

			if hint := filterMax(proto, iter.Value()); len(hint) > 0 {
				return fmt.Sprintf("item[%v] "+hint, iter.Key())
			}
		}

		return ""
	}

	return MsgInvalidValue
}

func IsEachMin(proto, value reflect.Value) string {
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for n := 0; n < value.Len(); n++ {
			if hint := filterMin(proto, value.Index(n)); len(hint) > 0 {
				return fmt.Sprintf("item[%v] "+hint, n)
			}
		}

		return ""

	case reflect.Map:
		iter := value.MapRange()

		for iter.Next() {
			// k := iter.Key()
			// v := iter.Value()

			if hint := filterMin(proto, iter.Value()); len(hint) > 0 {
				return fmt.Sprintf("item[%v] "+hint, iter.Key())
			}
		}

		return ""
	}

	return MsgInvalidValue
}
