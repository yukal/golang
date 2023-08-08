package main

import (
	"reflect"
	"time"
)

func Compare(rule string, filterVal, comparableVal any) bool {
	switch rule {

	case "min":
		return validateMin(filterVal, comparableVal)

	case "max":
		return validateMax(comparableVal, filterVal)

	case "eq":
		return validateEq(comparableVal, filterVal)

	case "strLen":
		return validateStrLen(comparableVal, filterVal)

	case "minLen":
		return validateMinLen(comparableVal, filterVal)

	case "maxLen":
		return validateMaxLen(comparableVal, filterVal)

	case "year":
		return validateYear(comparableVal, filterVal)

	}

	return true
}

// https://go.dev/ref/spec#Numeric_types
// 
func validateMin(filterVal, val any) bool {
	types := reflect.TypeOf(val).Kind().String() + ":" + reflect.TypeOf(filterVal).Kind().String()

	switch types {

	// int

	case "int8:int8":     return val.(int8) >= filterVal.(int8)
	case "int8:int16":    return int16(val.(int8)) >= filterVal.(int16)
	case "int8:int32":    return int32(val.(int8)) >= filterVal.(int32)
	case "int8:int64":    return int64(val.(int8)) >= filterVal.(int64)
	case "int8:int":      return int(val.(int8)) >= filterVal.(int)

	case "int16:int8":    return val.(int16) >= int16(filterVal.(int8))
	case "int16:int16":   return val.(int16) >= filterVal.(int16)
	case "int16:int32":   return int32(val.(int16)) >= filterVal.(int32)
	case "int16:int64":   return int64(val.(int16)) >= filterVal.(int64)
	case "int16:int":     return int(val.(int16)) >= filterVal.(int)

	case "int32:int8":    return val.(int32) >= int32(filterVal.(int8))
	case "int32:int16":   return val.(int32) >= int32(filterVal.(int16))
	case "int32:int32":   return val.(int32) >= filterVal.(int32)
	case "int32:int64":   return int64(val.(int32)) >= filterVal.(int64)
	case "int32:int":     return int(val.(int32)) >= filterVal.(int)

	case "int64:int8":    return val.(int64) >= int64(filterVal.(int8))
	case "int64:int16":   return val.(int64) >= int64(filterVal.(int16))
	case "int64:int32":   return val.(int64) >= int64(filterVal.(int32))
	case "int64:int64":   return val.(int64) >= filterVal.(int64)
	case "int64:int":     return val.(int64) >= int64(filterVal.(int))

	case "int:int8":      return val.(int) >= int(filterVal.(int8))
	case "int:int16":     return val.(int) >= int(filterVal.(int16))
	case "int:int32":     return val.(int) >= int(filterVal.(int32))
	case "int:int64":     return int64(val.(int)) >= filterVal.(int64)
	case "int:int":       return val.(int) >= filterVal.(int)

	// uint

	case "uint8:uint8":   return val.(uint8) >= filterVal.(uint8)
	case "uint8:uint16":  return uint16(val.(uint8)) >= filterVal.(uint16)
	case "uint8:uint32":  return uint32(val.(uint8)) >= filterVal.(uint32)
	case "uint8:uint64":  return uint64(val.(uint8)) >= filterVal.(uint64)
	case "uint8:uint":    return uint(val.(uint8)) >= filterVal.(uint)

	case "uint16:uint8":  return val.(uint16) >= uint16(filterVal.(uint8))
	case "uint16:uint16": return val.(uint16) >= filterVal.(uint16)
	case "uint16:uint32": return uint32(val.(uint16)) >= filterVal.(uint32)
	case "uint16:uint64": return uint64(val.(uint16)) >= filterVal.(uint64)
	case "uint16:uint":   return uint(val.(uint16)) >= filterVal.(uint)

	case "uint32:uint8":  return val.(uint32) >= uint32(filterVal.(uint8))
	case "uint32:uint16": return val.(uint32) >= uint32(filterVal.(uint16))
	case "uint32:uint32": return val.(uint32) >= filterVal.(uint32)
	case "uint32:uint64": return uint64(val.(uint32)) >= filterVal.(uint64)
	case "uint32:uint":   return uint(val.(uint32)) >= filterVal.(uint)

	case "uint64:uint8":  return val.(uint64) >= uint64(filterVal.(uint8))
	case "uint64:uint16": return val.(uint64) >= uint64(filterVal.(uint16))
	case "uint64:uint32": return val.(uint64) >= uint64(filterVal.(uint32))
	case "uint64:uint64": return val.(uint64) >= filterVal.(uint64)
	case "uint64:uint":   return val.(uint64) >= uint64(filterVal.(uint))

	case "uint:uint8":    return val.(uint) >= uint(filterVal.(uint8))
	case "uint:uint16":   return val.(uint) >= uint(filterVal.(uint16))
	case "uint:uint32":   return val.(uint) >= uint(filterVal.(uint32))
	case "uint:uint64":   return uint64(val.(uint)) >= filterVal.(uint64)
	case "uint:uint":     return val.(uint) >= filterVal.(uint)

	// int & uint

	case "int8:uint8":
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

	case "int16:uint8":
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

	case "int32:uint8":
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

	case "int64:uint8":
		return val.(int64) >= int64(filterVal.(uint8))

	case "int64:uint16":
		return val.(int64) >= int64(filterVal.(uint16))

	case "int64:uint32":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) >= uint64(filterVal.(uint32))
		}
		return false

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

	case "int:uint8":
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

	// uint & int

	case "uint8:int8":
		if filterVal.(int8) >= 0 {
			return val.(uint8) >= uint8(filterVal.(int8))
		}
		return true

	case "uint8:int16": return int16(val.(uint8)) >= filterVal.(int16)
	case "uint8:int32": return int32(val.(uint8)) >= filterVal.(int32)
	case "uint8:int64": return int64(val.(uint8)) >= filterVal.(int64)
	case "uint8:int": return int(val.(uint8)) >= filterVal.(int)

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

	case "uint16:int32": return int32(val.(uint16)) >= filterVal.(int32)
	case "uint16:int64": return int64(val.(uint16)) >= filterVal.(int64)
	case "uint16:int": return int(val.(uint16)) >= filterVal.(int)

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

	case "uint32:int32":
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

	case "uint64:int32":
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

	// rune

	case "rune:int8":  return val.(rune) >= int32(filterVal.(int8))
	case "rune:int16": return val.(rune) >= int32(filterVal.(int16))
	case "rune:int64": return int64(val.(rune)) >= filterVal.(int64)
	case "rune:int":   return int(val.(rune)) >= filterVal.(int)
	case "int8:rune":  return int32(val.(int8)) >= filterVal.(rune)
	case "int16:rune": return int32(val.(int16)) >= filterVal.(rune)
	case "int64:rune": return val.(int64) >= int64(filterVal.(rune))
	case "int:rune":   return val.(int) >= int(filterVal.(rune))
	case "rune:rune",
		 "rune:int32",
		 "int32:rune": return val.(rune) >= filterVal.(int32)

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

	// ...

	case "uint8:rune":
		if filterVal.(rune) >= 0 {
			return uint32(val.(uint8)) >= uint32(filterVal.(rune))
		}
		return true

	case "uint16:rune":
		if filterVal.(rune) >= 0 {
			return uint32(val.(uint16)) >= uint32(filterVal.(rune))
		}
		return true

	case "uint32:rune":
		if filterVal.(rune) >= 0 {
			return uint32(val.(uint32)) >= uint32(filterVal.(rune))
		}
		return true

	case "uint64:rune":
		if filterVal.(rune) >= 0 {
			return uint64(val.(uint64)) >= uint64(filterVal.(rune))
		}
		return true

	case "uint:rune":
		if filterVal.(rune) >= 0 {
			return val.(uint) >= uint(filterVal.(rune))
		}
		return true

	// byte

	case "byte:uint16": return uint16(val.(byte)) >= filterVal.(uint16)
	case "byte:uint32": return uint32(val.(byte)) >= filterVal.(uint32)
	case "byte:uint64": return uint64(val.(byte)) >= filterVal.(uint64)
	case "byte:uint":   return uint(val.(byte)) >= filterVal.(uint)
	case "uint16:byte": return val.(uint16) >= uint16(filterVal.(byte))
	case "uint32:byte": return val.(uint32) >= uint32(filterVal.(byte))
	case "uint64:byte": return val.(uint64) >= uint64(filterVal.(byte))
	case "uint:byte":   return val.(uint) >= uint(filterVal.(byte))
	case "byte:byte",
		 "byte:uint8",
		 "uint8:byte":  return val.(byte) >= filterVal.(uint8)

	// ...

	case "byte:int8":
		if filterVal.(int8) >= 0 {
			return val.(byte) >= uint8(filterVal.(int8))
		}
		return true

	case "byte:int16":
		if filterVal.(int16) >= 0 {
			return uint16(val.(byte)) >= uint16(filterVal.(int16))
		}
		return true

	case "byte:int32":
		if filterVal.(int32) >= 0 {
			return uint32(val.(byte)) >= uint32(filterVal.(int32))
		}
		return true

	case "byte:int64":
		if filterVal.(int64) >= 0 {
			return uint64(val.(byte)) >= uint64(filterVal.(int64))
		}
		return true

	case "byte:int":
		if filterVal.(int) >= 0 {
			return uint(val.(byte)) >= uint(filterVal.(int))
		}
		return true

	// ...

	case "int8:byte":
		if val.(int8) >= 0 {
			return uint8(val.(int8)) >= filterVal.(byte)
		}
		return false

	case "int16:byte":
		if val.(int16) >= 0 {
			return uint16(val.(int16)) >= uint16(filterVal.(byte))
		}
		return false

	case "int32:byte":
		if val.(int32) >= 0 {
			return uint32(val.(int32)) >= uint32(filterVal.(byte))
		}
		return false

	case "int64:byte":
		if val.(int64) >= 0 {
			return uint64(val.(int64)) >= uint64(filterVal.(byte))
		}
		return false

	case "int:byte":
		if val.(int) >= 0 {
			return uint(val.(int)) >= uint(filterVal.(byte))
		}
		return false

	}

	return false
}

