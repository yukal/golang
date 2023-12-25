package test

import (
	"math"
	"reflect"
	"testing"
	"yu/golang/internal/app"

	. "github.com/franela/goblin"
)

// go test ./test/unit/types...
// go test -v -run TestNumericTypes ./test/unit/types

func TestNumericTypes(t *testing.T) {
	g := Goblin(t)

	type TestItem struct {
		In  any
		Out any
	}

	g.Describe("Invalid Types", func() {
		// int

		g.Xit("return int8 when calling MinInt8, MaxInt8", func() {
			g.Assert(reflect.ValueOf(math.MinInt8).Kind()).Equal(reflect.Int8)
			g.Assert(reflect.ValueOf(math.MaxInt8).Kind()).Equal(reflect.Int8)
		})

		g.Xit("return int16 when calling MinInt16, MaxInt16", func() {
			g.Assert(reflect.ValueOf(math.MinInt16).Kind()).Equal(reflect.Int16)
			g.Assert(reflect.ValueOf(math.MaxInt16).Kind()).Equal(reflect.Int16)
		})

		g.Xit("return int32 when calling MinInt32, MaxInt32", func() {
			g.Assert(reflect.ValueOf(math.MinInt32).Kind()).Equal(reflect.Int32)
			g.Assert(reflect.ValueOf(math.MaxInt32).Kind()).Equal(reflect.Int32)
		})

		g.Xit("return int64 when calling MinInt64, MaxInt64", func() {
			g.Assert(reflect.ValueOf(math.MinInt64).Kind()).Equal(reflect.Int64)
			g.Assert(reflect.ValueOf(math.MaxInt64).Kind()).Equal(reflect.Int64)
		})

		g.It("return int when calling MinInt, MaxInt", func() {
			g.Assert(reflect.ValueOf(math.MinInt).Kind()).Equal(reflect.Int)
			g.Assert(reflect.ValueOf(math.MaxInt).Kind()).Equal(reflect.Int)
		})

		// unit

		g.Xit("return uint8 when calling MaxUint8", func() {
			g.Assert(reflect.ValueOf(math.MaxUint8).Kind()).Equal(reflect.Uint8)
		})

		g.Xit("return uint16 when calling MaxUint16", func() {
			g.Assert(reflect.ValueOf(math.MaxUint16).Kind()).Equal(reflect.Uint16)
		})

		g.Xit("return uint32 when calling MaxUint32", func() {
			g.Assert(reflect.ValueOf(math.MaxUint32).Kind()).Equal(reflect.Uint32)
		})

		g.It("return uint64 when calling MaxUint64", func() {
			g.Assert(reflect.ValueOf(uint64(math.MaxUint64)).Kind()).Equal(reflect.Uint64)
		})

		g.It("return uint when calling MaxUint", func() {
			g.Assert(reflect.ValueOf(uint(math.MaxUint)).Kind()).Equal(reflect.Uint)
		})
	})

	g.Describe("Valid Types", func() {
		// int

		g.It("return int8 when calling MinInt8, MaxInt8", func() {
			g.Assert(reflect.ValueOf(app.MIN_INT8).Kind()).Equal(reflect.Int8)
			g.Assert(reflect.ValueOf(app.MAX_INT8).Kind()).Equal(reflect.Int8)
		})

		g.It("return int16 when calling MinInt16, MaxInt16", func() {
			g.Assert(reflect.ValueOf(app.MIN_INT16).Kind()).Equal(reflect.Int16)
			g.Assert(reflect.ValueOf(app.MAX_INT16).Kind()).Equal(reflect.Int16)
		})

		g.It("return int32 when calling MinInt32, MaxInt32", func() {
			g.Assert(reflect.ValueOf(app.MIN_INT32).Kind()).Equal(reflect.Int32)
			g.Assert(reflect.ValueOf(app.MAX_INT32).Kind()).Equal(reflect.Int32)
		})

		g.It("return int64 when calling MinInt64, MaxInt64", func() {
			g.Assert(reflect.ValueOf(app.MIN_INT64).Kind()).Equal(reflect.Int64)
			g.Assert(reflect.ValueOf(app.MAX_INT64).Kind()).Equal(reflect.Int64)
		})

		g.It("return int when calling MinInt, MaxInt", func() {
			g.Assert(reflect.ValueOf(app.MIN_INT).Kind()).Equal(reflect.Int)
			g.Assert(reflect.ValueOf(app.MAX_INT).Kind()).Equal(reflect.Int)
		})

		// unit

		g.It("return uint8 when calling MaxUint8", func() {
			g.Assert(reflect.ValueOf(app.MIN_UINT8).Kind()).Equal(reflect.Uint8)
			g.Assert(reflect.ValueOf(app.MAX_UINT8).Kind()).Equal(reflect.Uint8)
		})

		g.It("return uint16 when calling MaxUint16", func() {
			g.Assert(reflect.ValueOf(app.MIN_UINT16).Kind()).Equal(reflect.Uint16)
			g.Assert(reflect.ValueOf(app.MAX_UINT16).Kind()).Equal(reflect.Uint16)
		})

		g.It("return uint32 when calling MaxUint32", func() {
			g.Assert(reflect.ValueOf(app.MIN_UINT32).Kind()).Equal(reflect.Uint32)
			g.Assert(reflect.ValueOf(app.MAX_UINT32).Kind()).Equal(reflect.Uint32)
		})

		g.It("return uint64 when calling MaxUint64", func() {
			g.Assert(reflect.ValueOf(app.MIN_UINT64).Kind()).Equal(reflect.Uint64)
			g.Assert(reflect.ValueOf(app.MAX_UINT64).Kind()).Equal(reflect.Uint64)
		})

		g.It("return uint when calling MaxUint", func() {
			g.Assert(reflect.ValueOf(app.MIN_UINT).Kind()).Equal(reflect.Uint)
			g.Assert(reflect.ValueOf(app.MAX_UINT).Kind()).Equal(reflect.Uint)
		})
	})
}
