package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type user2 struct {
	Name string
}

type user3 struct {
	Slave []*user2
	Master *user2
}

func main() {
	var u user2
	u.Name = "fizz"
	res,err := json.Marshal(u)
	fmt.Println(string(res), err)

	var u2 user3
	//decode(`{"Master":{"Name":"fizz2"}}`, &u2)
	decode(`{"Name":"fizz2"}`, &u2)
	fmt.Println(u2)
	fmt.Println(u2.Master)
}

func decode(str string, v interface{})  {
	srcElem := reflect.Indirect(reflect.ValueOf(v))
	fieldType := srcElem.FieldByName("Master").Type().Elem()
	fieldElem := reflect.New(fieldType)
	tmp := fieldElem.Interface()
	json.Unmarshal([]byte(str), &tmp)
	srcElem.FieldByName("Master").Set(fieldElem)
}
