package main

import "fmt"

type a struct {
	A string
}
type b struct {
	a
	//B string
}
func main() {
	var b1 b
	b1.A = "aa"

	fmt.Println(b1)
}
