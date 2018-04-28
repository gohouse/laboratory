package main

import (
	"github.com/gohouse/study/try"
	"fmt"
)

func main() {
	try.Try(func() {
		panic("asdf")
		fmt.Println(234)
	}).Catch(func(e interface{}) {
		fmt.Println(e)
	})
}
