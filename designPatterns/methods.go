package main

import "fmt"

type aT struct {
	field []string
}

func (a aT) a1() aT {
	a.field = append(a.field, "a1")
	return a
}
func (a aT) b1() aT {
	a.field = append(a.field, "b1")
	return a
}
func (a aT) show()  {
	fmt.Println(a)
}

func main()  {
	var a aT
	tmp := a.a1().b1()
	tmp=tmp.a1()
	tmp.show()
}
