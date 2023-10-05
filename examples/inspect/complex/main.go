package main

import (
	"fmt"
	"unsafe"
	"yu/golang/internal/app"
)

type CplxInt struct {
	i8  int8
	i16 int16
	i32 int32
	i64 int64
	i   int
}

type CplxUint struct {
	u8  uint8
	u16 uint16
	u32 uint32
	u64 uint64
	u   uint
}

type CplxFloat struct {
	f32 float32
	f64 float64
}

type CplxComplex struct {
	c64  complex64
	c128 complex128
}
type MapStr map[string]MapStr

type ComplexStruct struct {
	Bool    bool
	Int     CplxInt
	Uint    CplxUint
	Float   CplxFloat
	Complex CplxComplex

	String string

	ArrayOfBool [2]bool
	Slice       []bool
	EmptyStruct struct{}
	Map         map[string]string
	MapStr      MapStr
	Channel     chan int
	Function    func()

	Pointer       *ComplexStruct
	UnsafePointer unsafe.Pointer
	Uintptr       uintptr
}

func main() {
	str := "Hello World!"

	mapStr := MapStr{"key": MapStr{}}
	// TODO: avoid infinite tree
	// mapStr["key"] = mapStr

	complex := ComplexStruct{
		Bool: true,
		Int: CplxInt{
			i8:  app.MAX_INT8,
			i16: app.MAX_INT16,
			i32: app.MAX_INT32,
			i64: app.MAX_INT64,
			i:   app.MAX_INT,
		},
		Uint: CplxUint{
			u8:  app.MAX_UINT8,
			u16: app.MAX_UINT16,
			u32: app.MAX_UINT32,
			u64: app.MAX_UINT64,
			u:   app.MAX_UINT,
		},
		Float: CplxFloat{
			f32: app.MAX_FLOAT32,
			f64: app.MAX_FLOAT64,
		},
		Complex: CplxComplex{
			c64:  complex64(3.4028234663852886e+38),
			c128: complex128(1.7976931348623157e+308),
		},

		String: str,

		ArrayOfBool: [2]bool{false, true},
		Slice:       []bool{false, true},
		EmptyStruct: struct{}{},
		Map:         map[string]string{},
		MapStr:      mapStr,
		Channel:     make(chan int),
		Function:    func() {},

		UnsafePointer: unsafe.Pointer(&str),
		Uintptr:       uintptr(unsafe.Pointer(&str)),
	}

	// TODO: avoid infinite tree
	// complex.Pointer = &complex

	fmt.Print(app.InspectData(&complex))
}
