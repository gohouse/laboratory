package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"fmt"
)

func main() {

	//fi, err := os.Open("/Users/fizz/go/src/github.com/gohouse/study/test.json")
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	return
	//}
	//defer fi.Close()
	//
	//br := bufio.NewReader(fi)
	//for {
	//	a, _, c := br.ReadLine()
	//	if c == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(a))
	//}
	ReadLine("/Users/fizz/go/src/github.com/gohouse/study/test.json", Print)
}
func Print(line string) {
	fmt.Println(line)
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}