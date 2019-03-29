package main

import (
	"fmt"
	"time"
)

func main() {
	tickerdemo()
}
func tickafter() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	var flag = 1
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			flag += 1
			if flag == 3 {
				return
			}
		default:
			fmt.Println(" .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func tickerdemo() {
	ticker := time.NewTicker(time.Millisecond * 500)
	c := make(chan int, 3) //num为指定的执行次数
	go func(ticker *time.Ticker) {
		for t := range ticker.C {
			c <- 1
			fmt.Println("Tick at", t)
		}
	}(ticker)

	//ticker.Stop()


	for {
		fmt.Println(<-c)
	}

	time.Sleep(1000*time.Microsecond)

	//fmt.Println(len(c))
}


func tickerdemo2() {
	done := make(chan struct{},2) // 用来等待协程结束

	timer := time.NewTimer(time.Second * 3)

	go func() {
		for item := range timer.C {
			fmt.Printf("Now is %s\n", item)
			done <- struct{}{}
		}
	}()

	timer.Stop()

	fmt.Println("Print in main")

	<-done
	close(done)
}

