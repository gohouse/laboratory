package main

import "fmt"
import "reflect"

type T struct {}

func (t *T) Foo() {
	fmt.Println("foo")
}

type MyStruct struct {
	id int
}

func (t *T) Bar() {
	fmt.Println("foo")
}

//func main() {
//	var t T
//	reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})
//}

// modified main () that throws errors
func main() {
	var t T
	bar := reflect.ValueOf(&t).​MethodByName("Bar")
	ms := &MyStruct{5}

	params := make([]reflect.Value,0)
	params = append(params, reflect.ValueOf(ms))
	bar.Call(params)
}
