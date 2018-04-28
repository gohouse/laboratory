package main

import (
	"github.com/sipt/GoJsoner"
	"fmt"
)
func main() {

	result, err := GoJsoner.Discard(`
		{//test comment1
			"name": "测试",
			"url":"http://url.cn",
			/**
			test comment2
			1
			2
			3
			end
			*/
			"age":26 //test comment3
			/***asdf**/
		}
	`)

	fmt.Println(err)
	fmt.Println(result)
}
