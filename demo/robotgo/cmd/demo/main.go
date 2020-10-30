package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"time"
)
type Pos struct {
	X,Y int
}

func NewPos(x int, y int) *Pos {
	return &Pos{X: x, Y: y}
}
var (
	qunMingZi = NewPos(400,20)
	searchEntry = NewPos(50,25)
)
func main() {
	var app = "Mango"
	pos := findPos(app)
	clickSearchEntry(app, pos)
	searchText("ron")
}
func searchText(text string)  {
	time.Sleep(1*time.Second)
	robotgo.KeyTap(text)
}
func clickSearchEntry(app string, pos *Pos)  {
	robotgo.ActiveName(app)
	robotgo.Move(pos.X+searchEntry.X, pos.Y+searchEntry.Y)
	time.Sleep(1*time.Second)
	robotgo.MouseClick("left")
}
func findPos(app string) (pos *Pos) {
	//var pos = searchEntry
	w, h := robotgo.GetScreenSize()
	log.Printf("当前窗口大小: %v x %v\n", w, h)

	robotgo.ActiveName(app)

	//bitmap:=robotgo.OpenBitmap("1.jpg")
	//cutBm:=robotgo.GetPortion(bitmap,5,5,200,60)
	//robotgo.SaveBitmap(cutBm,"cutBm.png")
	bitmap := robotgo.OpenBitmap("search.png")
	time.Sleep(1*time.Second)
	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------", fx, fy)
	//time.Sleep(1*time.Second)
	return NewPos(fx,fy)
}

func capture(app string)  {

	robotgo.ActiveName(app)
	//robotgo.KeyTap("f", "control", "command")
	time.Sleep(1*time.Second)
	//bitmap := robotgo.CaptureScreen(5, 5, 200, 60)
	bitmap := robotgo.CaptureScreen()
	robotgo.SaveBitmap(bitmap, "screen.png")
	time.Sleep(1*time.Second)
	//robotgo.KeyTap("f", "control", "command")
}
