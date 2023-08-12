package validation

type FilterMap map[string]map[string]any

func (filter FilterMap) IsValid(anyStruct any) bool {
	flag, _ := _checkIsValid(filter, anyStruct, true)
	return flag
}

func (filter FilterMap) Validate(anyStruct any) (errList []string) {
	_, wrongFields := _checkIsValid(filter, anyStruct, false)
	return wrongFields
}
