package main

import "fmt"

type aaa map[string]interface{}

func main()  {
	var b = aaa{}
	b["a"] = 3
	fmt.Println(b)
}
