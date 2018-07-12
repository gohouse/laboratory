package main
import (
	"fmt"
	"net"
	"sync"
)
var bufPool sync.Pool
type buf struct {
	b []byte
}


func main() {
	ln, _ := net.Listen("tcp", ":8082")


	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
		}


		//为每一条进来的连接都开启一个goroutine处理
		go func() {
			var bf *buf
			v := bufPool.Get()
			if v == nil {
				//若不存在buf，创建新的
				fmt.Println("no buffer ,need create!")
				bf = &buf{
					b: make([]byte, 10),
				}
			} else {
				// 池里存在buf,v这里是interface{}，需要做类型转换
				bf = v.(*buf)
			}


			//从conn中连续读取数据
			//注意conn.Read有个特点，如果bf.b容量不足，那么会分两次读取
			//而不会去扩展bf.b
			for {
				_, err := conn.Read(bf.b)
				//这里只是例子，因此没有对io.EOF做特殊处理
				if err != nil {
					fmt.Printf("conn error: %#v", err)
					break
				}
			}
			// bf使命完成，放入池中
			bufPool.Put(bf)
		}()
	}
}