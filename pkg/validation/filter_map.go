package validation

type FilterMap map[string]map[string]any

func (filter FilterMap) IsValid(anyStruct any) bool {
	flag, _ := CheckIsValid(filter, anyStruct, true)
	return flag
}

func (filter FilterMap) Validate(anyStruct any) (errList []string) {
	_, wrongFields := CheckIsValid(filter, anyStruct, false)
	return wrongFields
}
