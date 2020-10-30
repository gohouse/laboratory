package main

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"log"
)

func main()  {
	mapCapLenTest()
}

func mapCapLenTest() {
	s := map[int]int{1:1,2:2,3:3,4:4,5:5,6:6}
	printMap(s)
	s[31] = 3
	printMap(s)
	s[41] = 4
	printMap(s)
	s[51] = 5
	printMap(s)
}
func printMap(s map[int]int) {
	fmt.Printf("len=%d  %v\n", len(s), s)
}

func sliceCapLenTest() {
	s := []int{2}
	printSlice(s)
	s = append(s, 33)
	printSlice(s)
	s = append(s, 33)
	printSlice(s)
	s = append(s, 33)
	printSlice(s)
	s = append(s, 33)
	printSlice(s)
	//s[1] = 2

	s = append(s, 33,44,2,2,3,4,9,1,3,5,9,9,2,5,6,2,44,2,2,3,4,9,1,3,5,9,9,2,5,6,2,44,2,2,3,4,9,1,3,5,9,9,2,5,6,2,44,2,2,3,4,9,1,3,5,9,9,2,2,3,3,5,9,9,2,2,3,3)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main3() {
	id, err := machineid.ProtectedID("myAppName")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
func main2() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
