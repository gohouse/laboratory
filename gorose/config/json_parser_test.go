package config

import (
	"testing"
	"fmt"
)

func TestConfigParser_Json(test *testing.T) {
	//var file = "/Users/fizz/go/src/github.com/gohouse/laboratory/dp/config/mysql.json"
	var file = ConfigFiles["json"]

	var confP = &JsonConfigParser{}

	pr, err := confP.Parse(file)

	if err != nil {
		test.Error("FAIL: json parser failed.", err)
		return
	}

	test.Log(fmt.Sprintf("PASS: json parser %v", pr))
}


