package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i, math.Pow(2, float64(i)), math.Pow(3, float64(i)))
	}
}
