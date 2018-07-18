package main

import "fmt"

var panicErr interface{}

func init()  {
	defer recovers()
}

func main() {
	//fmt.Println(1)
	//boot()
	fmt.Println(2)
	work()
	fmt.Println(panicErr)

}

func work() {
	panic("erroasfsdofjsdf")
}


func recovers() {
	if err := recover(); err!=nil{
		panicErr = err
	}
}

func boot() {
	defer recovers()
}

