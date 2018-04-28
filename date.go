package main

import (
	"time"
	"fmt"
)
type date time.Time
func main() {
	var a = time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(a)
}
