package main

import "fmt"

func main()  {
	a := method_str

	fmt.Println(a)
	fmt.Println(a(2))
}

func method_str(arg int) int {
	return arg
}
