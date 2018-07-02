package main

import (
	"fmt"
	"reflect"
)

func main()  {


	type Task struct {
		Id int64
		Money float64 `xorm:"money"`
		PlatformName string `xorm:"platformName"`
	}

	var task Task


	t := reflect.TypeOf(task)
	n := t.Name()
	e := reflect.ValueOf(task)
	fmt.Println(n,e.Kind())
}
