package config

import (
	"testing"
	"fmt"
)

func TestConfigParser_NewConfigParser(test *testing.T) {

	pr, err := NewConfigParser(ConfigFiles["json"],"json")

	if err != nil {
		test.Error("FAIL: read file failed.", err)
		return
	}

	//fmt.Println(os.Getenv("GOPATH"))
	//fmt.Println(build.Default.GOPATH)

	test.Log(fmt.Sprintf("PASS: json %v", pr))
}


