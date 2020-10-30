package main

import (
	"github.com/go-vgo/robotgo"
	"log"
	"time"
)

func main() {
	w, h := robotgo.GetScreenSize()
	log.Printf("当前窗口大小: %v x %v\n", w, h)
	log.Printf("当前应用的pid: %v\n", robotgo.GetPID())
	log.Printf("GetActive: %v\n", robotgo.GetActive())
	log.Printf("GetHandle: %v\n", robotgo.GetHandle())
	findName, _ := robotgo.FindName(robotgo.GetPID())
	log.Println("FindName:", findName)	// goland

	//return
	//log.Printf("ActivePID: %v\n", robotgo.ActivePID(147539))	// 不能用
	//log.Printf("ActiveName: %v\n", robotgo.ActiveName("goland"))	// goland
	//log.Printf("ActiveName: %v\n", robotgo.ActiveName("Telegram"))	// Telegram
	//log.Printf("ActiveName: %v\n", robotgo.ActiveName("Sublime Text"))	// sublime
	log.Printf("ActiveName: %v\n", robotgo.ActiveName("go4rdm"))	// sublime
	//names, _ := robotgo.Process()
	//for _,v := range names {
	//	log.Println(v.Pid, v.Name)
	//}

	robotgo.KeyTap("f", "control", "command")
	time.Sleep(800*time.Millisecond)
	screen := robotgo.CaptureScreen()
	robotgo.SaveBitmap(screen, "screen.png")//保存位图为图片
	time.Sleep(500*time.Millisecond)
	robotgo.KeyTap("f", "control", "command")

	log.Printf("ActiveName: %v\n", robotgo.ActiveName(findName))	// iTerm2
}
