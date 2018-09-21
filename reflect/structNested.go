package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

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

func main() {
	var cluster DbConfigCluster
	decode(&cluster)
	fmt.Println(cluster.Master)
}

func decode(dec interface{}) {

	srcVal := reflect.ValueOf(dec)
	srcElem := srcVal.Elem()
	fmt.Println(srcVal.Kind())
	fmt.Println(srcElem)
	fmt.Println(srcElem.FieldByName("Master").String())
	fieldMaster := srcElem.FieldByName("Master")
	fmt.Println(fieldMaster.IsValid())
	fmt.Println(fieldMaster.Type().Elem())
	singleConf := reflect.New(fieldMaster.Type().Elem())

	fmt.Println(singleConf)

	//=======
	var fp []byte
	fp, err2 := ioutil.ReadFile(
		"/Users/fizz/go/src/github.com/gohouse/gorose/examples/demoParserFiles/mysql_cluster.json")
	if err2 != nil {
		return
	}
	//fmt.Println(string(fp))
	var tmp DbConfigSingle
	var master = fieldMaster.Elem()
	err := json.Unmarshal(fp, &dec)
	if err != nil {
		return
	}

	fmt.Println(master)
	fmt.Println(tmp)
	fmt.Println(dec)
}