package main

import (
	_ "github.com/go-sql-driver/mysql" //import DB driver
	"fmt"
	"errors"
)


func main() {
	res := test222(test333)
	fmt.Println(res())
}

func test222(task func() error) func() error {
	return task
}

func test333() error {
	return errors.New("error 222")
}
