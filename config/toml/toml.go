package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"io/ioutil"
)

type songInfo struct {
	Name     string
	Id int
	JokeName string `toml:"jokeName"`
	Area []map[string]string
}

type config struct {
	Province string
	Tag string
	City songInfo
}

func test_toml() {
	var cg config
	var cpath string = "src/github.com/gohouse/laboratory/config/toml/config.toml"
	if _, err := toml.DecodeFile(cpath, &cg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v %v\n", cg.Province, cg.City.JokeName)
	fmt.Printf("%v %v\n", cg.Tag, cg.City.Area)
}

func test_unmarshal() {
	var cg config
	var cpath string = "src/github.com/gohouse/laboratory/config/toml/config.toml"
	res,_ := ioutil.ReadFile(cpath)
	if err := toml.Unmarshal(res, &cg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v %v\n", cg.Province, cg.City.JokeName)
	fmt.Printf("%v %v\n", cg.Tag, cg.City.Area)
}

func main() {
	test_unmarshal()
}