package main

import "fmt"

func pass(left, right chan int){
	left <- 1 + <- right
}

func main(){
	const n = 50
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i< n; i++ {
		right = make(chan int)
		// the chain is constructed from the end
		go pass(left, right) // the first goroutine holds (leftmost, new chan)
		left = right         // the second and following goroutines hold (last right chan, new chan)
	}
	go func(c chan int){ c <- 1}(right)
	fmt.Println("sum:", <- leftmost)
}