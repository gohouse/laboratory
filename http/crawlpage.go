package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)
//https://blog.csdn.net/Maggie_up有问题欢迎指正!
func determineEncodings(r io.Reader) []byte {
	OldReader := bufio.NewReader(r)
	bytes, err := OldReader.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	reader := transform.NewReader(OldReader, e.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return all
}

func main() {
	//https://blog.csdn.net/Maggie_up有问题欢迎指正!
	resp, err := http.Get(`http://0-w.cc/`)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf(`%s`, determineEncodings(resp.Body))
}

