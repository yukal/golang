package validation

type FilterMap map[string]map[string]any

func (filter FilterMap) IsValid(anyStruct any) bool {
	return IsValid(filter, anyStruct)
}

func (filter FilterMap) Validate(anyStruct any) (errList []string) {
	return Validate(filter, anyStruct)
}
