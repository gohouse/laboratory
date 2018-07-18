package main

import (
	"fmt"
	"time"
)

func main() {
	a := fmt.Sprintf("static/%v-%v-%v.xlsx",time.Now().Format("20060102150405"),
		"名字", 23)
	fmt.Println(a)
}
