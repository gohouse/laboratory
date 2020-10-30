package pubsub

import (
	"fmt"
	"testing"
)

func TestNewPubsub(t *testing.T) {
	var client1 = NewMember("client1")
	var client2 = NewMember("client2")
	var ps = NewPubsub(NewChannel("testchannel", 0))
	ps.Subscribe("testchannel",client1)
	ps.Subscribe("testchannel",client2)

	go ps.Broadcast()

	interactive(ps)
}

func interactive(ps *Pubsub)  {
	var i int
	for {
		fmt.Printf("请输入发布信息:")
		var msg string
		fmt.Scanln(&msg)
		ps.Publish("testchannel", &Topic{
			Title:   fmt.Sprintf("msg of %d",i),
			Content: msg,
		})
		i++
	}
}
