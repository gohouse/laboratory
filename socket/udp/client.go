package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

func main() {
	network, err := net.ResolveUDPAddr("udp4", ":1025")
	socket, err := net.DialUDP("udp4", nil, network)
	if err != nil {
		fmt.Println("connect fail !", err)
		return
	}
	defer socket.Close()

	//for {
		for i := 0; i < 3; i++ {
			senddata := []byte("hi server!"+strconv.Itoa(i))
			_, err = socket.Write(senddata)
			if err != nil {
				fmt.Println("send fail !", err)
				return
			}

			data := make([]byte, 20)
			read, remoteAddr, err := socket.ReadFromUDP(data)
			if err != nil {
				fmt.Println("read fail !", err)
				return
			}
			fmt.Println(read, remoteAddr)
			fmt.Printf("%s\n", data)
			time.Sleep(time.Second * 1)
		}
	//}
}