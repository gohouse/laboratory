package main

import (
	"fmt"
	"reflect"
)

type Student2 struct {
	Name  string `json:"stu_name"`
	Age   int	`json:"age"`
	Score float32

}

func (s Student2) Print(){
	fmt.Println(s)
}


func (s Student2) Set(name string,age int,score float32){
	s.Age = age
	s.Name = name
	s.Score = score
}


func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("Tag:%s\n",tag)
}

func main() {
	var a = Student2{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}
	TestStruct(&a)
}