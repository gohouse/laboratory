package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = []interface{}{"a", "b"}
	convert(a)
}

func convert(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	switch arg.(type) {
	case []string:
	case []int:
	case []interface{}:
		fmt.Println("aaa")
	}
}
