package main

import (
	"github.com/skratchdot/open-golang/open"
)
func main() {
	//cmd := exec.Command("/Users/mac/go/bin/fyne_demo")
	//cmd.Start()
	//open.RunWith("/Users/mac/go/bin/fyne_demo", "iTerm2")
	open.Run("/Users/mac/go/src/github.com/gohouse/demo/gl/screen.png")
	//cmd := exec.Command("/Users/mac/go/src/github.com/gohouse/demo/gl/screen.png")
	//open.Run("screen.png")
}
