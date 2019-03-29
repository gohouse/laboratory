package main

import (
	"fmt"
	"github.com/gohouse/gorose/utils"
	"time"
)

func main() {
	var result = make(chan int,100)
	var num = 50

	defer close(result)

	go producer(result, num)

	go consumer(result, num)

	time.Sleep(time.Second)
}

func consumer(c chan int, num int)  {
	for i := 0; i < num; i++ {
		select {
		case res := <-c:
			fmt.Println(res)
		}
	}
}

func producer(c chan int, num int)  {
	for i := 0; i < num; i++ {
		go func() {
			c <- utils.GetRandomNum(2)
		}()
	}
}
