package validation

import (
	"reflect"
	"regexp"
)

func IsMatch(reg, value any) (flag bool) {
	if reflect.ValueOf(reg).Kind() == reflect.String &&
		reflect.ValueOf(value).Kind() == reflect.String {
		flag, _ = regexp.MatchString(reg.(string), value.(string))
	}

	return
}
