package main

import (
	"fmt"
	"github.com/gohouse/superType"
)

func main()  {
	var ss superType.String = " asdf,sdfsa"

	a:= ss.Split(",")
	b:= ss.Length()
	c:= ss.String()
	d:= ss.Index(",")
	e:= ss.Trim()
	f:= ss.Trim("a")
	g:= ss.Replace("a", "m", -1)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}