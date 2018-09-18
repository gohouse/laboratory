package main

import (
	"fmt"
	"unsafe"
)

type TestStructTobytes struct {
	data int64
}

func main() {
	var testStruct = &TestStructTobytes{100}
	a := StructToBytes(testStruct)
	b := BytesToStruct(a)
	fmt.Println(a)
	fmt.Println(b)
}

func StructToBytes(testStruct *TestStructTobytes) []byte {
	type sliceMock struct {
		addr uintptr
		len  int
		cap  int
	}
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &sliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))

	return data
}

func BytesToStruct(b []byte) *TestStructTobytes {
	return *(**TestStructTobytes)(unsafe.Pointer(&b))
}
