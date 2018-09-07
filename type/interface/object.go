package main

import (
	"reflect"
	"fmt"
)

type Where map[string]interface{}
type MultiWhere []Where

func main()  {
	var where = Where{"a":1}
	//var where2 = MultiWhere{Where{"a":1},Where{"b":2}}
	var where2 []Where
	where2 = append(where2, Where{"c":3})
	res := getWhere(where)

	fmt.Println(res)
	fmt.Println(where2)
}

func getWhere(args ...interface{}) reflect.Type {
	switch args[0].(type) {
	case Where:
	case MultiWhere:
	case string:
	default:

	}

	return reflect.TypeOf(args[0])
}
