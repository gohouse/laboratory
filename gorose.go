package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/examples/config"
	"github.com/gohouse/gorose"
	"fmt"
)

func main() {
	c,err := gorose.Open(config.DbConfig, "mysql_dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
	var conn *gorose.Connection
	fmt.Println(conn)
	//getUserList(conn)
}

func getUserList(conn *gorose.Connection)  {
	res,err := conn.Table("users").First()

	fmt.Println(err)
	fmt.Println(res)
}
