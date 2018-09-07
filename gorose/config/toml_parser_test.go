package config

import (
	"testing"
	"fmt"
)

func TestConfigParser_Toml(test *testing.T) {
	//var file = "/Users/fizz/go/src/github.com/gohouse/laboratory/dp/config/mysql.toml"
	var file = ConfigFiles["toml"]

	var confP = &TomlConfigParser{}

	pr, err := confP.Parse(file)

	if err != nil {
		test.Error("FAIL: toml parser failed.", err)
		return
	}

	test.Log(fmt.Sprintf("PASS: toml parser %v", pr))
}


