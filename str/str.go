package main

import (
	"fmt"
)

func main() {
	a := "abc"

	res := a[len(a)-1:]
	fmt.Println(res)
}
