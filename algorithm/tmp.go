package main

import (
	"fmt"
	"reflect"
)

type a struct {
	key interface{}
	val interface{}
}
type b struct {
	c []*a
	d []a
}

func main() {
	var a = &b{make([]*a,3),make([]a,3)}

	fmt.Println(cap(a.c), cap(a.d))

	fmt.Println(a.c)
	fmt.Println("a.c[1]",a.c[1])
	fmt.Println(a.d)
	fmt.Println("a.d[1]",a.d[1])
	fmt.Println(a.d[1].key==nil)
	fmt.Println(reflect.TypeOf(a.c))
}

func hash(k interface{}) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
		//fmt.Println(h)
	}
	return h
}
