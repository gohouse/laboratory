package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(rand.Intn(10))
	//}
	randomA()
}

func randomA() {

	for i := 0; i < 10; i++ {

	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Float64())
}
