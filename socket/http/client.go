package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type person struct {
	NAME string
	AGE  int
}

func main() {
	tmp := person{"xx", 1}
	data, err := json.Marshal(tmp) //转化为JSON
	if err != nil {
		fmt.Println("json Marshal err", err)
	}
	resp, err := http.Post("http://127.0.0.1:1026/test", "application/x-www-form-urlencoded", strings.NewReader(string(data)))
	fmt.Println(data)
	fmt.Println(string(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //读取服务器返回的信息
	if err != nil {
		fmt.Println("read err")
	}
	fmt.Println(string(body))
}
