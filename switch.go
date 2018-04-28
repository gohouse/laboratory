package main

import "fmt"

func main()  {
	var a interface{} = 3

	switch res := a.(type) {
	case int:
		fmt.Println(res)
	}
}
