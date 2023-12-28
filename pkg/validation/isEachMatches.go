package validation

import (
	"reflect"
	"strings"
)

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
