package main

import (
	"fmt"
	"reflect"
)

type Article struct{
	Name string
	Age int
}

func main()  {
	data := &Article{}
	data.Name = "a"
	data2 := new(Article)

	data3 := &Article{}

	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))
	fmt.Println(data2)
	fmt.Println(reflect.TypeOf(data2))
	fmt.Println(data3)
}
