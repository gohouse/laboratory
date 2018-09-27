package main

// 冒泡排序
import (
	"fmt"
)

func BubbleSort(values []int) {
	Len := len(values)
	for i := 0; i < Len-1; i++ {
		for j := i + 1; j < Len; j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	BubbleSort(values)
	fmt.Println(values)
	return
}
