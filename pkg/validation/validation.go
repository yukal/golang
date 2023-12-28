package validation

import (
	"fmt"
	"reflect"
	"strings"
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
	MsgInvalidValue    = "invalid value"
	MsgInvalidRangeVal = "invalid range value"
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

			if strings.Contains(action, "-fields") {
				value = reflect.ValueOf(successFields)
			}

			if hint := compare(action, proto, value); hint != "" {
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
				if hint := compare("range", rules, value); hint != "" {
					return hint
				}

			case "validation.Rule":
				action := item.Index(0).Elem().String()
				proto := item.Index(1).Elem()

				if hint := compare(action, proto, value); hint != "" {
					return hint
				}

			case "string":
				action := item.Elem().String()
				proto := reflect.ValueOf(nil)

				if hint := compare(action, proto, value); hint != "" {
					return hint
				}
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

	if valMin.IsZero() || valMax.IsZero() {
		return MsgInvalidRangeVal
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
