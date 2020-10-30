package main

import (
	"fmt"
	"unsafe"
)
import "reflect"

// 通过反射，对user进行赋值
type user struct {
	Name string
	Age  int
	//feature map[string]interface{}
}

func (u user) test() {
	fmt.Println("test")
}

func main() {
	//var u interface{}
	//u=new(user)
	//value:=reflect.ValueOf(u)
	////value:=reflect.TypeOf(u).Elem()
	//if value.Kind()==reflect.Ptr{
	//	elem:=value.Elem()
	//	name:=elem.FieldByName("name")
	//	if name.Kind()==reflect.ToPredict{
	//		*(*string)(unsafe.Pointer(name.Addr().Pointer())) = "fangwendong"
	//	}
	//
	//	age:=elem.FieldByName("age")
	//	if age.Kind()==reflect.Int{
	//		*(*int)(unsafe.Pointer(age.Addr().Pointer())) =24
	//	}
	//
	//	feature:=elem.FieldByName("feature")
	//	if feature.Kind()==reflect.Map{
	//		*(*map[string]interface{})(unsafe.Pointer(feature.Addr().Pointer())) =map[string]interface{}{
	//			"爱好":"篮球",
	//			"体重":60,
	//			"视力":5.2,
	//		}
	//	}
	//
	//}
	//
	//fmt.Println(u)

	//setValue(&user{})
	var u []user
	//u = append(u, user{Name:"d"})
	SetStructFieldByJsonName(&u, map[string]interface{}{"name": "aaa", "age": 8})

	fmt.Println(u)
}

func setValue(structName interface{}) {
	value := reflect.ValueOf(structName)
	//value:=reflect.TypeOf(structName)
	if value.Kind() == reflect.Ptr {
		elem := value.Elem()
		name := elem.FieldByName("name")
		if name.Kind() == reflect.String {
			*(*string)(unsafe.Pointer(name.Addr().Pointer())) = "fangwendong"
		}

		age := elem.FieldByName("age")
		if age.Kind() == reflect.Int {
			*(*int)(unsafe.Pointer(age.Addr().Pointer())) = 24
		}

		feature := elem.FieldByName("feature")
		if feature.Kind() == reflect.Map {
			*(*map[string]interface{})(unsafe.Pointer(feature.Addr().Pointer())) = map[string]interface{}{
				"爱好": "篮球",
				"体重": 60,
				"视力": 5.2,
			}
		}

	}

	fmt.Println(structName)
}

func newElemFunc(sliceValue reflect.Value, fields []string) reflect.Value {
	elemType := sliceValue.Type().Elem()
	switch elemType.Kind() {
	case reflect.Slice:
		slice := reflect.MakeSlice(elemType, len(fields), len(fields))
		x := reflect.New(slice.Type())
		x.Elem().Set(slice)
		return x
		//case reflect.Map:
		//	mp := reflect.MakeMap(elemType)
		//	x := reflect.New(mp.Type())
		//	x.Elem().Set(mp)
		//	return x
	}
	return reflect.New(elemType)
}

//将结构体里的成员按照json名字来赋值
func SetStructFieldByJsonName(ptr interface{}, fields map[string]interface{}) {
	fmt.Println("fields:", fields)
	// 结果集
	var result reflect.Value
	// 获取slice的Elem	([]UserStruct)
	sliceValue := reflect.Indirect(reflect.ValueOf(ptr))
	// 如果是切片Struct
	if sliceValue.Kind() == reflect.Slice {
		// slice的Elem获取类型的Elem (UserStruct)
		sliceElementType := sliceValue.Type().Elem()
		fmt.Println(sliceElementType, sliceElementType.Kind()) // struct
		slice := reflect.MakeSlice(sliceValue.Type(), 0, 0)
		//result = reflect.New(slice.Type())
		//result.Elem().Set(slice)

		tmp := reflect.New(sliceElementType)
		tmpElem := tmp.Elem()
		//result2 := reflect.Append(sliceValue, tmp)
		tmpElem.FieldByName("Name").Set(reflect.ValueOf("fizz"))
		tmpElem.FieldByName("Age").Set(reflect.ValueOf(18))
		fmt.Printf("%v\n",tmp)
		fmt.Println(tmp.Type())
		fmt.Println(tmp.Elem().NumField())
		fmt.Println(tmp.Elem().Field(0).Type())
		fmt.Println(slice.Type())
		sliceValue = reflect.Append(sliceValue, tmpElem)
		//sliceValue = reflect.Append(sliceValue, tmpElem)
		slice = reflect.Append(slice, tmpElem)
		//slice = reflect.Append(slice, tmp.Elem())
		fmt.Println(slice, sliceValue)

		//sliceValue.FieldByIndex([]int{0}).Set(tmpElem.Elem())

		//sliceValue.Set(tmpElem)
	} else if sliceValue.Kind() == reflect.Struct { // 单一struct对象
		fmt.Println(sliceValue.Kind()) // struct
		result = reflect.New(sliceValue.Type())
		result.Elem().Set(sliceValue)

		result.Interface()
	}

	//for rows.Next() {
	//	var newValue = newElemFunc(fields)
	//	bean := newValue.Interface()
	//
	//	switch elemType.Kind() {
	//	case reflect.Slice:
	//		err = rows.ScanSlice(bean)
	//	case reflect.Map:
	//		err = rows.ScanMap(bean)
	//	default:
	//		err = rows.Scan(bean)
	//	}
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	if err := containerValueSetFunc(&newValue, nil); err != nil {
	//		return err
	//	}
	//}

	return
}
