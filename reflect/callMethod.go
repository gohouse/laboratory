package main

import (
	"fmt"
	"reflect"
)

type FooBar struct {
}

func (f *FooBar) FooBarAdd(argOne, argTwo float64) float64 {

	return argOne + argTwo
}

func main() {
	foobar := &FooBar{}
	resultCallByName :=reflect.ValueOf(foobar).MethodByName("FooBarAdd2")
	//.
	//	Call([]reflect.Type{reflect.ValueOf(123.4),reflect.ValueOf(432.1)})

	//fmt.Println(resultCallByName[0].Float())
	fmt.Println(resultCallByName.IsValid())


	//if methodValue := reflect.ValueOf(foobar).MethodByName("FooBarAdd2");
	//	methodValue.IsValid() {
	//
	//}
}
