package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
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
	// 单一数据库配置
	type DbConfigSingle struct {
		Driver          string // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog  bool   // 是否开启sql日志
		SetMaxOpenConns int    // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns int    // (连接池)闲置的连接数, 默认-1
		Prefix          string // 表前缀
		Dsn             string // 数据库链接
	}

	// 数据库集群配置
	// 如果不启用集群, 则直接使用 DbConfig 即可
	// 如果仍然使用此配置为非集群, 则 Slave 配置置空即可, 等同于使用 DbConfig
	type DbConfigCluster struct {
		Slave         []*DbConfigSingle // 多台读服务器, 如果启用则需要放入对应的多台从服务器配置
		Master        *DbConfigSingle   // 一台主服务器负责写数据
	}

	var cg DbConfigCluster
	var cpath string = "src/github.com/gohouse/laboratory/config/toml/file.toml"
	res,_ := ioutil.ReadFile(cpath)
	if err := toml.Unmarshal(res, &cg); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cg)
	js,_:= json.Marshal(cg)
	fmt.Println(string(js))
	fmt.Println(cg.Master)
}

func main() {
	test_unmarshal()
}