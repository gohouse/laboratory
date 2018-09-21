package bench_gorose

import (
	"fmt"
	"github.com/gohouse/gorose"
	_ "github.com/gohouse/gorose/driver/mysql"
	"testing"
)

type users struct {
	Name string `orm:"name"`
	Age int `orm:"age"`
	Job string `orm:"job"`
}
// DB Config.(Recommend to use configuration file to import)
var dbConfig = &gorose.DbConfigSingle {
	Driver:          "mysql", // 驱动: mysql/sqlite/oracle/mssql/postgres
	EnableQueryLog:  false,   // 是否开启sql日志
	SetMaxOpenConns: 0,    // (连接池)最大打开的连接数，默认值为0表示不限制
	SetMaxIdleConns: 0,    // (连接池)闲置的连接数, 默认-1
	Prefix:          "", // 表前缀
	Dsn:             "gcore:gcore@tcp(192.168.200.248:3306)/test?charset=utf8", // 数据库链接
}
var Connection *gorose.Connection

func init()  {
	var err error
	Connection, err = gorose.Open(dbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestStruct_First(test *testing.T)  {

	db := Connection.NewSession()
	var user users
	err2 := db.Table(&user).Select()
	if err2 != nil {
		test.Error("FAIL: open failed.", err2)
		return
	}
	test.Log(fmt.Sprintf("PASS: open: %v", user))
}


func BenchmarkStruct_First(bmtest *testing.B) {
	db := Connection.NewSession()
	var user users
	for cnt := 0; cnt < bmtest.N; cnt++ {
		db.Table(&user).First() // 200	   6498108 ns/op
	}
}