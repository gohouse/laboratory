package main

import "fmt"

func main() {

	a := []int{1,2,3,4,5}

	lenA := len(a)

	fmt.Println(a[lenA-1:])
}
