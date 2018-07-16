package main

import (
	"fmt"
	"github.com/gohouse/laboratory/utils"
)

func main() {
	a := fmt.Sprint(1)

	fmt.Println(a+"sdfs")

	var s string = "12312sf"
	s = fmt.Sprintf("%s%s",s,"123123")
	fmt.Println(s)
	utils.ParseStr(s)
}
