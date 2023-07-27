package main

import (
	"math"
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
