package main

import (
	"net"
	"fmt"
	"time"
	"strings"
)

var (
	connType string // long / short
)

func init()  {
	connType = "long"
}

func main() {
	// 超时
	timeout := 10
	// 设置监听地址和端口
	network := "localhost:1024"
	// 创建socket链接, 监听链接端口
	netListen, err := net.Listen("tcp", network)
	if err != nil {
		LogErr("创建socket链接失败: ", err)
	}
	// 延迟关闭监听链接
	defer netListen.Close()
	fmt.Println("开始监听链接......")

	// 开始接收任务操作
	for {
		receiver(netListen, timeout)
	}
}

func receiver(netListen net.Listener, timeout int) {
	// 开始接收任务操作
	conn, err := netListen.Accept()
	if err != nil {
		//continue
	}
	fmt.Println(conn.RemoteAddr().String(), ": 链接成功!")

	if connType == "short" {
		LogErr("定时器~短链接测试...")
		// 短链接, 使用定时器
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		// 处理链接
		handleConn(conn)
	} else {
		LogErr("心跳~长链接测试...")
		go handleConnection(conn, timeout)
	}
}

//长连接入口
func handleConnection(conn net.Conn, timeout int) {

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			checkConnErr(conn, err)
			return
		}

		Data := (buffer[:n])
		msgChain := make(chan byte)
		//心跳计时  
		go HeartBeating(conn, msgChain, timeout)
		//检测每次Client是否有数据传来  
		go GravelChannel(Data, msgChain)
		LogErr("receive data length:", n)
		LogErr(conn.RemoteAddr().String(), "receive data string:", string(Data))
	}
}

//心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func HeartBeating(conn net.Conn, msgChain chan byte, timeout int) {
	select {
	case fk := <-msgChain:
		LogErr(conn.RemoteAddr().String(), "receive data string:", string(fk))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * time.Duration(timeout)):
		LogErr("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func GravelChannel(n []byte, msgChain chan byte) {
	for _, v := range n {
		msgChain <- v
	}
	close(msgChain)
}
func handleConn(conn net.Conn) {
	// 数据缓冲区
	buffer := make([]byte, 2048)

	// 开始处理
	for {
		strLen, err := conn.Read(buffer)

		if err != nil {
			checkConnErr(conn, err)
			return
		}

		fmt.Println(conn.RemoteAddr().String(), " 获取数据: ", string(buffer[:strLen]))
	}
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
