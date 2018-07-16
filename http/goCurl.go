package main

import (
	"github.com/mikemintang/go-curl"
	"fmt"
	"encoding/json"
)

type rr struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}

func main() {
	url := "http://fooddrug.service-alpha.wochacha.cn/tasklist/test"
	resp, _ := curl.NewRequest().SetUrl(url).Get()
	fmt.Println(resp.Body)

	var js rr
	json.Unmarshal([]byte(resp.Body), &js)

	fmt.Println(js)
}