func validateMax(val any, filterVal any) bool {
	types := reflect.TypeOf(val).Kind().String() + ":" + reflect.TypeOf(filterVal).Kind().String()

	switch types {

	// int

	case "int64:int64":   return val.(int64) <= filterVal.(int64)
	case "int64:int32":   return val.(int64) <= int64(filterVal.(int32))
	case "int64:int16":   return val.(int64) <= int64(filterVal.(int16))
	case "int64:int8":    return val.(int64) <= int64(filterVal.(int8))
	case "int64:int":     return val.(int64) <= int64(filterVal.(int))

	case "int32:int64":   return int64(val.(int32)) <= filterVal.(int64)
	case "int32:int32":   return val.(int32) <= filterVal.(int32)
	case "int32:int16":   return val.(int32) <= int32(filterVal.(int16))
	case "int32:int8":    return val.(int32) <= int32(filterVal.(int8))
	case "int32:int":     return val.(int32) <= int32(filterVal.(int))

	case "int16:int64":   return int64(val.(int16)) <= filterVal.(int64)
	case "int16:int32":   return int32(val.(int16)) <= filterVal.(int32)
	case "int16:int16":   return val.(int16) <= filterVal.(int16)
	case "int16:int8":    return val.(int16) <= int16(filterVal.(int8))
	case "int16:int":     return val.(int16) <= int16(filterVal.(int))

	case "int8:int64":    return int64(val.(int8)) <= filterVal.(int64)
	case "int8:int32":    return int32(val.(int8)) <= filterVal.(int32)
	case "int8:int16":    return int16(val.(int8)) <= filterVal.(int16)
	case "int8:int8":     return val.(int8) <= filterVal.(int8)
	case "int8:int":      return val.(int8) <= int8(filterVal.(int))

	case "int:int64":     return int64(val.(int)) <= filterVal.(int64)
	case "int:int32":     return int32(val.(int)) <= filterVal.(int32)
	case "int:int16":     return int16(val.(int)) <= filterVal.(int16)
	case "int:int8":      return int8(val.(int)) <= filterVal.(int8)
	case "int:int":       return val.(int) <= filterVal.(int)

	// int & uint

	case "int64:uint64":  return val.(int64) <= int64(filterVal.(uint64))
	case "int64:uint32":  return val.(int64) <= int64(filterVal.(uint32))
	case "int64:uint16":  return val.(int64) <= int64(filterVal.(uint16))
	case "int64:uint8":   return val.(int64) <= int64(filterVal.(uint8))
	case "int64:uint":    return val.(int64) <= int64(filterVal.(uint))

	case "int32:uint64":  return int64(val.(int32)) <= int64(filterVal.(uint64))
	case "int32:uint32":  return val.(int32) <= int32(filterVal.(uint32))
	case "int32:uint16":  return val.(int32) <= int32(filterVal.(uint16))
	case "int32:uint8":   return val.(int32) <= int32(filterVal.(uint8))
	case "int32:uint":    return val.(int32) <= int32(filterVal.(uint))

	case "int16:uint64":  return int64(val.(int16)) <= int64(filterVal.(uint64))
	case "int16:uint32":  return int32(val.(int16)) <= int32(filterVal.(uint32))
	case "int16:uint16":  return val.(int16) <= int16(filterVal.(uint16))
	case "int16:uint8":   return val.(int16) <= int16(filterVal.(uint8))
	case "int16:uint":    return val.(int16) <= int16(filterVal.(uint))

	case "int8:uint64":   return int64(val.(int8)) <= int64(filterVal.(uint64))
	case "int8:uint32":   return int32(val.(int8)) <= int32(filterVal.(uint32))
	case "int8:uint16":   return int16(val.(int8)) <= int16(filterVal.(uint16))
	case "int8:uint8":    return val.(int8) <= int8(filterVal.(uint8))
	case "int8:uint":     return val.(int8) <= int8(filterVal.(uint))

	case "int:uint64":    return int64(val.(int)) <= int64(filterVal.(uint64))
	case "int:uint32":    return int32(val.(int)) <= int32(filterVal.(uint32))
	case "int:uint16":    return int16(val.(int)) <= int16(filterVal.(uint16))
	case "int:uint8":     return int8(val.(int)) <= int8(filterVal.(uint8))
	case "int:uint":      return val.(int) <= int(filterVal.(uint))

	// uint

	case "uint64:uint64": return val.(uint64) <= filterVal.(uint64)
	case "uint64:uint32": return val.(uint64) <= uint64(filterVal.(uint32))
	case "uint64:uint16": return val.(uint64) <= uint64(filterVal.(uint16))
	case "uint64:uint8":  return val.(uint64) <= uint64(filterVal.(uint8))
	case "uint64:uint":   return val.(uint64) <= uint64(filterVal.(uint))

	case "uint32:uint64": return uint64(val.(uint32)) <= filterVal.(uint64)
	case "uint32:uint32": return val.(uint32) <= filterVal.(uint32)
	case "uint32:uint16": return val.(uint32) <= uint32(filterVal.(uint16))
	case "uint32:uint8":  return val.(uint32) <= uint32(filterVal.(uint8))
	case "uint32:uint":   return val.(uint32) <= uint32(filterVal.(uint))

	case "uint16:uint64": return uint64(val.(uint16)) <= filterVal.(uint64)
	case "uint16:uint32": return uint32(val.(uint16)) <= filterVal.(uint32)
	case "uint16:uint16": return val.(uint16) <= filterVal.(uint16)
	case "uint16:uint8":  return val.(uint16) <= uint16(filterVal.(uint8))
	case "uint16:uint":   return val.(uint16) <= uint16(filterVal.(uint))

	case "uint8:uint64":  return uint64(val.(uint8)) <= filterVal.(uint64)
	case "uint8:uint32":  return uint32(val.(uint8)) <= filterVal.(uint32)
	case "uint8:uint16":  return uint16(val.(uint8)) <= filterVal.(uint16)
	case "uint8:uint8":   return val.(uint8) <= filterVal.(uint8)
	case "uint8:uint":    return val.(uint8) <= uint8(filterVal.(uint))

	case "uint:uint64":   return uint64(val.(uint)) <= filterVal.(uint64)
	case "uint:uint32":   return uint32(val.(uint)) <= filterVal.(uint32)
	case "uint:uint16":   return uint16(val.(uint)) <= filterVal.(uint16)
	case "uint:uint8":    return uint8(val.(uint)) <= filterVal.(uint8)
	case "uint:uint":     return val.(uint) <= filterVal.(uint)

	// uint & int

	case "uint64:int64":  return int64(val.(uint64)) <= filterVal.(int64)
	case "uint64:int32":  return int64(val.(uint64)) <= int64(filterVal.(int32))
	case "uint64:int16":  return int64(val.(uint64)) <= int64(filterVal.(int16))
	case "uint64:int8":   return int64(val.(uint64)) <= int64(filterVal.(int8))
	case "uint64:int":    return int64(val.(uint64)) <= int64(filterVal.(int))

	case "uint32:int64":  return int64(val.(uint32)) <= filterVal.(int64)
	case "uint32:int32":  return int32(val.(uint32)) <= filterVal.(int32)
	case "uint32:int16":  return int32(val.(uint32)) <= int32(filterVal.(int16))
	case "uint32:int8":   return int32(val.(uint32)) <= int32(filterVal.(int8))
	case "uint32:int":    return int32(val.(uint32)) <= int32(filterVal.(int))

	case "uint16:int64":  return int64(val.(uint16)) <= filterVal.(int64)
	case "uint16:int32":  return int32(val.(uint16)) <= filterVal.(int32)
	case "uint16:int16":  return int16(val.(uint16)) <= filterVal.(int16)
	case "uint16:int8":   return int16(val.(uint16)) <= int16(filterVal.(int8))
	case "uint16:int":    return int16(val.(uint16)) <= int16(filterVal.(int))

	case "uint8:int64":   return int64(val.(uint8)) <= filterVal.(int64)
	case "uint8:int32":   return int32(val.(uint8)) <= filterVal.(int32)
	case "uint8:int16":   return int16(val.(uint8)) <= filterVal.(int16)
	case "uint8:int8":    return int8(val.(uint8))  <= filterVal.(int8)
	case "uint8:int":     return int8(val.(uint8))  <= int8(filterVal.(int))

	case "uint:int64":    return int64(val.(uint))  <= filterVal.(int64)
	case "uint:int32":    return int32(val.(uint))  <= filterVal.(int32)
	case "uint:int16":    return int16(val.(uint))  <= filterVal.(int16)
	case "uint:int8":     return int8(val.(uint))   <= filterVal.(int8)
	case "uint:int":      return int(val.(uint))    <= filterVal.(int)

	// byte

	case "byte:int64":    return int64(val.(byte)) <= filterVal.(int64)
	case "byte:int32":    return int32(val.(byte)) <= filterVal.(int32)
	case "byte:int16":    return int16(val.(byte)) <= filterVal.(int16)
	case "byte:int8":     return int8(val.(byte))  <= filterVal.(int8)
	case "byte:int":      return int(val.(byte))   <= filterVal.(int)

	case "byte:uint64":   return uint64(val.(byte)) <= filterVal.(uint64)
	case "byte:uint32":   return uint32(val.(byte)) <= filterVal.(uint32)
	case "byte:uint16":   return uint16(val.(byte)) <= filterVal.(uint16)
	case "byte:uint8":    return uint8(val.(byte))  <= filterVal.(uint8)
	case "byte:uint":     return uint(val.(byte))   <= filterVal.(uint)

	}

	return false
}

