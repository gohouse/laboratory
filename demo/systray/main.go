package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"time"
)
func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("新消息+22")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)

	for {
		systray.SetTitle("+22")
		time.Sleep(1500*time.Millisecond)
		systray.SetTitle("+++")
		time.Sleep(500*time.Millisecond)
	}

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.SetTitle("test 2344214123421341241234123")
		}
	}
}

func onExit() {
	// clean up here
	systray.Quit()
}