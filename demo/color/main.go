package main

import (
	"fmt"
	"os"
	"time"

	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
)

func main() {
	testColor()
}
func testColor()  {
	ct.Foreground(ct.Green, false)
	fmt.Println("Green text starts here...")
	ct.ResetColor()
	fmt.Println("abc")
	ct.ChangeColor(ct.Red, true, ct.White, false)
	fmt.Println("123")
	ct.ResetColor()
}
func msgTest() {
	writer := goterminal.New(os.Stdout)
	for i := 0; i < 5; i++ {
		ct.Foreground(ct.Red, false)
		fmt.Fprintln(writer, "I'm in yellow.")
		writer.Print()
		ct.ResetColor()

		time.Sleep(time.Second)

		fmt.Fprintln(writer, "Lets change above text to green.")
		writer.Print()
		time.Sleep(time.Second)

		// processing done here, after which color should change or text should be over-written.
		writer.Clear()

		ct.Foreground(ct.Green, false)
		fmt.Fprintln(writer, "I'm in green now.")
		writer.Print()
		ct.ResetColor()

		fmt.Fprintln(writer, "Great!")
		writer.Print()
		time.Sleep(time.Second)

		writer.Reset()
	}
}
func main2() {
	writer := goterminal.New(os.Stdout)
	for i := 0; i < 5; i++ {
		ct.Foreground(ct.Red, false)
		fmt.Fprintln(writer, "I'm in yellow.")
		writer.Print()
		ct.ResetColor()

		time.Sleep(time.Second)

		fmt.Fprintln(writer, "Lets change above text to green.")
		writer.Print()
		time.Sleep(time.Second)

		// processing done here, after which color should change or text should be over-written.
		writer.Clear()

		ct.Foreground(ct.Green, false)
		fmt.Fprintln(writer, "I'm in green now.")
		writer.Print()
		ct.ResetColor()

		fmt.Fprintln(writer, "Great!")
		writer.Print()
		time.Sleep(time.Second)

		writer.Reset()
	}

}
