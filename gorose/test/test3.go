package main

import "fmt"

type iusb interface {
	input()string
	output()string
}
type mouse struct {
}
type upan struct {
}

func (m mouse) input() string {
	return "mouse"
}
func (m mouse) output() string {
	return "output"
}

func (u upan) input() string {
	return "upan"
}
func (u upan) output() string {
	return "output"
}

func main() {

	//var usb iusb = new(mouse)
	var usb iusb = new(upan)
	fmt.Println(usb.input())
}
