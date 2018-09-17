package parser

import (
	"fmt"
	"github.com/gohouse/laboratory/gorose/config"
	"testing"
)

func TestFileParser_New(test *testing.T) {

	pr, err := NewFileParser(config.DemoParserFiles["json"], "json")

	if err != nil {
		test.Error("FAIL: read file failed.", err)
		return
	}

	//fmt.Println(os.Getenv("GOPATH"))
	//fmt.Println(build.Default.GOPATH)

	test.Log(fmt.Sprintf("PASS: json %v", pr))
}
