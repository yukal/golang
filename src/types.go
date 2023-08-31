package src

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"
)

// https://go.dev/ref/spec#Numeric_types
//

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
const MIN_RUNE = MIN_INT32
const MAX_RUNE = MAX_INT32

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
const MIN_BYTE = MIN_UINT8
const MAX_BYTE = MAX_UINT8

const MIN_FLOAT32 = float32(0)
const MAX_FLOAT32 = float32(math.MaxFloat32)
const MIN_FLOAT64 = float64(0)
const MAX_FLOAT64 = float64(math.MaxFloat64)

const MIN_COMPLEX64 = complex64(0)
const MIN_COMPLEX128 = complex128(0)

func Uint8ToInt8(val uint8) int8 {
	if val >= math.MaxInt8 {
		return int8(math.MaxInt8)
	}

	return int8(val)
}

func Uint8ToInt16(val uint8) int16 {
	return int16(val)
}

func Uint8ToInt32(val uint8) int32 {
	return int32(val)
}

func Uint8ToInt64(val uint8) int64 {
	return int64(val)
}

func Uint8ToInt(val uint8) int {
	return int(val)
}

func Uint16ToInt8(val uint16) int8 {
	if val >= math.MaxInt8 {
		return int8(math.MaxInt8)
	}

	return int8(val)
}

func Uint16ToInt16(val uint16) int16 {
	if val >= math.MaxInt16 {
		return int16(math.MaxInt16)
	}

	return int16(val)
}

func Uint16ToInt32(val uint16) int32 {
	return int32(val)
}

func Uint32ToInt8(val uint32) int8 {
	if val >= math.MaxInt8 {
		return int8(math.MaxInt8)
	}

	return int8(val)
}

func Uint32ToInt16(val uint32) int16 {
	if val >= math.MaxInt16 {
		return int16(math.MaxInt16)
	}

	return int16(val)
}

func Uint32ToInt32(val uint32) int32 {
	if val >= math.MaxInt32 {
		return int32(math.MaxInt32)
	}

	return int32(val)
}

// TODO: Refactoring
func Uint8ToInt8Range(val uint8) int8 {
	return int8(Uint8ToInt16Range(val))
}

func Uint8ToInt16Range(val uint8) int16 {
	return int16(val) - math.MaxInt8 - 1
}

func Uint8ToInt32Range(val uint8) int32 {
	return int32(val) - math.MaxInt8 - 1
}

func Uint8ToInt64Range(val uint8) int64 {
	return int64(val) - math.MaxInt8 - 1
}

func Uint8ToIntRange(val uint8) int {
	return int(val) - math.MaxInt8 - 1
}

// TODO: Refactoring
func Int8ToUint8Range(val int8) uint8 {
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

func Int8ToUint16Range(val int8) uint16 {
	return uint16(Int8ToUint8Range(val))
}

func Int8ToUint32Range(val int8) uint32 {
	return uint32(Int8ToUintRange(val))
}

func Int8ToUint64Range(val int8) uint64 {
	return uint64(Int8ToUintRange(val))
}

func Int8ToUintRange(val int8) uint {
	return uint(Int8ToUint8Range(val))
}

func InspectData(data any) string {
	buf := new(bytes.Buffer)

	inspectRecursively(buf, reflect.ValueOf(data), 0)

	return strings.TrimSpace(buf.String())
}

func InspectDataW(writer io.Writer, data any) {
	inspectRecursively(writer, reflect.ValueOf(data), 0)
}

// TODO: fix generation the nested data
// Problem: Concurrently scans the nested struct. Each time it generates a string with a random ordering of nested keys. So it is hard to test with generated data
func inspectRecursively(writer io.Writer, data reflect.Value, depth int) {
	indent := strings.Repeat(" ", depth*2)

	switch data.Kind() {
	case reflect.Slice:
		for i := 0; i < data.Len(); i++ {
			value := data.Index(i)

			switch {
			case value.Kind() == reflect.String:
				writer.Write([]byte(
					fmt.Sprintf("\n%s%v: %#v", indent, i, value.Interface())))

			case IsPrimitive(value):
				writer.Write([]byte(
					fmt.Sprintf("\n%s%v: %v", indent, i, value.Interface())))

			default:
				writer.Write([]byte(
					fmt.Sprintf("\n%s%v:", indent, i)))
				inspectRecursively(writer, reflect.ValueOf(value.Interface()), depth+1)
			}
		}

	case reflect.Map:
		for _, key := range data.MapKeys() {
			value := data.MapIndex(key)

			switch {
			case value.Kind() == reflect.String:
				writer.Write([]byte(
					fmt.Sprintf("\n%s%s: %#v", indent, key.Interface(), value.Interface())))

			case IsPrimitive(value):
				writer.Write([]byte(
					fmt.Sprintf("\n%s%s: %v", indent, key.Interface(), value.Interface())))

			default:
				writer.Write([]byte(
					fmt.Sprintf("\n%s%s: ", indent, key.Interface())))
				inspectRecursively(writer, reflect.ValueOf(value.Interface()), depth+1)
			}
		}

	case reflect.Struct:
		refType := data.Type()

		for i := 0; i < refType.NumField(); i++ {
			dataField := refType.Field(i)
			dataValue := reflect.Indirect(data).
				FieldByName(dataField.Name)

			switch {
			case dataValue.Kind() == reflect.String:
				writer.Write([]byte(
					fmt.Sprintf("\n%s%v: %#v", indent, dataField.Name, dataValue.Interface())))

			case IsPrimitive(dataValue):
				writer.Write([]byte(
					fmt.Sprintf("%s%s: %v\n", indent, dataField.Name, dataValue.Interface())))

			default:
				writer.Write([]byte(
					fmt.Sprintf("\n%s%v: ", indent, dataField.Name)))
				inspectRecursively(writer, dataValue, depth+1)
			}
		}

	case reflect.String:
		writer.Write([]byte(
			fmt.Sprintf("%#v", data.Interface())))

	default:
		writer.Write([]byte(
			fmt.Sprintf("%T(%[1]v)", data.Interface())))
	}
}

func IsNumeric(val any) bool {
	ref := reflect.ValueOf(val)
	if ref.Type().String() == "reflect.Value" {
		ref = val.(reflect.Value)
	}

	switch ref.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	}

	return false
}

func IsPrimitive(val any) bool {
	ref := reflect.ValueOf(val)
	if ref.Type().String() == "reflect.Value" {
		ref = val.(reflect.Value)
	}

	switch ref.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Bool, reflect.String:
		return true
	}

	return false
}
