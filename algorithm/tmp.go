package main

import "fmt"

func main() {
	var a = "deg34_@#$%^&*()wabc"

	res := hash(a)

	fmt.Println(res)
}

func hash(k interface{}) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
		//fmt.Println(h)
	}
	return h
}
