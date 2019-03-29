package main

import (
	"fmt"
)

func main() {
	var days int = 365
	var cash float64 = 10000
	var rate float64 = 4.5
	var rateYear float64 = 5.28
	// 年华总收益
	yearTotal := cash + cash * rateYear / 100
	// 日化福利年总收益
	var tmpDay float64 = cash
	for i := 0; i < days; i++ {
		tmpDay = tmpDay + tmpDay * rate / 100 / float64(days)
	}

	fmt.Println(yearTotal, tmpDay)
}
