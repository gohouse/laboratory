package main

import (
	"fmt"
	"reflect"
)

type T2 struct {
	A int
	B string
}

func main() {
	t := T2{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset_Strip")
	fmt.Println("t is now", t)
}