package main

import (
	"fmt"
)

func test(a []int) {
	a = append(a, 4)
	a[0] = 10
	fmt.Println(a, len(a), cap(a))
}

func main() {
	var a []int

	a = []int{1, 2, 3}
	fmt.Println(a, len(a), cap(a))
	test(a)
	fmt.Println(a, len(a), cap(a))
}