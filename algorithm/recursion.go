package main

import "fmt"

func main()  {
	var data = []int{1,2,3,4}

	fmt.Println(recursion(data,[]int{},0))
}

func recursion(data []int, res []int, i int) []int  {
	if len(data)==i {
		return res
	}

	res = append(res, data[i])
	i++
	return recursion(data,res,i)
}
