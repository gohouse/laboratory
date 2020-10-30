package main

import (
	"fmt"
	"github.com/gohouse/demo/pubsub"
)

var ch1 = "ch1"
var ch2 = "ch2"
func main() {
	var client1 = pubsub.NewClient("client1")
	var client2 = pubsub.NewClient("client2")
	var ps = pubsub.NewPubsub(pubsub.NewChannel(ch1, 10))
	ps.AddChannel(pubsub.NewChannel(ch2, 10))

	ps.Subscribe(ch1, client1)
	ps.Subscribe(ch2, client1)
	ps.Subscribe(ch2, client2)

	go ps.Broadcast()

	interactive(ps)
}

func interactive(ps *pubsub.Pubsub)  {
	var i int
	for {
		fmt.Println("请输入发布信息:")
		var ch, msg string
		fmt.Scanln(&ch, &msg)
		ps.Publish(ch, &pubsub.Topic{
			Title:   fmt.Sprintf("sequence %d",i),
			Content: msg,
		})
		i++
	}
}
