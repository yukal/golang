package main

import (
	"errors"
	"math"
	"reflect"
	"time"
)

const MIN_INT8 = int8(math.MinInt8)
const MAX_INT8 = int8(math.MaxInt8)
const MIN_INT16 = int16(math.MinInt16)
const MAX_INT16 = int16(math.MaxInt16)
const MIN_INT32 = int32(math.MinInt32)
const MAX_INT32 = int32(math.MaxInt32)
const MIN_INT64 = int64(math.MinInt64)
const MAX_INT64 = int64(math.MaxInt64)
const MIN_INT = math.MinInt
const MAX_INT = math.MaxInt

const MIN_UINT8 = uint8(0)
const MAX_UINT8 = uint8(math.MaxUint8)
const MIN_UINT16 = uint16(0)
const MAX_UINT16 = uint16(math.MaxUint16)
const MIN_UINT32 = uint32(0)
const MAX_UINT32 = uint32(math.MaxUint32)
const MIN_UINT64 = uint64(0)
const MAX_UINT64 = uint64(math.MaxUint64)
const MIN_UINT = uint(0)
const MAX_UINT = uint(math.MaxUint)

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

// TODO: Refactoring
func Uint8ToInt8(val uint8) int8 {
	return int8(Uint8ToInt16(val))
}

func Uint8ToInt16(val uint8) int16 {
	return int16(val) - math.MaxInt8 - 1
}

func Uint8ToInt32(val uint8) int32 {
	return int32(val) - math.MaxInt8 - 1
}

func Uint8ToInt64(val uint8) int64 {
	return int64(val) - math.MaxInt8 - 1
}

func Uint8ToInt(val uint8) int {
	return int(val) - math.MaxInt8 - 1
}

// TODO: Refactoring
func Int8ToUint8(val int8) uint8 {
	half := uint8(math.MaxInt8)
	residue := uint8(val)

	if val < 0 {
		offset := residue - half - 1
		return uint8(0) + offset
	} else {
		offset := half - residue
		return uint8(255) - offset
	}
}

func Int8ToUint16(val int8) uint16 {
	return uint16(Int8ToUint8(val))
}

func Int8ToUint32(val int8) uint32 {
	return uint32(Int8ToUint8(val))
}

func Int8ToUint64(val int8) uint64 {
	return uint64(Int8ToUint8(val))
}

