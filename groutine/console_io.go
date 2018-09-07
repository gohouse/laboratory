package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

var scanner2  *bufio.Scanner = bufio.NewScanner(os.Stdin)
var inputChan2 chan bool = make(chan bool)

func GetInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := strings.Trim(scanner.Text(),"\n")
	return  input
}

func jiaohu()  {
	// 用户输入协程
	go func(inputChan chan bool) {
		fmt.Println("按'b'回到菜单列表,按'q'退出程序...")
		for{
			input := strings.Trim(GetInput(scanner2),"\n")
			switch input {
			case "b":
				inputChan <- true
				return
			case "q":
				inputChan <- false
				return
			default:
				fmt.Println("输入错误,请重新输入...")
			}
		}
	}(inputChan2)
	// 主协程阻塞channel
	//for  {
		select {
		case r := <-inputChan2:
			if r {
				//this.DrawMenu()
				//close(inputChan)
				fmt.Println(r)
				//close(inputChan2)
				//os.Exit(0)
			}else if !r{
				fmt.Println(r)
				//close(inputChan2)
				//os.Exit(0)
			}
		}
	//}
}

func main()  {
	jiaohu()
}