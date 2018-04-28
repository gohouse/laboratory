package main

import (
	"github.com/gohouse/utils"
	"fmt"
)

func main() {
	//res := utils.SuccessReturn("fail",204)
	res := utils.FailReturn()

	fmt.Println(res)
}
