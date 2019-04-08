package main

import (
	"fmt"
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
	fmt.Println(hash(1000))
}

func hash(k interface{}) uint64 {
	key := fmt.Sprintf("%s", k)
	var h int64 = 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int64(key[i])
		//fmt.Println(h)
	}
	return uint64(h)
}
