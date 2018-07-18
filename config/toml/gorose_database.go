package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"fmt"
	"github.com/gohouse/gorose/utils"
	"github.com/gohouse/gorose"
)

type mysql_items struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
	Charset  string `toml:"charset"`
	Protocol string `toml:"protocol"`
	Prefix   string `toml:"prefix"`
	Driver   string `toml:"driver"`
}
type conf struct {
	Default         string                 `toml:"Default"`
	SetMaxOpenConns int                    `toml:"SetMaxOpenConns"`
	SetMaxIdleConns int                    `toml:"SetMaxIdleConns"`
	Connections     map[string]mysql_items `toml:"Connections"`
}


func main() {
	var cg conf
	var tmp interface{} = cg
	fmt.Println(tmp)
	var cpath string = "src/github.com/gohouse/laboratory/config/toml/gorose_database.toml"
	if _, err := toml.DecodeFile(cpath, &cg); err != nil {
		log.Fatal(err)
	}
	//var tmp = make(map[string]interface{})
	//tmp["DbConf"] = cg

	fmt.Println(cg)
	connection, err := gorose.Open(cg, "mysql_dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	// close DB
	defer connection.Close()

	db := connection.GetInstance()
	fmt.Println(db)
	res, err := db.Table("users").Where("id", "<", 1).First()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(res))
	fmt.Println(db.LastSql)
	fmt.Println(res)
	//fmt.Printf("%v\n",   cg.Connections[cg.DbConfig.Default])
	//
	var jsonstr, _ = utils.JsonEncode(cg)
	////var jsonstr = `{"DbConfig":{"Default":"mysql_yz","SetMaxOpenConns":300,"SetMaxIdleConns":10},"Connections":{"mysql_yz":{"Host":"180.97.188.201","Username":"idfas","Password":"Mysql777","Port":"3306","Database":"idfas","Charset":"utf8","Protocol":"tcp","Prefix":"idfa_","Driver":"mysql"}}}`
	fmt.Printf("%v\n", jsonstr)
	//var js conf
	//json.Unmarshal([]byte(jsonstr), &js)
	//
	//fmt.Println(js)
}
