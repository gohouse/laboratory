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

type Ala string

func (a Ala) String() string {
	return string(a)
}
func (a *Ala) Join() string {
	return "";
}
func (a Ala) Split(sep string) []string {

	return strings.Split(a.String(), sep);
}

func parseStr(arg interface{}) string {
	return fmt.Sprint(arg)
}

func main()  {
	var as Ala = "asdf,sdfsa"
	//var as Ala = "asdsa"

	a := as.String()
	fmt.Println(a)

	fmt.Println(as.Split(","))
	//fmt.Println(as.ToPredict())
}