func validateEq(val any, filterVal any) bool {
	types := reflect.TypeOf(val).Kind().String() + ":" + reflect.TypeOf(filterVal).Kind().String()

	switch types {

	// int

	case "int64:int64":   return val.(int64) == filterVal.(int64)
	case "int64:int32":   return val.(int64) == int64(filterVal.(int32))
	case "int64:int16":   return val.(int64) == int64(filterVal.(int16))
	case "int64:int8":    return val.(int64) == int64(filterVal.(int8))
	case "int64:int":     return val.(int64) == int64(filterVal.(int))

	case "int32:int64":   return int64(val.(int32)) == filterVal.(int64)
	case "int32:int32":   return val.(int32) == filterVal.(int32)
	case "int32:int16":   return val.(int32) == int32(filterVal.(int16))
	case "int32:int8":    return val.(int32) == int32(filterVal.(int8))
	case "int32:int":     return val.(int32) == int32(filterVal.(int))

	case "int16:int64":   return int64(val.(int16)) == filterVal.(int64)
	case "int16:int32":   return int32(val.(int16)) == filterVal.(int32)
	case "int16:int16":   return val.(int16) == filterVal.(int16)
	case "int16:int8":    return val.(int16) == int16(filterVal.(int8))
	case "int16:int":     return val.(int16) == int16(filterVal.(int))

	case "int8:int64":    return int64(val.(int8)) == filterVal.(int64)
	case "int8:int32":    return int32(val.(int8)) == filterVal.(int32)
	case "int8:int16":    return int16(val.(int8)) == filterVal.(int16)
	case "int8:int8":     return val.(int8) == filterVal.(int8)
	case "int8:int":      return val.(int8) == int8(filterVal.(int))

	case "int:int64":     return int64(val.(int)) == filterVal.(int64)
	case "int:int32":     return int32(val.(int)) == filterVal.(int32)
	case "int:int16":     return int16(val.(int)) == filterVal.(int16)
	case "int:int8":      return int8(val.(int)) == filterVal.(int8)
	case "int:int":       return val.(int) == filterVal.(int)

	// int & uint

	case "int64:uint64":  return val.(int64) == int64(filterVal.(uint64))
	case "int64:uint32":  return val.(int64) == int64(filterVal.(uint32))
	case "int64:uint16":  return val.(int64) == int64(filterVal.(uint16))
	case "int64:uint8":   return val.(int64) == int64(filterVal.(uint8))
	case "int64:uint":    return val.(int64) == int64(filterVal.(uint))

	case "int32:uint64":  return int64(val.(int32)) == int64(filterVal.(uint64))
	case "int32:uint32":  return val.(int32) == int32(filterVal.(uint32))
	case "int32:uint16":  return val.(int32) == int32(filterVal.(uint16))
	case "int32:uint8":   return val.(int32) == int32(filterVal.(uint8))
	case "int32:uint":    return val.(int32) == int32(filterVal.(uint))

	case "int16:uint64":  return int64(val.(int16)) == int64(filterVal.(uint64))
	case "int16:uint32":  return int32(val.(int16)) == int32(filterVal.(uint32))
	case "int16:uint16":  return val.(int16) == int16(filterVal.(uint16))
	case "int16:uint8":   return val.(int16) == int16(filterVal.(uint8))
	case "int16:uint":    return val.(int16) == int16(filterVal.(uint))

	case "int8:uint64":   return int64(val.(int8)) == int64(filterVal.(uint64))
	case "int8:uint32":   return int32(val.(int8)) == int32(filterVal.(uint32))
	case "int8:uint16":   return int16(val.(int8)) == int16(filterVal.(uint16))
	case "int8:uint8":    return val.(int8) == int8(filterVal.(uint8))
	case "int8:uint":     return val.(int8) == int8(filterVal.(uint))

	case "int:uint64":    return int64(val.(int)) == int64(filterVal.(uint64))
	case "int:uint32":    return int32(val.(int)) == int32(filterVal.(uint32))
	case "int:uint16":    return int16(val.(int)) == int16(filterVal.(uint16))
	case "int:uint8":     return int8(val.(int)) == int8(filterVal.(uint8))
	case "int:uint":      return val.(int) == int(filterVal.(uint))

	// uintX : uintX

	case "uint64:uint64": return val.(uint64) == filterVal.(uint64)
	case "uint64:uint32": return val.(uint64) == uint64(filterVal.(uint32))
	case "uint64:uint16": return val.(uint64) == uint64(filterVal.(uint16))
	case "uint64:uint8":  return val.(uint64) == uint64(filterVal.(uint8))
	case "uint64:uint":   return val.(uint64) == uint64(filterVal.(uint))

	case "uint32:uint64": return uint64(val.(uint32)) == filterVal.(uint64)
	case "uint32:uint32": return val.(uint32) == filterVal.(uint32)
	case "uint32:uint16": return val.(uint32) == uint32(filterVal.(uint16))
	case "uint32:uint8":  return val.(uint32) == uint32(filterVal.(uint8))
	case "uint32:uint":   return val.(uint32) == uint32(filterVal.(uint))

	case "uint16:uint64": return uint64(val.(uint16)) == filterVal.(uint64)
	case "uint16:uint32": return uint32(val.(uint16)) == filterVal.(uint32)
	case "uint16:uint16": return val.(uint16) == filterVal.(uint16)
	case "uint16:uint8":  return val.(uint16) == uint16(filterVal.(uint8))
	case "uint16:uint":   return val.(uint16) == uint16(filterVal.(uint))

	case "uint8:uint64":  return uint64(val.(uint8)) == filterVal.(uint64)
	case "uint8:uint32":  return uint32(val.(uint8)) == filterVal.(uint32)
	case "uint8:uint16":  return uint16(val.(uint8)) == filterVal.(uint16)
	case "uint8:uint8":   return val.(uint8) == filterVal.(uint8)
	case "uint8:uint":    return val.(uint8) == uint8(filterVal.(uint))

	case "uint:uint64":   return uint64(val.(uint)) == filterVal.(uint64)
	case "uint:uint32":   return uint32(val.(uint)) == filterVal.(uint32)
	case "uint:uint16":   return uint16(val.(uint)) == filterVal.(uint16)
	case "uint:uint8":    return uint8(val.(uint)) == filterVal.(uint8)
	case "uint:uint":     return val.(uint) == filterVal.(uint)

	// uintX : intX

	case "uint64:int64":  return int64(val.(uint64)) == filterVal.(int64)
	case "uint64:int32":  return int64(val.(uint64)) == int64(filterVal.(int32))
	case "uint64:int16":  return int64(val.(uint64)) == int64(filterVal.(int16))
	case "uint64:int8":   return int64(val.(uint64)) == int64(filterVal.(int8))
	case "uint64:int":    return int64(val.(uint64)) == int64(filterVal.(int))

	case "uint32:int64":  return int64(val.(uint32)) == filterVal.(int64)
	case "uint32:int32":  return int32(val.(uint32)) == filterVal.(int32)
	case "uint32:int16":  return int32(val.(uint32)) == int32(filterVal.(int16))
	case "uint32:int8":   return int32(val.(uint32)) == int32(filterVal.(int8))
	case "uint32:int":    return int32(val.(uint32)) == int32(filterVal.(int))

	case "uint16:int64":  return int64(val.(uint16)) == filterVal.(int64)
	case "uint16:int32":  return int32(val.(uint16)) == filterVal.(int32)
	case "uint16:int16":  return int16(val.(uint16)) == filterVal.(int16)
	case "uint16:int8":   return int16(val.(uint16)) == int16(filterVal.(int8))
	case "uint16:int":    return int16(val.(uint16)) == int16(filterVal.(int))

	case "uint8:int64":   return int64(val.(uint8)) == filterVal.(int64)
	case "uint8:int32":   return int32(val.(uint8)) == filterVal.(int32)
	case "uint8:int16":   return int16(val.(uint8)) == filterVal.(int16)
	case "uint8:int8":    return int8(val.(uint8))  == filterVal.(int8)
	case "uint8:int":     return int8(val.(uint8))  == int8(filterVal.(int))

	case "uint:int64":    return int64(val.(uint))  == filterVal.(int64)
	case "uint:int32":    return int32(val.(uint))  == filterVal.(int32)
	case "uint:int16":    return int16(val.(uint))  == filterVal.(int16)
	case "uint:int8":     return int8(val.(uint))   == filterVal.(int8)
	case "uint:int":      return int(val.(uint))    == filterVal.(int)

	// byte

	case "byte:int64":    return int64(val.(byte)) == filterVal.(int64)
	case "byte:int32":    return int32(val.(byte)) == filterVal.(int32)
	case "byte:int16":    return int16(val.(byte)) == filterVal.(int16)
	case "byte:int8":     return int8(val.(byte))  == filterVal.(int8)
	case "byte:int":      return int(val.(byte))   == filterVal.(int)

	case "byte:uint64":   return uint64(val.(byte)) == filterVal.(uint64)
	case "byte:uint32":   return uint32(val.(byte)) == filterVal.(uint32)
	case "byte:uint16":   return uint16(val.(byte)) == filterVal.(uint16)
	case "byte:uint8":    return uint8(val.(byte))  == filterVal.(uint8)
	case "byte:uint":     return uint(val.(byte))   == filterVal.(uint)

	}

	return false
}

func validateStrLen(val any, filterVal any) bool {
	switch reflect.TypeOf(val).String() {

	case "[]string":
		var result bool = len(val.([]string)) > 0

		for _, str := range val.([]string) {
			result = result && validateEq(len(str), filterVal)
		}

		return result

	case "string":
		length := len(val.(string))
		return validateEq(length, filterVal)

	}

	return false
}

func validateMinLen(val any, filterVal any) bool {
	switch reflect.TypeOf(val).String() {

	case "[]string":
		length := len(val.([]string))
		return validateMin(filterVal, length)

	case "string":
		length := len(val.(string))
		return validateMin(filterVal, length)

	}

	return false
}

func validateMaxLen(val any, filterVal any) bool {
	switch reflect.TypeOf(val).String() {

	case "[]string":
		length := len(val.([]string))
		return validateMax(length, filterVal)

	case "string":
		length := len(val.(string))
		return validateMax(length, filterVal)

	}

	return false
}

func validateYear(val any, filterVal any) bool {
	if reflect.TypeOf(val).String() != "time.Time" {
		return false
	}

	year := val.(time.Time).Year()
	return validateEq(year, filterVal)
}
