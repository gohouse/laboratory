package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {

	//fmt.Println(time.Now()[:18])
	timestr := fmt.Sprintf("%s", time.Now())
	fmt.Println(timestr[:19])

	var a = map[string]interface{} {"a":"b"}
	fmt.Println(reflect.TypeOf(a["a"]))
}
