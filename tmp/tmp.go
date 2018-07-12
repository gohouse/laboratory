package main

import (
	"fmt"
	"github.com/gohouse/laboratory/utils"
)

type HttpHander struct{}
type Mstring string

func main() {

	a := 104
	for i := 0; i < 100; i++ {
		fmt.Println("insert into idfa_users(mobile) values("+utils.ParseStr(a+i)+");")
	}

}
