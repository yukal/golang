package validation

type FilterMap map[string]map[string]any

func (filter FilterMap) IsValid(anyStruct any) bool {
	flag, _ := checkIsValid(filter, anyStruct, true)
	return flag
}

func (filter FilterMap) Validate(anyStruct any) (errList []string) {
	_, wrongFields := checkIsValid(filter, anyStruct, false)
	return wrongFields
}
