package main

import (
	"flag"
	"fmt"
)

func parseFlag(key string) string {
	s := flag.String(key,
		"/Users/fizz/go/src/github.com/gohouse/laboratory/script/wcc/compile.sh", "配置文件地址")
	//"/Users/fizz/go/src/github.com/gohouse/laboratory/script/wcc/db.toml", "配置文件地址")
	flag.Parse()

	return *s
}

func main() {
	fmt.Println(parseFlag("conf"))
}
