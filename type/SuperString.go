package main

import (
	"fmt"
	"strings"
)

type ala interface {
	String() string
	Join() string
	Split() []string
}

type SuperString string

func (ss SuperString) String() string {
	return string(ss)
}
func (ss SuperString) Length() int {
	return len(ss.String())
}
func (ss SuperString) Index(sep string) int {
	return strings.Index(ss.String(), sep)
}
func (ss SuperString) Split(sep string) []string {
	return strings.Split(ss.String(), sep)
}
func (ss SuperString) Trim(flag ...string) string {
	var sub = " "
	if len(flag) > 0 {
		sub = flag[0]
	}
	return strings.Trim(ss.String(), sub)
}
func (ss SuperString) Replace(old, new string, n int) string {
	return strings.Replace(ss.String(), old, new, n)
}

//func parseStr(arg interface{}) string {
//	return fmt.Sprint(arg)
//}

func main()  {
	var ss SuperString = " asdf,sdfsa"

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
