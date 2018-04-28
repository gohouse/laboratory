package main

import (
	"net"
	"fmt"
	"os"
	"time"
	"strconv"
)

func main() {
	// 设置链接地址和端口
	network := "localhost:1024"
	// 设置协议版本和地址
	tcpAddr,err := net.ResolveTCPAddr("tcp4", network)
	if err!=nil{
		fmt.Fprintf(os.Stderr, "设置协议版本和地址失败: %s", err.Error())
		os.Exit(1)
	}
	// 开始链接
	conn,err := net.DialTCP("tcp",nil, tcpAddr)
	if err!=nil{
		fmt.Fprintf(os.Stderr, "链接失败: %s", err.Error())
		os.Exit(1)
	}

	// 开始发送消息
	//handlerSend(conn)
	//sender(conn)

	for i := 20; i <25; i++ {
		words:= strconv.Itoa(i)+" 定时器测试~~~"
		conn.Write([]byte(words))
		time.Sleep(1*time.Second)
	}
	fmt.Println("send over")
}

