package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Println(GetFieldName(Student{}))
	fmt.Println(GetFieldName(&Student{}))
	fmt.Println(GetFieldName(""))
	fmt.Println(GetTagName(&Student{}))
	fmt.Println(GetStructName(&Student{}))
	//fmt.Println(Student{}.test())
	RunStructMethod(&Student{})
}

type table interface {
	test() string
}

type Student struct {
	Name  string `orm:"name" validate:"string"`
	Age   int    `orm:"age"`
	Grade int    `orm:"-"`
}

func (s Student) test2() string {
	return "tablename"
}

func GetStructName(structName interface{}) string {
	// 获取type
	t := reflect.TypeOf(structName)
	// 如果是反射Ptr类型, 就获取他的 element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func RunStructMethod(structName interface{})  {
	if i,ok := structName.(table);ok{
		fmt.Println(i.test())
	} else {
		fmt.Println("未实现")
	}
}

func StructSetValue(structName interface{})  {
	// 获取type
	t := reflect.TypeOf(structName)
	// 如果是反射Ptr类型, 就获取他的 element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	//structName.(Student).test()
	// 判断是否是struct
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return
	}
	//t.MethodByName("test")


	m, ok := t.MethodByName("test")
	if !ok {
		fmt.Println("ptr method by name failed")
		return
	}
	table := m.Func.Call([]reflect.Value{reflect.ValueOf(t), reflect.ValueOf("23")})[0].String()
	fmt.Println(table)
}

//获取结构体中字段的名称
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

//获取结构体中Tag的值，如果没有tag则返回字段值
func GetTagName(structName interface{}) []string {
	// 获取type
	t := reflect.TypeOf(structName)
	// 如果是反射Ptr类型, 就获取他的 element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	//fmt.Println(t.Name())

	// 判断是否是struct
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	// 获取字段数量
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		//fieldName := t.Field(i).Name
		// tag 名字
		tagName := t.Field(i).Tag.Get("orm")
		// tag为-时, 不解析
		if tagName=="-" || tagName=="" {
			// 字段名字
			tagName = t.Field(i).Name
		}
		result = append(result, tagName)
	}
	return result
}
