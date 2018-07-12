package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"fmt"
	"github.com/gohouse/laboratory/utils"
	"encoding/json"
)

type dbConfig struct {
	Default         string
	SetMaxOpenConns int
	SetMaxIdleConns int
}
type mysql_items struct {
	Host string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Port string `toml:"port"`
	Database string `toml:"database"`
	Charset string `toml:"charset"`
	Protocol string `toml:"protocol"`
	Prefix string `toml:"prefix"`
	Driver string `toml:"driver"`
}
type conf struct {
	DbConfig dbConfig
	Connections map[string]mysql_items
}

func main() {
	var cg conf
	var cpath string = "src/github.com/gohouse/laboratory/config/toml/gorose_database.toml"
	if _, err := toml.DecodeFile(cpath, &cg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", cg)
	fmt.Printf("%v\n",   cg.Connections[cg.DbConfig.Default])

	var jsonstr = utils.JsonEncode(cg)
	//var jsonstr = `{"DbConfig":{"Default":"mysql_yz","SetMaxOpenConns":300,"SetMaxIdleConns":10},"Connections":{"mysql_yz":{"Host":"180.97.188.201","Username":"idfas","Password":"Mysql777","Port":"3306","Database":"idfas","Charset":"utf8","Protocol":"tcp","Prefix":"idfa_","Driver":"mysql"}}}`
	fmt.Printf("%v\n", jsonstr)
	var js conf
	json.Unmarshal([]byte(jsonstr), &js)

	fmt.Println(js)
}
