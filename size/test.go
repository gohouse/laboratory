package main

import (
	"fmt"
	"unsafe"
)

func main() {
	printInt()
}

func printUint()  {
	var a uint
	var b uint32
	var c uint64

	fmt.Println("size of a: ", unsafe.Sizeof(a))
	fmt.Println("size of b: ", unsafe.Sizeof(b))
	fmt.Println("size of c: ", unsafe.Sizeof(c))
}

func printInt()  {
	var a int
	var b int32
	var c int64

	fmt.Println("size of a: ", unsafe.Sizeof(a))
	fmt.Println("size of b: ", unsafe.Sizeof(b))
	fmt.Println("size of c: ", unsafe.Sizeof(c))
}