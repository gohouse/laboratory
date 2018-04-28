package main

import (
	"net"
	"bufio"
	"ftj-data-synchro/protocol"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"bytes"
	"ftj-data-synchro/logic"
	"fmt"
	"strconv"
)

/*
 客户端结构体
 */
type Client struct {
	DeviceID string        //客户端连接的唯标志
	Conn     net.Conn      //连接
	reader   *bufio.Reader //读取
	writer   *bufio.Writer //输出
	Data     []byte        //接收数据
}

func NewClient(conn *net.TCPConn) *Client {
	reader := bufio.NewReaderSize(conn, 10240)
	writer := bufio.NewWriter(conn)
	c := &Client{Conn: conn, reader: reader, writer: writer}
	return c
}

/**
 数据读取（粘包处理）
 */
func (this *Client) read() {
	for {
		var data []byte
		var err error
		//读取指令头 返回输入流的前4个字节，不会移动读取位置
		data, err = this.reader.Peek(4)
		if len(data) == 0 || err != nil {
			continue
		}
		//返回缓冲中现有的可读取的字节数
		var byteSize = this.reader.Buffered()
		fmt.Printf("读取字节长度：%d\n", byteSize)
		//生成一个字节数组，大小为缓冲中可读字节数
		data = make([]byte, byteSize)
		//读取缓冲中的数据
		this.reader.Read(data)
		fmt.Printf("读取字节：%d\n", data)
		//保存到新的缓冲区
		for _, v := range data {
			this.Data = append(this.Data, v)
		}
		if len(this.Data) < 4 {
			//数据包缓冲区清空
			this.Data = []byte{}
			fmt.Printf("非法数据，无指令头...\n")
			continue
		}
		data, err = protocol.HexBytesToBytes(this.Data[:4])
		instructHead, _ := strconv.ParseUint(string(data), 16, 16)
		//指令头效验
		if uint16(instructHead) != 42330 {
			fmt.Printf("非法数据\n")
			//数据包缓冲区清空
			this.Data = []byte{}
			continue
		}
		data = this.Data[:protocol.HEADER_SIZE]
		var p = protocol.Decode(data)
		fmt.Printf("消息体长度：%d\n", p.Len)
		var bodyLength = len(this.Data)
		/**
		 判断数据包缓冲区的大小是否小于协议请求头中数据包大小
		 如果小于，等待读取下一个客户端数据包，否则对数据包解码进行业务逻辑处理
		 */
		if int(p.Len) > len(this.Data)-protocol.HEADER_SIZE {
			fmt.Printf("body体长度：%d,读取的body体长度：%d\n", p.Len, bodyLength)
			continue
		}
		fmt.Printf("实际处理字节：%v\n", this.Data)
		p = protocol.Decode(this.Data)
		//逻辑处理
		go this.logicHandler(p)
		//数据包缓冲区清空
		this.Data = []byte{}
	}
}
