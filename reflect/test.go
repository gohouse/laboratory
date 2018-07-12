package main

import (
	"fmt"
)

type User struct {
	Name string `orm:"name"`
	Age int `orm:"age"`
}

func main() {
	//u := &User{
	//	"fizz",
	//	18,
	//}

	var u User

	fmt.Println(&u.Name)

	//v := reflect.ValueOf(u)
	//fmt.Println(v)
	//k := v.Kind()
	//fmt.Println(k)
	//
	//iv := v.Interface()
	//fmt.Println(iv)
	//
	//fmt.Println(reflect.TypeOf(iv))
}
