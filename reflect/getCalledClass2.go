package main

import (
	"fmt"
	"runtime"
)

type a3 struct {

}

func (a a3) foo() {
	pc, _, _, ok := runtime.Caller(0)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fmt.Printf("called from %s\n", details.Name())
	}
}


func main() {
	var a a3
	a.foo()

}