package a

import (
	"fmt"
	"github.com/gohouse/laboratory/import/b"
	"reflect"
)

type AA struct {
	bT b.B
	AAA string
}

func A() string {
	return "aaa"
}

func CallC() interface{} {
	//var tmpB b.B
	var aaT = &AA{}
	aaT.AAA = "bbbbbb"
	fmt.Println(reflect.TypeOf(aaT.bT))
	//res := c.C(aaT.bT)
	//fmt.Println(res)

	return aaT
}