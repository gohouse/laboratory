package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	//// 超时
	//timeout := 10
	// 设置监听地址和端口
	network,_ := net.ResolveUDPAddr("udp4", ":1025")
	// 创建socket链接, 监听链接端口
	socket, err := net.ListenUDP("udp4", network)
	if err != nil {
		LogErr("创建socket链接失败: ", err)
	}
	// 延迟关闭监听链接
	defer socket.Close()
	fmt.Println("开始监听链接......")

	// 开始接收任务操作
	for {
		receiver(socket)
	}
}

func receiver(socket *net.UDPConn) {
	data := make([]byte, 20)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read data ", err)
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)
	go func() {
		for i := 0; i < 3; i++ {
			send_data := []byte("hi client!")
			_, err = socket.WriteToUDP(send_data, remoteAddr)
			if err != nil {
				fmt.Println("send fail!", err)
				return
			}
		}
	}()
}

func checkConnErr(conn net.Conn, err error) {
	if err.Error() == "EOF" {
		LogErr(conn.RemoteAddr().String(), " 获取数据结束: ", err)
	} else if strings.Contains(err.Error(), "timeout") {
		LogErr(conn.RemoteAddr().String(), " 获取数据超时: ", err)
	} else {
		LogErr(conn.RemoteAddr().String(), " 获取数据失败: ", err)
	}
}

func LogErr(err ...interface{}) {
	fmt.Println(err...)
}
