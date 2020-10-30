package main

import (
	"github.com/gen2brain/beeep"
	"time"
)

func main() {
	act()
	ticker := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-ticker.C:
			act()
		}
	}
}

func act()  {
	beeep.Alert("喝水提醒", "喝水的时候到了, 可以眺望一下远方, 放松一下眼睛 ^_^", "Clock.icns")
}
