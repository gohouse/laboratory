package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`(?=.*[a-zA-Z])(?=.*\d)`)
	//reg := regexp.MustCompile(`([a-zA-Z]+)|(\d+)`)
	//reg := regexp.MustCompile(`[[a-zA-Z]*|\d*]*`)
	res := reg.FindAllString("asdf2323a2s", -1)
	fmt.Println(res)

	//sourceStr := `asdf2323a2s`
	//matched, _ := regexp.MatchString(`(?=.*[a-zA-Z])(?=.*\d){6,16}`, sourceStr)
	//fmt.Printf("%v", matched)
}
