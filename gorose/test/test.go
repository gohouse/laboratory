package main

import (
	"fmt"
	"github.com/gohouse/laboratory/dp"
)

func main() {
	connection := gorose.Open("mysql")
	//db := connection.NewDB()
	res := connection.Table("aaa").BuildSql("select")

	fmt.Println(res)
}
