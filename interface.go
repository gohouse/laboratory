package main

import "fmt"

type A interface{
	Test() string
}
type AT struct {

}

func (a *AT) Test() string {
	return "babsbsf"
}


func main()  {
	//var a interface{} = 3

	var b AT
	var c A
	//b = 3

	fmt.Println(b.Test())
	fmt.Println(c.Test())
}
