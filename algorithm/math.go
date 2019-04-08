package main

import (
	"fmt"
	"math"
)

var leaveArr []float64

// 声明公因数变量,余数变量
var commonLowestArr []float64

//var leaveArr []float64
// Lowest Common Multiple 最小公倍数
func LCM(intArr []float64) float64 {
	var i float64
	for i = 1; i < 10; i++ {
		var leaveArr2 []float64
		for k, item := range intArr {
			res := float64(item) / i
			leaveArr2 = append(leaveArr2, res)
			if k == len(intArr)-1 && res == math.Floor(res) {
				commonLowestArr = append(commonLowestArr, i)
				//fmt.Println(commonLowestArr)
				leaveArr = leaveArr2
				LCM(leaveArr2)
			}
			//commonLowestArr = append(commonLowestArr, i)
		}
		//leaveArr = make([]float64,0)
	}

	fmt.Println(leaveArr, commonLowestArr)
	return 0
}

func main() {
	intArr := []float64{4, 6, 8}
	res := LCM(intArr)
	//res := float64(5)/2
	//if res==math.Floor(res) {
	//
	//} else {
	//	fmt.Println(234)
	//}

	fmt.Println(res)
	//for _, item := range intArr {
	//	fmt.Println(item)
	//}
}
