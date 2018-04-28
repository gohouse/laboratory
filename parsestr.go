package main

import (
	"github.com/gohouse/gorose/utils"
	"fmt"
)

func main() {
	var a = 3.1415926535

	var b float32 = 2

	var c = "I'm Jim "

	var res = utils.ParseStr(a)
	var res2 = utils.ParseStr(b)

	fmt.Println(utils.AddSingleQuotes(res))
	fmt.Println(utils.AddSingleQuotes(c))
	fmt.Println(res2)
}
