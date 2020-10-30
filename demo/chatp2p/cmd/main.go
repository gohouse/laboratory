package main

import "github.com/gohouse/demo/chatp2p"

func main() {
	chat := chatp2p.NewChat()
	println(chat)
	chat.GetUserInfo()
}