func Int8ToUint(val int8) uint {
	return uint(Int8ToUint8(val))
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

	// int & uint

	// solved
	case "int8:uint8":    return Int8ToUint8(val.(int8)) >= filterVal.(uint8)
	case "int8:uint16":   return Int8ToUint16(val.(int8)) >= filterVal.(uint16)
	case "int8:uint32":   return Int8ToUint32(val.(int8)) >= filterVal.(uint32)
	case "int8:uint64":   return Int8ToUint64(val.(int8)) >= filterVal.(uint64)
	case "int8:uint":     return Int8ToUint(val.(int8)) >= filterVal.(uint)

	case "int16:uint8":   panic(errors.New("int16:uint8"))
	case "int16:uint16":  panic(errors.New("int16:uint16"))
	case "int16:uint32":  panic(errors.New("int16:uint32"))
	case "int16:uint64":  panic(errors.New("int16:uint64"))
	case "int16:uint":    panic(errors.New("int16:uint"))

	case "int32:uint8":   panic(errors.New("int32:uint8"))
	case "int32:uint16":  panic(errors.New("int32:uint16"))
	case "int32:uint32":  panic(errors.New("int32:uint32"))
	case "int32:uint64":  panic(errors.New("int32:uint64"))
	case "int32:uint":    panic(errors.New("int32:uint"))

	case "int64:uint8":   panic(errors.New("int64:uint8"))
	case "int64:uint16":  panic(errors.New("int64:uint16"))
	case "int64:uint32":  panic(errors.New("int64:uint32"))
	case "int64:uint64":  panic(errors.New("int64:uint64"))
	case "int64:uint":    panic(errors.New("int64:uint"))

	case "int:uint8":     panic(errors.New("int:uint8"))
	case "int:uint16":    panic(errors.New("int:uint16"))
	case "int:uint32":    panic(errors.New("int:uint32"))
	case "int:uint64":    panic(errors.New("int:uint64"))
	case "int:uint":      panic(errors.New("int:uint"))

	// uint

	case "uint:uint8":    return val.(uint) >= uint(filterVal.(uint8))
	case "uint:uint16":   return val.(uint) >= uint(filterVal.(uint16))
	case "uint:uint32":   return val.(uint) >= uint(filterVal.(uint32))
	case "uint:uint64":   return uint64(val.(uint)) >= filterVal.(uint64)
	case "uint:uint":     return val.(uint) >= filterVal.(uint)

	case "uint64:uint8":  return val.(uint64) >= uint64(filterVal.(uint8))
	case "uint64:uint16": return val.(uint64) >= uint64(filterVal.(uint16))
	case "uint64:uint32": return val.(uint64) >= uint64(filterVal.(uint32))
	case "uint64:uint64": return val.(uint64) >= filterVal.(uint64)
	case "uint64:uint":   return val.(uint64) >= uint64(filterVal.(uint))

	case "uint32:uint8":  return val.(uint32) >= uint32(filterVal.(uint8))
	case "uint32:uint16": return val.(uint32) >= uint32(filterVal.(uint16))
	case "uint32:uint32": return val.(uint32) >= filterVal.(uint32)
	case "uint32:uint64": return uint64(val.(uint32)) >= filterVal.(uint64)
	case "uint32:uint":   return uint(val.(uint32)) >= filterVal.(uint)

	case "uint16:uint8":  return val.(uint16) >= uint16(filterVal.(uint8))
	case "uint16:uint16": return val.(uint16) >= filterVal.(uint16)
	case "uint16:uint32": return uint32(val.(uint16)) >= filterVal.(uint32)
	case "uint16:uint64": return uint64(val.(uint16)) >= filterVal.(uint64)
	case "uint16:uint":   return uint(val.(uint16)) >= filterVal.(uint)

	case "uint8:uint8":   return val.(uint8) >= filterVal.(uint8)
	case "uint8:uint16":  return uint16(val.(uint8)) >= filterVal.(uint16)
	case "uint8:uint32":  return uint32(val.(uint8)) >= filterVal.(uint32)
	case "uint8:uint64":  return uint64(val.(uint8)) >= filterVal.(uint64)
	case "uint8:uint":    return uint(val.(uint8)) >= filterVal.(uint)

	// uint & int

	// solved
	case "uint8:int8":    return Uint8ToInt8(val.(uint8)) >= filterVal.(int8)
	case "uint8:int16":   return Uint8ToInt16(val.(uint8)) >= filterVal.(int16)
	case "uint8:int32":   return Uint8ToInt32(val.(uint8)) >= filterVal.(int32)
	case "uint8:int64":   return Uint8ToInt64(val.(uint8)) >= filterVal.(int64)
	case "uint8:int":     return Uint8ToInt(val.(uint8))  >= filterVal.(int)

	case "uint16:int8":   panic(errors.New("uint16:int8"))
	case "uint16:int16":  panic(errors.New("uint16:int16"))
	case "uint16:int32":  panic(errors.New("uint16:int32"))
	case "uint16:int64":  panic(errors.New("uint16:int64"))
	case "uint16:int":    panic(errors.New("uint16:int"))

	case "uint32:int8":   panic(errors.New("uint32:int8"))
	case "uint32:int16":  panic(errors.New("uint32:int16"))
	case "uint32:int32":  panic(errors.New("uint32:int32"))
	case "uint32:int64":  panic(errors.New("uint32:int64"))
	case "uint32:int":    panic(errors.New("uint32:int"))

	case "uint64:int8":   panic(errors.New("uint64:int8"))
	case "uint64:int16":  panic(errors.New("uint64:int16"))
	case "uint64:int32":  panic(errors.New("uint64:int32"))
	case "uint64:int64":  panic(errors.New("uint64:int64"))
	case "uint64:int":    panic(errors.New("uint64:int"))

	case "uint:int8":     panic(errors.New("uint:int8"))
	case "uint:int16":    panic(errors.New("uint:int16"))
	case "uint:int32":    panic(errors.New("uint:int32"))
	case "uint:int64":    panic(errors.New("uint:int64"))
	case "uint:int":      panic(errors.New("uint:int"))

	// byte

	case "byte:int64":    return int64(val.(byte))  >= filterVal.(int64)
	case "byte:int32":    return int32(val.(byte))  >= filterVal.(int32)
	case "byte:int16":    return int16(val.(byte))  >= filterVal.(int16)
	case "byte:int8":     return int8(val.(byte))   >= filterVal.(int8)
	case "byte:int":      return int(val.(byte))    >= filterVal.(int)
	case "byte:uint64":   return uint64(val.(byte)) >= filterVal.(uint64)
	case "byte:uint32":   return uint32(val.(byte)) >= filterVal.(uint32)
	case "byte:uint16":   return uint16(val.(byte)) >= filterVal.(uint16)
	case "byte:uint":     return uint(val.(byte))   >= filterVal.(uint)

	// ...

	case "byte:uint8":    return val.(byte)   >= filterVal.(uint8)
	case "uint8:byte":    return val.(uint8)  >= filterVal.(byte)
	case "int32:rune":    return val.(int32)  >= filterVal.(rune)
	case "rune:int32":    return val.(rune)   >= filterVal.(int32)

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
